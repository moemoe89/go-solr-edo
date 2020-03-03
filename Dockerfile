FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/practicing-solr-golang

WORKDIR /go/src/github.com/moemoe89/practicing-solr-golang

COPY . /go/src/github.com/moemoe89/practicing-solr-golang

RUN go mod download
RUN go install

ENTRYPOINT /go/bin/practicing-solr-golang
