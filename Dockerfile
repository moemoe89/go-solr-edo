FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/go-solr-edo

WORKDIR /go/src/github.com/moemoe89/go-solr-edo

COPY . /go/src/github.com/moemoe89/go-solr-edo

RUN go mod download
RUN go install

ENTRYPOINT /go/bin/go-solr-edo
