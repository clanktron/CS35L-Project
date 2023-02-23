# Notes

- Routing? Bare `net/http` it is.
- Using `golang.org/x/crypto/bcrypt` for password hashing
- most likely using `github.com/golang-jwt/jwt` for auth

#### Misc 

- consider maybe serving react app from go binary
    - might be "harder" than keeping them separate
    - though might be a nice thing to do as an alternative deployment method (if there's time)
