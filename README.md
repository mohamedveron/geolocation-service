# geolocation-service
This component is developed using go-chi https://github.com/go-chi/chi router and oapi-codegen https://github.com/deepmap/oapi-codegen to generate http client for other services and using go mod for dependancies management.

# Use api/api.yml file in the swagger ui to see the description of the client interface

# How to use the component in your service

The component generated http client to expose the geolocation data for other services.

You have to import the latest version of the exposed client in your go mod file by run the following command

```bash
go get github.com/mohamedveron/geolocation-service/client
```


## Setup of the component

Must have golang installed version >= 12.0.0

make file consists of 4 steps: generate, test, build, run
you can run all of them 

```bash
make all
```
This app use an AWS RDS postgres instance: host-> geolocation.czqumefsqwp6.eu-central-1.rds.amazonaws.com:5432

If you need to use local database, run this qurey in the deafult postgres db to start importing process

CREATE TABLE public.geolocation (
	ip_address varchar(25) PRIMARY KEY,
	country varchar(25) NULL,
	country_code varchar(25) NULL,
	city varchar(25) NULL,
	latitude float8 NULL,
	longitude float8 NULL,
	mystery_value int8 NULL
);


## Test import geolocation data process by run

```bash
make test
```
or

go test -v ./tests

You will start the procees of importing the geolocation data to the db and will take around 1m and 33.66s time if you used the local db,
If you used the default AWS rds instance it will take more time because of the network latency 


# Start the http server on port 8080:

```bash
make run
```

# Build and run docker image

```bash
docker build --tag geolocation-service .
```

```bash
docker run -p 9090:9090 geolocation-service
```
