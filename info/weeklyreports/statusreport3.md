# Weekly Status Report #3

This week was spent focused on breaking my code down into more logical abstractions. I 
eventually ended up breaking most of my `main.go` into smaller go files that'll be called when necessary. 
I also made notable progress with my sql commands and http routing. My custom http handlers now support all 
types of requests (get, post, etc.). For a MVP (minimum viable product) the only things left to finish are leftover 
sql commands and id based routing. As I mentioned last time using a library like `Echo` might make it much easier to create the api. 
It would make id and slug based routing much easier and my job would essentially be over by now. However, since I'm stubborn and want to 
minimize external dependencies I'm going to stick with the basic http library that's built into go. I'll have to do some custom matching logic 
in order to implement id routing but that doesn't appear too daunting. I might accomplish it with a lengthy case-switch statement or maybe with some 
custom regex pattern matching. Since that doesn't seem like too big of a task it's likely I'll be able to get 
to some niceties like jwt authentication, proper logging, triple tls termination, etc. If time allows I might even 
integrate better command-line interaction; so flags can be passed to the binary, etc (this is lower on the list since it doesn't 
seem like a necessary or possibly even desirable feature). From a group perspective the rest of my teammates are getting close to prototyping their individual react component. 
As of now we have a very bare bones create new user page and login page. Due to the simplicity of the project design 
I hope our frontend stays as straightforward and manageable as possible...that being said react state is always a pain to manage.
I'm looking forward to meeting with my group next week and getting all the pieces of our application working together.
