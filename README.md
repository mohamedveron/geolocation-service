# geolocation-service
This component is developed using go-chi https://github.com/go-chi/chi router and oapi-codegen https://github.com/deepmap/oapi-codegen to generate http client for other services and using go mod for dependancies management.

# Use api/api.yml file in the swagger ui to see the description of the generated apis

# How to use the compnent in your service
The component generated http client to expose the geolocation data for other services.

You have to import the client in your go mod file by run the following command

```bash
go get github.com/mohamedveron/geolocation-service/client
```
## Test import geolocation data process
By run 
```bash
make test
```
You will start the procees of importing the geolocation data to the db and will take around 1m and 33.66s time 

## Setup

Must have golang installed version >= 12.0.0

make file consists of 4 steps: generate, test, build, run
you can run all of them 

```bash
make all
```

# Run the unit tests:
```bash
make test
```
If you have issue with code generation in generate step you can run go test -v ./tests

# Start the http server on port 8080:

```bash
make run
```

If you have issue with code generation in generate step you can copy the api/api.gen.go file in repo and run go run main.go to start the server

# Build and run docker image

```bash
docker build --tag geolocation-service .
```

``bash
docker run -p 9090:9090 geolocation-service
```
