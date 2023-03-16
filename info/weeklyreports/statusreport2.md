# Weekly Status Report #2

This week was spent building the "base-layer" of the project. I "finalized" the development environments for both the frontend and backend code. For now we're using
vanilla react for our frontend and nodejs as our runtime but there are certainly other options we could switch to later on. Astro or Next.js are straightforward options if we want to integrate
routing easily (which might be desirable for the login page and create-user page). For the backend I was initially leaning towards golang's builtin http library for implementing the API, however after a little 
research it seems that there's a framework called Echo that could significantly speed up development time. As of now I'm going to stick with the http library to keep dependencies at a minimum and gain some comfortability
dealing with raw http requests. I always have the option to switch later as the migration would be fairly trivial. Roles for the frontend were also chosen this week. Marlee is going to start with the login page and 
Daniel is going to take the create new user page. The main app page will be handled by Leo and Bryan. So far things have been going smoothly with the setup. Concerning backend development I've been able to serve 
some basic data over http along with basic routing for such. For a smoother development experience I'm using `entr` to hot reload my go code anytime I save my `.go` files. Next week I hope to create a rudimentary 
implementation of a REST api. It'll be (mostly) identical to our final spec but it'll just serve dummy data for now; no real database queries will be made. I'm prioritizing such to ensure the frontend team members have 
enough time to practice querying and handling the backend data. 

Note: I projected looking into JWTs last week but did not get to it :(. Since this is more of a nicety for the project I'm leaving it for after a MVP is working.
