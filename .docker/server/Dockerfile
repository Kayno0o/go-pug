FROM golang:1.22-alpine as golang

WORKDIR /app/server

COPY watcher.sh go.mod go.sum ./

RUN apk update && apk add --no-cache inotify-tools make bash
RUN go install github.com/Joker/jade/cmd/jade@latest
RUN go get -u github.com/Joker/jade

EXPOSE 8080

CMD ["make", "docker"]
