# MeGu App

This is a simple CRUD app.

## Architecture

The application consists of:
  * frontend server to provide interaction with a browser
  * a backend that connects with a database and provides a REST API to the frontend

## Docker image

```
docker pull mendelgusmao/me_gu:latest
```

## Environment variables

#### Frontend

* **ME_GU_FRONTEND_ADDRESS**
 - Defines the address in which the frontend server will be listening to
* **ME_GU_FRONTEND_BACKENDADDRESS**
 - Defines the address in which the backend server is listening
* **ME_GU_FRONTEND_SESSIONKEY**
 - Cryptographic key for the secure session and cookie
* **ME_GU_FRONTEND_TEMPLATES**
 - Absolute path to the directory where the HTML templates are stored

#### Backend

* **ME_GU_BACKEND_ADDRESS**
 - Defines the address in which the backend server will be listening to
* **ME_GU_BACKEND_DATABASE**
 - The database DSN used to connect to a database. Example: `root:password@/me_gu`
* **ME_GU_BACKEND_PASSWORDRESETURL**
 - A template for the URL that will be sent to the e-mail address to an user that is willing to reset its account password. The reset token will be appended.
* **ME_GU_BACKEND_PASSWORDRESETEXPIRATION**
 - Defines the maximum time a password reset token is valid.
* **ME_GU_BACKEND_PASSWORDRESETFROMADDRESS**
 - Defines the sender e-mail address for the password reset e-mail
* **ME_GU_BACKEND_SMTPADDRESS**
 - SMTP server address
* **ME_GU_BACKEND_SMTPPORT**
 - SMTP server port
* **ME_GU_BACKEND_SMTPUSER**
 - SMTP server username
* **ME_GU_BACKEND_SMTPPASSWORD**
 - SMTP server password

## Backend API

#### Users

* **GET /users/{id}** - Retrieves an user and responds with:
  * *200 OK* with a *application/json* body if successful
  * *404 Not Found* if there's no user wich such id
  * *500 Internal Server Error* if there's an unexpected error

  ```
  {
    "id": 1,
    "email": "john.doe@mail.org",
    "full_name": "John Doe",
    "telephone": "+1 (949) 555-1234",
    "address": "4590 MacArthur Blvd"
  }
  ```

* **POST /users** - Creates an user and responds with:
  * *201 Created* if successful
  * *400 Bad Request* if the JSON payload is malformed
  * *409 Conflict* if there's already an user with the e-mail address
  * *500 Internal Server Error* if there's an unexpected error

  ```
  {
    "email": "john.doe@mail.org",
    "full_name": "John Doe",
    "telephone": "+1 (949) 555-1234",
    "password": "p455w0rd",
    "address": "4590 MacArthur Blvd"
  }
  ```

* **PATCH /users/{id}** - Updates an user and responds with:
  * *204 No Content* if successful
  * *400 Bad Request* if the JSON payload is malformed
  * *404 Not Found* if there's no user wich such id
  * *409 Conflict* if there's already an user with the e-mail address
  * *500 Internal Server Error* if there's an unexpected error

#### User authentication

* **POST /users/authenticate** - Authenticates an user with its email and password
  * *200 OK* if email and password matches
  * *400 Bad Request* if the JSON payload is malformed
  * *403 Forbidden* if there's no user with such e-mail address or password doesn't match
  * *500 Internal Server Error* if there's an unexpected error

  ```
  {
    "email": "john.doe@mail.org",
    "password": "p455w0rd"
  }
  ```

#### Password recovery

* **POST /users/password-reset** - Creates a token for password reset and sends an e-mail if there's an user with that e-mail
  * *201 Created* if successful
  * *400 Bad Request* if the JSON payload is malformed
  * *404 Not Found* if there's no user with such e-mail address
  * *500 Internal Server Error* if there's an unexpected error

  ```
  {
    "email": "john.doe@mail.org",
    "password": "p455w0rd"
  }
  ```

  * **GET /users/password-reset/token/{token}** - Checks validity of password reset token
    * *204 No Content* if the token is valid
    * *403 Forbidden* if the token is not valid
    * *404 Not Found* if there's no token
    * *500 Internal Server Error* if there's an unexpected error


  * **POST /users/password-reset/password** - Updates user password
    * *204 No Content* if successful
    * *403 Forbidden* if the token is not valid
    * *404 Not Found* if there's no token
    * *500 Internal Server Error* if there's an unexpected error

    ```
    {
      "email": "john.doe@mail.org",
      "token": "e26d3dac-9ed7-4c0e-8c91-0cd145c2f058"
    }
    ```

## Database structure

```
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `full_name` varchar(255) DEFAULT '',
  `telephone` varchar(255) DEFAULT '',
  `address` varchar(255) DEFAULT '',
  `password_reset_token` varchar(255) DEFAULT '',
  `password_reset_token_expiration` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
```
