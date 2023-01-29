# Final Project

## Over-engineered To-do App

### Features:

* users
    * username (email)
        * stored in plaintext
    * password 
        * minimum 8 chars
        * require a special and uppercase char
        * stored vals should be hashed and salted
* lists (groups of notes)
    * title
    * uncompleted tasks
    * completed tasks
* tasks
    * task content
    * completed or not completed
    * completed tasks are archived

#### Software Architecture

* Frontend (JS or Rust/WASM if you're feeling spicy)
    * pretty webui to display todo item content
        * login page
            * logo with username and password field (nothing else)
        * main app page
            * sidebar with todo lists
            * remaining right hand side should just be todo items
* Backend (Golang)
    * communicates with DB 
    * serves a REST API for the Frontend to consume
    * REST Schema: (prefix: https://<backendurl>/api/v0/)
        * /login
            * POST: should accept username and password
                * returns authentication cookie/token (stored in browser by frontend)
        * /lists
            * GET: return all lists for user
            * POST: create new list
        * /lists/{listid}
            * GET: return list metadata
            * POST: update list metadata
            * DELETE: delete list
        * /lists/{listid}/notes
            * GET: return all notes for list
            * POST: create new note
        * /lists/{listid}/notes/{noteid}
            * GET: returns note data
            * PUT: update note data
            * DELETE: delete note
* CockroachDB
    * Postgres compatible DB
    * built run in a distributed/fault-tolerant fashion
    * DB Schema: TODO


## Other Options

### Markdown Notetaking App

#### Features:

* users
* folders
* files (markdown)

### Other notes

Three popular software architectures:
- Monolith (all on a single bare metal or VM server)
- Distributed Monolith (solves the problem of scaling but services are still tightly coupled)
- Microservices (pub/sub architecture, if one part of the app goes down the others continue to function)
