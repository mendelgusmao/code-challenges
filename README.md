# Supereasy App

This is a simple app exposing an API to authenticate users and retrieve car service providers information.

## Architecture

The application consists of a backend that connects with a MongoDB instance and provides a REST API.

## Setup

- Get an instance of MongoDB up and running
- Execute `go run backend/main.go`
- Execute `cd data; ./bootstrap.sh`

## Environment variables

* **SUPEREASY_ADDRESS**
 - Defines the address in which the server will be listening to,
* **SUPEREASY_DATABASEURL**
 - Defines the database URL used to connect to a MongoDB instance. Example: `mongodb://mongo:27017`.
* **SUPEREASY_DATABASENAME**
 - Defines MongoDB collection name. Example: `supereasy`.
* **SUPEREASY_JWTSECRET**
 - Defines a cryptographic secret used for authenticating API requests.
* **SUPEREASY_LOCATIONIQAPIKEY**
 - Defines an API key that is used to retrieve geocoding data from [LocationIQ](https://locationiq.com).

## REST API

#### Users

* **POST /users/authenticate** - Authenticates an user
  * *204 OK* and an `Authorization` header with the token if successful
  * *400 Bad Request* if the JSON payload is malformed
  * *403 Forbidden* if there's no such combination of user and password
  * *500 Internal Server Error* if there's an unexpected error


  Request:

  ```
  {
    "email": "john.doe@mail.org",
    "password": "p455w0rd"
  }
  ```

  The token should be stored and sent in the `Authorization` HTTP header when requesting the `/partners` endpoint.

#### Partners

* **GET /partners/one_by_location_and_service?service=&coordinates=** - Retrieves a partner
  * *200 OK* with an `application/json` body if at least one partner is within 10 kilometers of the provided coordinates and offers the desired service
  * *404 Not Found* if there are no partners offering the desired service within 10 kilometers of the provided coordinates
  * *403 Forbidden* if the Authorization token is invalid
  * *500 Internal Server Error* if there's an unexpected error


  Parameters:
  - coordinates: latitude and latitude pair of the desired location. Example: `-23.550278,-46.633889`
  - service: The desired car service. Example: `OIL_CHANGE` or `CAR_WASHING`


  Response:

  ```
  {
      "availableServices": [
          "OIL_CHANGE",
          "DRY_WASHING"
      ],
      "id": "cjsgiytw40006enwge0ushrqx",
      "location": {
          "address": "Praça Leonor Kaupa, 100 - Jardim da Saúde",
          "city": "São Paulo",
          "country": "Brazil",
          "lat": -23.619575,
          "long": -46.627023,
          "name": "Shopping Plaza Sul",
          "state": "SP"
      },
      "name": "Karla Albuquerque"
  }
  ```

* **GET /partners/by_address?address=** - Retrieves a list of partners
  * *200 OK* with an `application/json` body if at least one partner is within 10 kilometers of the provided address and offers the desired service
  * *404 Not Found* if there are no partners offering the desired service within 10 kilometers of the provided address
  * *403 Forbidden* if the Authorization token is invalid
  * *500 Internal Server Error* if there's an unexpected error


  Parameters:
  - address: Venue name of the desired location. Example: `Av. Brg. Faria Lima, 1355`


  Response:

  ```
  [
    {
        "availableServices": [
            "OIL_CHANGE",
            "DRY_WASHING"
        ],
        "name": "Karla Albuquerque"
    },
    {
        "availableServices": [
            "DRY_WASHING"
        ],
        "name": "Antônio Braga"
    },
    ...
  ```
