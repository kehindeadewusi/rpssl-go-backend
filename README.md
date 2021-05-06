# Overview
A go backend, implementing rock,paper,scissors,spock,lizard. It's supposed to be very popular :)

# RPSSL Backend (This code)
RPSSL = Rock, Paper, Scissors, Spock, Lizard.

This implements all the game service.

## Project Setup
1. Clone the repo and change directory inside rpssl-go-backend `git clone github.com/kehindeadewusi/rpssl-go-backend`
1. Run `go install` to setup dependencies. 
1. The application uses some reasonable default for the IP & PORT to start the application on. You can override these values in a number of ways, including via environment variables.
You can override the following 2 values.
    - RPSSL_HOST : Server host address, defaults to 127.0.0.1
    - RPSSL_PORT : Server port number, defaults to 8081
1. Run `go run main.go`

## Docker
The project includes a Docker file that uses a 2-stage build; builds in the first and creates a **FROM SCRATCH** base image in the second step.
You may use the docker build command to create an image
```
docker build -t tag:version .
```

## API Docs
I'm relying on Swagger for the documentation of the API. 

go-swagger dependency exists in the project, via `go get -u github.com/go-swagger/go-swagger/cmd/swagger`

To generate the specs, you would run `swagger generate spec -o ./swagger.yaml --scan-models`

To serve and preview in a browser, `swagger serve -F=swagger swagger.yaml`

Gracias!