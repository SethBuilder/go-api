# Intro

This is a rest api implemented with Go

It has two endpoints:

`GET /wines`
Returns all wines

`POST /wines`
Creates new wine


## Usage
Endpoints can be called by Curl or Postman.

### Curl
1. GET /wines: `curl https://seif-go-api.herokuapp.com/wines`
2. POST /wines: `curl https://seif-go-api.herokuapp.com/wines -X POST -d '{"name":"Bond Melbury","year":2008,"price": 500,"region":"Napa Valley","country":"United States"}' -H "Content-Type:application/json"`

### Postman
Simply import this [public collection](https://www.getpostman.com/collections/2457f64878b7ef56a2ff)

## How to run locally:
1. Download repo
2. `go run server.go`

## Tests
`go test`

## To enhance
1. Replace in-memory storge with a database.
2. Implement [the DDD Hexagon](https://youtu.be/1rxDzs0zgcE?t=1688) to replace the current flat file structure.
