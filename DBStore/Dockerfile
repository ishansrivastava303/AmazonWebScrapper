FROM golang:latest

WORKDIR /app/

COPY . /app/

RUN export GO111MODULE=on

RUN go get github.com/mattn/go-sqlite3

RUN chmod +x DBstore.go

ENTRYPOINT ["go","run","./DBstore.go"]

