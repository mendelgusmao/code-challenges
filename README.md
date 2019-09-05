# Tax Challenge App

This is a simple API server to track packages and compute them taxes as well

## Architecture

The application consists of:
  * a backend server that interacts with a BoltDB database and provides a REST API

## Building

#### Native Go
```
go run backend/main.go
```

#### Docker
```
docker build -t taxchallenge .
docker run -p 9091:9091 -v $PWD/data:/app/data taxchallenge
```

## Environment variables

* **TAXCHALLENGE_ADDRESS**
 - The network address the server will be listening to (default: **:9091**)
* **TAXCHALLENGE_DATABASE**
 - Path to the BoltDB database file (default: **taxchallenge.boltdb**)
* **TAXCHALLENGE_TAXRULES**
 - Path to the JSON file that contains the rules to calculate the taxes (default: **taxrules.json**)
