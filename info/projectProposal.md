---
geometry: "left=3cm,right=3cm,top=2cm,bottom=2cm"
---

# CS35L Project Proposal

## Team Members and Details:

- Marlee Kitchen, marlee.kitchen@gmail.com, Dis 1D, Vepa
- Yuzhou Gao, yuzhougao@g.ucla.edu, Dis 1E, Jason Kimko
- Daniel Wang, d.w.901102@gmail.com, Dis 1E, Jason Kimko
- Clayton Castro, crdcastro5@gmail.com, Dis 1D, Vepa
- Leo Naddell, leonaddell@gmail.com, Dis 1D, Vepa


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

>This is an ideal featureset, some minor nuances like hashed passwords may not be accomplished.

### Software Architecture:

##### Frontend (ReactJS (possibly with some Bootstrap?))

- pretty webui to display todo item content
    - login page
        - logo with username and password field (nothing else)
    - main app page
        - sidebar with todo lists
        - remaining right hand side should just be todo items
        - possible navbar on top?

##### Backend (Golang)

- communicates with CockroachDB
- serves a REST API for the Frontend to consume
- REST Schema: (prefix: https://\<backendurl\>/api/v0/)
    - /login
        - POST: should accept username and password
            - returns authentication cookie/token (stored in browser by frontend)
    - /lists
        - GET: return all lists for user
        - POST: create new list
    - /lists/{listid}
        - GET: return list metadata
        - POST: update list metadata
        - DELETE: delete list
    - /lists/{listid}/notes
        - GET: return all notes for list
        - POST: create new note
    - /lists/{listid}/notes/{noteid}
        - GET: returns note data
        - PUT: update note data
        - DELETE: delete note

>Exact REST Schema is subject to change

- CockroachDB
    - Postgres compatible DB
    - built to run in a distributed/fault-tolerant fashion
    - DB Schema: TODO
- All 3 services will be TLS terminated with a publicly valid cert.

Frontend and Backend will be deployed in separate containers on a Kubernetes cluster along with the required instance of CockroachDB.
Deployment can be done either on a local cluster on one of our laptops or on some dorm hardware. 
