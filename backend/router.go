package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
)

// ShiftPath splits off the first component of p, which will be cleaned of
// relative components before processing. head will never contain a slash and
// tail will always be a rooted path without trailing slash.
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

type App struct {
	LoginHandler *LoginHandler
	UserHandler  *UserHandler
	ListHandler  *ListHandler
}

type UserHandler struct {
	db *sql.DB
}

type LoginHandler struct {
	db *sql.DB
}

type ListHandler struct {
	NoteHandler *NoteHandler
	db          *sql.DB
}

type NoteHandler struct {
	db *sql.DB
}

func (h *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)
	switch head {
	case "user":
		h.UserHandler.ServeHTTP(w, r)
		return
	case "login":
		h.LoginHandler.ServeHTTP(w, r)
		return
	case "list":
		h.ListHandler.ServeHTTP(w, r)
		return
	default:
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

// accepts POST requests with new user payloads - responds with jwt or error
func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	enableCors(w)

	switch r.Method {
	case http.MethodPost:

		var newuser User
		var err error

		// read json payload into new user object
		if err := parseJSON(w, r, &newuser); err != nil {
			log.Print(err)
			return
		}

		// add user to database
		if err := addUser(h.db, newuser); err != nil {
			log.Print(err)
			return
		}

		// reassign stored values in user object
		newuser, err = getUser(h.db, newuser.Username)
		if err != nil {
			log.Print(err)
			return
		}
		log.Printf("User %s has been successfully created", newuser.Username)

		// add default list to user
		defaultList := List{
			Userid: newuser.Id,
			Name:   "my tasks",
		}
		if err := addList(h.db, defaultList); err != nil {
			log.Print(err)
			return
		}
		log.Printf("Default list has been created for %s", newuser.Username)

		// respond with jwt
		token, err := generateJWT(newuser)
		if err != nil {
			log.Print(err)
			return
		}
		fmt.Print(token, "\n")

		// provide token as cookie to frontend
		ck := http.Cookie{
			Name:     "token",
			Value:    token,
			MaxAge:   3600,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
		}
		http.SetCookie(w, &ck)
		w.Write([]byte(fmt.Sprintf("created user %s", newuser.Username)))

		return
	case http.MethodDelete:

		// parse username from jwt token
		userid, err := verifyJWT(r)
		if err != nil {
			log.Print(err)
			http.Error(w, "Missing or expired cookie", 401)
			return
		}

		var user User
		// parse username and password into JSON object
		if err := parseJSON(w, r, &user); err != nil {
			log.Print(err)
			return
		}
		// check credentials
		if err := verifyUser(h.db, user); err != nil {
			log.Print(err)
			return
		}

		// remove row with corresponding userid
		if err := deleteUser(h.db, userid); err != nil {
			log.Print(err)
			log.Printf("Failed to delete user with id %d", userid)
			w.Write([]byte(fmt.Sprintf("Failed to delete user with id %d", userid)))
			return
		}
		w.Write([]byte(fmt.Sprintf("Successfully deleted user with id %d", userid)))
		return
	case http.MethodOptions:
		return
	default:
		http.Error(w, fmt.Sprintf("Expected method POST, OPTIONS or DELETE, got %v", r.Method), http.StatusMethodNotAllowed)
		fmt.Printf("Expected method POST, OPTIONS or DELETE, got %v", r.Method)
		return
	}

}

