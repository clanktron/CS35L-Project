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

##### Frontend (ReactJS)


##### Backend (Golang)

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

- CockroachDB
    - Postgres compatible DB
    - built to run in a distributed/fault-tolerant fashion
    - DB Schema: TODO

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
