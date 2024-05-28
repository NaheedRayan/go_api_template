## This is a basic REST API using Golang and Mux


The Tutorial link I have followed [Link](https://hugo-johnsson.medium.com/rest-api-with-golang-and-mux-e934f581b8b5)


### Running the API (Default port 3000)

```
go run main.go
```


### Running the API on different port

```
go run main.go --listenAddr :2000
```


### Testing the API

```
curl http://localhost:3000
```