# Weekly Status Report #1

This was our first week of working on the project and we now have a basic repo structure and project outline. We're still in the process of sorting out
exactly how we'll handle pull requests but all members have been added to the repo (hosted on my personal Github account). Our project is comprised of 
a separate frontend and a backend that will be worked on in separate teams. I will be primarily responsible for backend api development
while my teammates will be taking the lead with the react based UI. For this starting week I started setting up the necessary development environments in the repo; this includes Dockerfiles,
basic setup scripts, etc. Concerning backend development I now have an extremely basic rest api working in go. For now it seems going with the builtin http library is the 
best option for building the API. I also did some testing with deploying cockroachdb to a local kubernetes cluster. The UI is TLS terminated as is the SQL port (though I
have yet to do proper testing of the sql port). Using a self-hosted instance of the database would allow the entire app to live in a single cluster, 
though we may end up using the free hosted version of cockroachdb if things go awry with a self-hosted instance. Most of the aforementioned tasks are still a work in progress and I hope to finalize
them in the coming week. Next week I also hope to get further with the api and build more robust http interaction (handling error codes, actually catching errors, further JWT token research).
I'm also considering using the external jwt library for go to easily integrate authentication. 
