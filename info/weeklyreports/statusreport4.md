# Weekly Status Report #4

Last week I projected meeting with my group to bring together all the different frontend pieces and connect them with the backend.
My group was unfortunately unable to get together, however each of our pieces seem to be coming along nicely. I did a little more refactoring
this week; moving the http handlers from a dedicated file to their respective object files and rearranging function abstractions in various places. So 
far I've finished these tasks in my `user.go` file and still have to clean up `list.go` and `note.go`.
I've also managed to add couple new features to the backend in the past few days. JWT authentication is now working at a bare-bones level and I can 
now successfully parse login requests and ensure they match their hashed, recorded counterparts. The largest task left to do is 
id based routing. I've been putting this off for a couple reasons, the first of which is I don't wish to reinvent the wheel. There are a number
of frameworks out there that have implemented features like this and creating my own routing solution may be a general waste of time. That being 
said, the current available frameworks are all not backed by any notable entities and could be abandoned without any prior warning. This taken into account, 
along with my desire to understand more about routing performance, is leading me to implement my own solution. My options at the moment seem to be either 
regex or custom string matching, though I plan on doing a little more research on what should be optimal performance-wise.
This weekend I plan on wrapping up any lingering database queries that need to be written and, at the very least, research routing solutions and 
their underlying mechanisms. Later in the week I hope to meet with my group to get an MVP up and running. 