// accepts POST requests with user credentials - responds with a jwt or error
func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	enableCors(w)

	_, err := verifyJWT(r)
	if err == nil {
		log.Print("User is already logged in\n")
		w.Write([]byte(fmt.Sprintf("User is already logged in")))
		return
	}

	switch r.Method {
	case http.MethodPost:

		var login User

		// payload should contain username and password of user (other fields are ignored)
		if err := parseJSON(w, r, &login); err != nil {
			log.Print(err)
			return
		}
		// checks user+pass against stored database values
		if err := verifyUser(h.db, login); err != nil {
			log.Print(err)
			return
		}
		// fill user object with necessary information (userid)
		login, err := getUser(h.db, login.Username)
		if err != nil {
			log.Print(err)
			return
		}
		// creates jwt for user
		token, err := generateJWT(login)
		if err != nil {
			log.Print(err)
			return
		}
		// provide token as cookie to frontend
		ck := http.Cookie{
			Name:     "token",
			Value:    token,
			MaxAge:   3600,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
		}
		// responds with message to store cookie in browser for future requests
		http.SetCookie(w, &ck)

		log.Printf("%s logged in successfully", login.Username)
		w.Write([]byte(fmt.Sprintf("%s logged in successfully", login.Username)))

		return
	case http.MethodOptions:
		return
	default:
		http.Error(w, fmt.Sprintf("Expected method POST or OPTIONS, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var listname string
	listname, r.URL.Path = ShiftPath(r.URL.Path)

	if r.URL.Path != "/" {
		var head string
		head, r.URL.Path = ShiftPath(r.URL.Path)
		switch head {
		case "note":
			h.NoteHandler.Handler(listname).ServeHTTP(w, r)
		default:
			http.Error(w, "Not Found", http.StatusNotFound)
		}
		return
	}

	enableCors(w)

	// verify user is authenticated and store userid
	userid, err := verifyJWT(r)
	if err != nil {
		log.Print(err)
		return
	}

	if listname == "" {
		switch r.Method {
		case http.MethodGet:

			// get array of list objects
			lists, err := getLists(h.db, userid)
			if err != nil {
				log.Print("Error getting lists from database")
				log.Print(err)
				return
			}
			// send array of lists as http response
			if err := renderJSON(w, lists); err != nil {
				log.Print(err)
				return
			}

			return
		case http.MethodPost:

			var newlist List

			// parse payload into newlist object
			if err := parseJSON(w, r, &newlist); err != nil {
				log.Print(err)
				return
			}

			// set userid to jwt result to prevent impersonation
			newlist.Userid = userid

			// add list to database
			if err := addList(h.db, newlist); err != nil {
				log.Print(err)
				return
			}

			log.Printf("Added \"%s\" to lists!", newlist.Name)
			w.Write([]byte(fmt.Sprintf("Added \"%s\" to lists!", newlist.Name)))

			return
		case http.MethodOptions:
			return
		default:
			http.Error(w, fmt.Sprintf("Expected method GET, POST, or OPTIONS got %v", r.Method), http.StatusMethodNotAllowed)
			return
		}
	}

	switch r.Method {
	case http.MethodGet:

		list, err := getList(h.db, listname, userid)
		if err != nil {
			log.Print("Error getting list from database")
			log.Print(err)
			return
		}
		if err := renderJSON(w, list); err != nil {
			log.Print(err)
			return
		}

		return
	case http.MethodPut:

		var updatedlist List
		if err := parseJSON(w, r, &updatedlist); err != nil {
			log.Print("Error parsing json payload\n")
			log.Print(err)
			return
		}
		if err := updateList(h.db, updatedlist, listname, userid); err != nil {
			log.Printf("Error updating list %s\n", listname)
			log.Print(err)
			return
		}
		log.Printf("Updated list \"%s\"\n", updatedlist.Name)
		w.Write([]byte(fmt.Sprintf("Updated list \"%s\"!", updatedlist.Name)))

		return
	case http.MethodDelete:
		var list List
		list.Userid = userid
		list.Name = listname
		if err := deleteList(h.db, list); err != nil {
			log.Print(err)
			log.Print("Failed to delete list\n")
		}
		log.Printf("Deleted list %s\n", list.Name)
		w.Write([]byte(fmt.Sprintf("Deleted list \"%s\"!", list.Name)))

		return
	case http.MethodOptions:
		return
	default:
		http.Error(w, fmt.Sprintf("Expected method GET, POST, PUT, OPTIONS, or DELETE, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}
}

func (h *NoteHandler) Handler(listname string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var head string
		head, r.URL.Path = ShiftPath(r.URL.Path)

		if r.URL.Path != "/" {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		enableCors(w)

		userid, err := verifyJWT(r)
		if err != nil {
			log.Print(err)
			return
		}

		// get details of parent list for use later
		var parentlist List
		parentlist, err = getList(h.db, listname, userid)
		if err != nil {
			log.Print(err)
			http.Error(w, fmt.Sprintf("Invalid list identifier %s", listname), http.StatusBadRequest)
			return
		}
		listid := parentlist.Id

		// set noteid to postfix of url if it ends with such
		var noteid int64
		if head != "" {
			noteid, err = strconv.ParseInt(head, 10, 64)
			if err != nil {
				log.Printf("Invalid note id %d", noteid)
				http.Error(w, fmt.Sprintf("Invalid note id %d", noteid), http.StatusBadRequest)
				return
			}
		}

		if noteid == 0 {
			switch r.Method {
			case http.MethodGet:

				// get all notes corresponding to userid in token and listid from listname
				notes, err := getNotes(h.db, userid, listid)
				if err != nil {
					log.Print(err)
					log.Print("Failed to get notes from database\n")
					return
				}
				if err := renderJSON(w, notes); err != nil {
					log.Print(err)
					log.Printf("Json response failed\n")
					return
				}

				return
			case http.MethodPost:

				var newnote Note
				// read note content into note object
				if err := parseJSON(w, r, &newnote); err != nil {
					log.Print(err)
					log.Printf("Failed to parse json payload\n")
					return
				}

				// ensure ids are not overwritten
				newnote.Userid = userid
				newnote.Listid = listid

				// add note to database
				if err := addNote(h.db, newnote); err != nil {
					log.Print(err)
					log.Print("Failed to add note to database\n")
					return
				}
				log.Printf("Added note to list with id %d", newnote.Listid)
				w.Write([]byte(fmt.Sprintf("Added note to list with id %d", newnote.Listid)))

				return
			case http.MethodOptions:
				return
			default:
				http.Error(w, fmt.Sprintf("Expected method GET, POST, or OPTIONS got %v", r.Method), http.StatusMethodNotAllowed)
				return
			}
		}

		switch r.Method {
		case http.MethodGet:

			// retrieve full note object corresponding to given noteid
			notes, err := getNote(h.db, noteid)
			if err != nil {
				log.Print(err)
				log.Print("Failed to get notes from database\n")
				return
			}
			if err := renderJSON(w, notes); err != nil {
				log.Print(err)
				log.Printf("Json response failed\n")
				return
			}

			return

		case http.MethodPut:

			var updatednote Note
			// read note content into note object
			if err := parseJSON(w, r, &updatednote); err != nil {
				log.Print(err)
				log.Printf("Failed to parse json payload\n")
				return
			}

			// ensure values are not overwritten
			updatednote.Userid = userid
			updatednote.Listid = listid
			updatednote.Id = noteid

			if err := updateNote(h.db, updatednote); err != nil {
				log.Print(err)
				log.Print("Failed to update note in database\n")
				return
			}
			log.Printf("Updated note with content \"%s\"", updatednote.Content)
			w.Write([]byte(fmt.Sprintf("Note with id %d has been updated with content \"%s\"", noteid, updatednote.Content)))

			return
		case http.MethodDelete:

			if err := deleteNote(h.db, noteid); err != nil {
				log.Print(err)
				log.Print("Failed to delete note from database\n")
				return
			}
			log.Printf("Note with id %d has been deleted", noteid)
			w.Write([]byte(fmt.Sprintf("Note with id %d has been deleted", noteid)))

			return
		case http.MethodOptions:
			return
		default:
			http.Error(w, fmt.Sprintf("Expected method GET, POST, PUT, or DELETE, got %v", r.Method), http.StatusMethodNotAllowed)
			return
		}

	})
}
