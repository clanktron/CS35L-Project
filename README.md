# CS35L Final Project : Jot

an over-engineered to-do app.

Tech Stack:
- Frontend: React
- Backend: Go, CockroachDB

## Development Setup:

For backend development:

Inside the backend directory theres a /dev folder, this contains a hot reload script (since its not built in to golang) that can help with development speed.

To build the container for production simply run:

```
./package/build.sh
```

For Frontend Development:

Download and Install [Docker Desktop](https://www.docker.com/products/docker-desktop/) 

>After installation ensure you can run the `docker` command in your terminal

In the root directory of the project run the following:

```
docker compose up
```

This will start two docker containers. One is the cockroachdb database and the other is the backend golang code.
These processes will both output to the terminal while they run, so if you need another terminal for commands I would open another tab etc.

Now Install [Insomnia](https://insomnia.rest/download)

You can now import the json file I sent in the discord chat and read the documentation on the different api endpoints. If you follow the imported document from
top to bottom it should be relatively self explanatory. Insomnia also supports generating code corresponding to the query so that might be useful.

At this point you can start up your react code like usual and query the data how you like. Use the urls from the insomnia document but just know the backend is 
listening on port 4000.

Notes:
- By default I create an admin user with a single list containing no notes so you'll have to POST some data before you can query it.
- You can also watch the output of the golang code as you run queries since it'll log notable events and errors.

