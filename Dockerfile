FROM golang:latest

WORKDIR /app

COPY . /app

RUN export GO111MODULE=on

RUN go get github.com/PuerkitoBio/goquery

RUN chmod +x amazonwebscraper.go

ENTRYPOINT ["go","run","./amazonwebscraper.go"]

