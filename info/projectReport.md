---
geometry: "left=3cm,right=3cm,top=2cm,bottom=2cm"
---

# CS35L Project Report

## Team Members and Details:

- Marlee Kitchen, marlee.kitchen@gmail.com, Dis 1D, Vepa
- Yuzhou Gao, yuzhougao@g.ucla.edu, Dis 1E, Jason Kimko
- Daniel Wang, d.w.901102@gmail.com, Dis 1E, Jason Kimko
- Clayton Castro, crdcastro5@gmail.com, Dis 1D, Vepa
- Leo Naddell, leonaddell@gmail.com, Dis 1D, Vepa


## Over-engineered To-do App

### Features:

* users
    * id (unique int64 identification number)
    * username 
        * stored in plaintext
    * password 
        * minimum 8 chars
        * require a special and uppercase char
        * stored vals should be hashed and salted
* lists (groups of notes)
    * id (unique int64 identification number)
    * userid (corresponding userid)
* notes
    * id (unique int64 identification number)
    * userid (corresponding userid)
    * listid (corresponding listid)

### Software Architecture:

#### Frontend (ReactJS)

- Login Page (Marlee Kitchen):
   - functionality implemented primarily using the imported useRef, useState, and useEffect from React
   - display of the page varied upon submission via clicking the "log in" button, hitting "enter" on the keyboard, or clicking the "create account" button
   - the functionality of the "log in" button and hitting "enter" on the keyboard were synonymous and would either switch the page to the main page by passing the props to the onFormSwitch function if the entered username and password corresponded to a matching pair stored in the backend database or would stay on the login page with an added "incorrect username or password" error message displayed if no such pair exists
   - clicking the "create account" button switches the displayed page from the login page to the register page, also, via the onFormSwitch function that React offers

- Create new user page (Daniel Wang):
    - Input username and password to create new user
    - Set limitation for not allowing special character in username
    - Set requirement for needing to include special character in password
    - Require the password to entered twice to comfirm the password
    - Check for reponse from backend whether the username has been used

#### Backend (Golang) (Clayton)

- communicates with CockroachDB
- serves a REST API for the Frontend to consume
- REST Schema: (prefix: https://\<backendurl\>/api/v0/)
    - /user
        - POST: should accept username and password
            - returns authentication cookie/token (stored in browser by frontend)
            - only returns cookie first time user is created
    - /login
        - POST: should accept username and password
            - returns authentication cookie/token (stored in browser by frontend)
    - /lists
        - GET: return all lists for user
        - POST: create new list
    - /lists/{listname}
        - GET: return list metadata
        - PUT: update list metadata
        - DELETE: delete list
    - /lists/{listname}/notes
        - GET: return all notes for list
        - POST: create new note
    - /lists/{listname}/notes/{noteid}
        - GET: returns note data
        - PUT: update note data
        - DELETE: delete note

#### Data Persistence (Clayton)

- CockroachDB
    - Postgres compatible DB
    - built to run in a distributed/fault-tolerant fashion

### General

- All 3 services (CockroachDB UI, golang api, react app) will be TLS terminated with a publicly valid cert if the manifests within the k8s directory
are used for deployment.

Frontend and Backend are to be deployed in separate containers on a Kubernetes cluster along with the required instance of CockroachDB.
Deployment can be also be done to plain docker or bare metal (though this is discouraged for production deployments as it would defeat the 
purpose of the fault-tolerant architecture).

All services can scale independently and can operate independently, meaning as long as the database (single source of truth) is valid the application state is stable.
This means the containers for all three can operate at essentially infinite scale, only being limited by the number of machines you can supply
to the cluster.

#### Future Thoughts

It would be nice to have some user attributes/preferences, like a light/dark mode or maybe some custom css for their homepage.

#### Takeaways (Clayton Castro)

I learned a lot from this project, notably how design a custom REST api in golang (a language I had no previous experience with). Since I didn't use 
an existing routing framework I ended up making my own, which is a nice bonus since I can use such in later projects.
