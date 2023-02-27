FROM golang:1.20.1-alpine3.17
WORKDIR /code
RUN apk add entr fish
CMD [ "fish" ]
