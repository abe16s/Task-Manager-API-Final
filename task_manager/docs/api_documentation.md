# Task Management Service

A task management service built with Go and MongoDB, providing functionality for adding, updating, deleting, and retrieving tasks. This service offers a RESTful API for managing tasks and demonstrates the use of MongoDB for data storage.

## Configuration

Before running the application, ensure you have a MongoDB instance running. Update the database connection string, also the server port number and jwt secret key in .env

```
JWT_SECRET
MONGODB_URI
SERVER_PORT
```

[Postman documetation](https://documenter.getpostman.com/view/32032637/2sA3s3GAhh)

## Authorization & Authentication

Endpoints accessed by all registered users

```
GET localhost:8080/tasks
GET localhost:8080/tasks/:id
```

Endpoints accessed by only registered admins

```
POST localhost:8080/tasks
PUT localhost:8080/tasks/:id
DELETE localhost:8080/tasks/:id
PATCH localhost:8080/promote
```

## Register new user

```
POST localhost:8080/register
```

Registers a new user by accepting a JSON object containing the username and password.

Request: 

```json
{
  "username": "string",
  "password": "string"
}
```

Responses:

* 201 Created:
* Description: User registered successfully.

Body:

```json
{
  "_id": "9e484920-0871-49a3-9bcf-2b9a29e7ec09",
  "username": "abe16s",
  "password": "$2a$10$n3watE2dZ7WPz4oZA.3yIOXKPrBG5GUrOt8gg5WmpI3EpTg.NhLH.",
  "is_admin": true
}
```
* First registered user would be an admin by default

## User Login

```
POST localhost:8080/login
```

Authenticates a user and returns a JWT token upon successful login.

#### Request:

```json
{
  "username": "string",
  "password": "string"
}
```

#### Responses:

* 200 OK:
* Description: User logged in successfully.

Body:
```json
{
  "message": "User logged in successfully",
  "token": "jwt-token-here"
}
```

## Promote user

```
PATCH localhost:8080/login
```

Promotes a user to admin status. Only accessible by users with an admin token.

#### Request:
  * Headers:
      * Content-Type: application/json
      * Authorization: Bearer <admin-token>
  * Query Parameters:
      * username: The username of the user to promote.

Example Request:

```bash
PATCH /promote?username=johndoe
Authorization: Bearer <admin-token>
```

#### Responses:

* 200 OK:
* Description: User promoted to admin successfully.

Body:
```json
{
  "message": "User promoted to admin successfully"
}
```

## GetAllTasks

```GET localhost:8080/tasks```

This endpoint makes an HTTP GET request to localhost:8080/tasks to retrieve a list of tasks. The request does not include a request body. The response will have a status code of 200 and a content type of application/json. The response body will be an array of task objects, each containing an id, title, description, due date, and status. 

* The header should include a proper authorization bearer token - only a registered user can get tasks

Here's an example of the response body:

```json
[
  {
    "id":"uuid",
    "title":"string",
    "description":"string",
    "due_date":"string (ISO 8601 format)",
    "status":"string"
  }
]
```

### Example

#### Request

```curl
curl --location 'localhost:8080/tasks'
```

#### Response

```JSON
[
    {
        "id": "9e484920-0871-49a3-9bcf-2b9a29e7ec09",
        "title": "Task 1",
        "description": "First task",
        "due_date": "2024-08-06T14:40:10.331133+03:00",
        "status": "Pending"
    },
    {
        "id": "9e484920-0871-49a3-9bcf-2b9a29e7ec09",
        "title": "Task 2",
        "description": "Second task",
        "due_date": "2024-08-07T14:40:10.331133+03:00",
        "status": "In Progress"
    }
]
```

## GET - GetTaskByID

```localhost:8080/tasks/:id```

This endpoint retrieves a specific task by its ID. The ID specified as a parameter

* The header should include a proper authorization bearer token - only a registered user can get task

#### Request
* Method: GET
* Endpoint: localhost:8080/tasks/:id

#### Response
* Status: 200
* Content-Type: application/json
* { "id": "uuid", "title": "string", "description": "string", "due_date": "string  (ISO 8601 format)", "status": "string"}

#### Example Response

```JSON
{
    "id":"uuid",
    "title":"string",
    "description":"string",
    "due_date":"string (ISO 8601 format)",
    "status":"string"
}
```


## PUT - UpdateTaskByID

```localhost:8080/tasks/:id```

This endpoint allows you to update a specific task identified by its ID. The request should be sent to localhost:8080/tasks/:id using the HTTP PUT method.

* The header should include a proper authorization admin bearer token - only a registered admin can get update tasks

#### Request Body
The request body should be in raw format and include the following parameters:
* Title (string): The updated title of the task.
* description (string): The updated description of the task.
* status (string): The updated status of the task.
* due_date (string (ISO 8601 format)): The updated due date of the task.

#### Response
Upon successful execution, the endpoint returns a status code of 201 and a JSON object with the updated details of the task, including the following properties:
* id (string): The ID of the task.
* title (string): The title of the task.
* description (string): The description of the task.
* due_date (string): The due date of the task.
* status (string (ISO 8601 format)): The status of the task.

#### Example Response
```JSON
{
    "id":"uuid",
    "title":"string",
    "description":"string",
    "due_date":"string (ISO 8601 format)",
    "status":"string"
}
```

## DELETE - DeleteTask

```localhost:8080/tasks/:id```

This endpoint is used to delete a specific task identified by its ID. 

* The header should include a proper authorization admin bearer token - only a registered admin can get delete tasks

#### Request

* Method: DELETE
* URL: localhost:8080/tasks/:id

#### Response

* Status: 200


## POST - AddTask

```localhost:8080/tasks```

This endpoint is used to create a new task.

* The header should include a proper authorization admin bearer token - only a registered admin can get add tasks

#### Request Body

* title (string, required): The title of the task.
* description (string, required): The description of the task.
* due_date (string, required): The due date of the task.
* status (string, required): The status of the task.

#### Response

The response is in JSON format with the following schema:
```JSON
{
  "type": "object",
  "properties": {
    "id": {
      "type": "uuid"
    },
    "title": {
      "type": "string"
    },
    "description": {
      "type": "string"
    },
    "due_date": {
      "type": "string"
    },
    "status": {
      "type": "string"
    }
  }
}
```


### Error Handling:
Each endpoint returns error messages in a standardized format, with appropriate HTTP status codes depending on the error encountered. It's important to handle these errors gracefully on the client side.