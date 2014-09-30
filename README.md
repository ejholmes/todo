# Todo

Todo is a simple Go backend for storing Todo's, running at http://r101-todo.herokuapp.com/todos

## API

### GET /todos

Returns all todos.

**Example Request**

```console
curl http://r101-todo.herokuapp.com/todos
```

**Example Response**

```json
[
  {
    "id": "6a84b9f1-8acf-4e37-a29b-e11fa115acdc",
    "text": "Take out the trash",
    "completed_at": "2014-09-30T17:10:11.272776746Z"
  }
]
```

### POST /todos

Creates a new todo.

**Example Request**

```console
curl -d '{"text": "Take out the trash"}' http://r101-todo.herokuapp.com/todos
```

**Example Response**

```json
{
  "id": "6a84b9f1-8acf-4e37-a29b-e11fa115acdc",
  "text": "Take out the trash",
  "completed_at": null
}
```

### DELETE /todos/:id

Delete a todo.

**Example Request**

```console
curl -X DELETE http://r101-todo.herokuapp.com/todos/6a84b9f1-8acf-4e37-a29b-e11fa115acdc
```

**Example Response**

```json
{
  "id": "6a84b9f1-8acf-4e37-a29b-e11fa115acdc",
  "text": "Take out the trash",
  "completed_at": null
}
```

### PATCH /todos/:id

Update a todo

**Example Request**

```console
curl -X PATCH -d '{"text":"Walk the dog"}' http://r101-todo.herokuapp.com/todos/6a84b9f1-8acf-4e37-a29b-e11fa115acdc
```

**Example Response**

```json
{
  "id": "6a84b9f1-8acf-4e37-a29b-e11fa115acdc",
  "text": "Walk the dog",
  "completed_at": null
}
```

### POST /todos/:id/complete

Mark the todo is complete.

**Example Request**

```console
curl -X POST http://r101-todo.herokuapp.com/todos/6a84b9f1-8acf-4e37-a29b-e11fa115acdc/complete
```

**Example Response**

```json
{
  "id": "6a84b9f1-8acf-4e37-a29b-e11fa115acdc",
  "text": "Take out the trash",
  "completed_at": "2014-01-01T00:00:00.000000000Z"
}
```

### DELETE /todos/:id/complete

Mark the todo as not complete.

**Example Request**

```console
curl -X DELETE http://r101-todo.herokuapp.com/todos/6a84b9f1-8acf-4e37-a29b-e11fa115acdc/complete
```

**Example Response**

```json
{
  "id": "6a84b9f1-8acf-4e37-a29b-e11fa115acdc",
  "text": "Take out the trash",
  "completed_at": null
}
```
