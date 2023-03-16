# Weekly Status Report #5

The project is nearing its finish at this point and now the main focus is tying everyone's work together. I made major strides on the backend this week.
All database queries are finished and I finally chose an approach to routing; I ended up using an approach that I'll refer to as 'ShiftPath'. This consisted of 
creating a custom root handler that passed http queries to other child handlers based off custom path parsing. The typical way to route in go is to create a mux object, 
and then add handler functions to your specified path strings. The shortcoming of this approach is the inability to route based on slugs or ids. Instead of this, I
elected to parse the url of the request manually. The allowed me to break the url into individual pieces that would be analyzable by their respective httphandler. I could then 
do my own case switching based off the url request and build out the api spec I originally intended. The spec is nearly identical to the one in the project proposal aside from 
a few minor changes. The main change was my decision to identify lists based off their name rather than their id number as per the original spec. This would allow for url queries 
to be much more human-readable...making scripting easier etc. I left note routing as Id numbers as those items are more likely to be short-lived and are usually not queried individually (though I did build in the 
functionality to do so). After I had implemented such the backend was finally ready to be queried/consumed by the frontend. I put together an api spec file using Insomnia and 
gave that to my group members along with a docker-compose file for interacting with the results of my backend work. The spec file includes documentation of all possible api queries 
and the docker-compose file spins up a cockroachdb container along with a containerized version of my go backend. I was able to then meet with my group and supply them with said resources, 
and now the only thing left to do is let them finalize the frontend.
