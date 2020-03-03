[![CircleCI](https://circleci.com/gh/moemoe89/practicing-solr-golang.svg?style=svg)](https://circleci.com/gh/moemoe89/practicing-solr-golang)
[![codecov](https://codecov.io/gh/moemoe89/practicing-solr-golang/branch/master/graph/badge.svg)](https://codecov.io/gh/moemoe89/practicing-solr-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/moemoe89/practicing-solr-golang)](https://goreportcard.com/report/github.com/moemoe89/practicing-solr-golang)

# PRACTICING-SOLR-GOLANG #

Practicing Solr Using Golang (Martini Framework) with Go Mod as Programming Language, Solr as Search Platform

## Directory structure
Your project directory structure should look like this
```
  + your_gopath/
  |
  +--+ src/github.com/moemoe89
  |  |
  |  +--+ practicing-solr-golang/
  |     |
  |     +--+ main.go
  |        + api/
  |        + routers/
  |        + ... any other source code
  |
  +--+ bin/
  |  |
  |  +-- ... executable file
  |
  +--+ pkg/
     |
     +-- ... all dependency_library required

```

## Setup and Build

* Setup Golang <https://golang.org/>
* Setup Solr <https://lucene.apache.org/solr/>
* Under `$GOPATH`, do the following command :
```
$ mkdir -p src/github.com/moemoe89
$ cd src/github.com/moemoe89
$ git clone <url>
$ mv <cloned directory> practicing-solr-golang
```

## Running Application
Make config file for local :
```
$ cp config-sample.json config.json
```
Change Solr address & collection based on your config :
```
http://localhost:8983/solr
collection1
```
Build
```
$ go build
```
Run
```
$ go run main.go
```

## How to Run with Docker
Make config file for docker :
```
$ cp config-sample.json config.json
```
Change Solr address & collection based on your docker config :
```
http://solr:8983/solr
collection1
```
Build
```
$ docker-compose build
```
Run
```
$ docker-compose up
```
Stop
```
$ docker-compose down
```

## How to Starting Solr with Docker
Create core :
```
$ docker exec -it --user=solr solr bin/solr create_core -c collection1
```
Access Solr Admin UI :
```
open http://localhost:8983/
```

## How to Run Unit Test
Run
```
$ go test ./...
```
Run with cover
```
$ go test ./... -cover
```
Run with HTML output
```
$ go test ./... -coverprofile=c.out && go tool cover -html=c.out
```


## License

MIT