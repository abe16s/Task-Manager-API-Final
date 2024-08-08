# Task Management Service

A task management service built with Go and MongoDB, providing functionality for adding, updating, deleting, and retrieving tasks. This service offers a RESTful API for managing tasks and demonstrates the use of MongoDB for data storage.

## Configuration

Before running the application, ensure you have a MongoDB instance running. Update the database connection string in main.go where it is set with clientOptions := options.Client().ApplyURI("").

## GET - GetAllTasks

localhost:8080/tasks

This endpoint makes an HTTP GET request to localhost:8080/tasks to retrieve a list of tasks. The request does not include a request body. The response will have a status code of 200 and a content type of application/json. The response body will be an array of task objects, each containing an id, title, description, due date, and status. Here's an example of the response body:


```json
[{"id":"","title":"","description":"","due_date":"","status":""}]
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

localhost:8080/tasks/:id

This endpoint retrieves a specific task by its ID.

#### Request
* Method: GET
* Endpoint: localhost:8080/tasks/:id

#### Response
* Status: 200
* Content-Type: application/json
* { "id": "", "title": "", "description": "", "due_date": "", "status": ""}

#### Example Response

```JSON
{
    "id": "",
    "title": "",
    "description": "",
    "due_date": "",
    "status": ""
}
```


## PUT UpdateTaskByID

localhost:8080/tasks/:id

This endpoint allows you to update a specific task identified by its ID. The request should be sent to localhost:8080/tasks/:id using the HTTP PUT method.

#### Request Body
The request body should be in raw format and include the following parameters:
* Title (string): The updated title of the task.
* description (string): The updated description of the task.
* status (string): The updated status of the task.
* due_date (string): The updated due date of the task.

#### Response
Upon successful execution, the endpoint returns a status code of 201 and a JSON object with the updated details of the task, including the following properties:
* id (string): The ID of the task.
* title (string): The title of the task.
* description (string): The description of the task.
* due_date (string): The due date of the task.
* status (string): The status of the task.

#### Example Response
```JSON
{
    "id": "",
    "title": "",
    "description": "",
    "due_date": "",
    "status": ""
}
```

## DELETE DeleteTask

localhost:8080/tasks/:id

This endpoint is used to delete a specific task identified by its ID.

#### Request

* Method: DELETE
* URL: localhost:8080/tasks/:id

#### Response

* Status: 200


## POST AddTask

localhost:8080/tasks


This endpoint is used to create a new task.

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