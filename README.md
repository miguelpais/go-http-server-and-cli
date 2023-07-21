# Go Small Http Server and CURL cli

This project puts forward a small socket-based toy http server leveraging Goroutines for concurrent processing of incoming requests and the serving of two routes, one json based and the other html-based.

The http-server includes a token-based rate limiter controlling the maximum rate at which the server handles requests, queeing the rate limited ones in an internal queue up to a certain maximum.

It then also includes a CURL-like tool to fetch the http response of any http based url.

## Usage

### Server

    go cmd/server/main/main.go

### CURL tool

    go cmd/cli/main/main.go http://localhost:8000
