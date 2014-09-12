## API

**GET /todos**

Returns all todos.

**Example Response**

```json
[{"id":"faf24dae-ce0d-4b93-9efc-1051e1680366","text":"Hello","completed_at":null}]
```

**POST /todos**

Creates a new todo.

**Example Response**

```json
{"id":"faf24dae-ce0d-4b93-9efc-1051e1680366","text":"Hello","completed_at":null}
```

**DELETE /todos/:id**

Delete a todo.

**Example Response**

```json
{"id":"faf24dae-ce0d-4b93-9efc-1051e1680366","text":"Hello","completed_at":null}
```

**POST /todos/:id/complete**

Mark the todo is complete.

```json
{"id":"faf24dae-ce0d-4b93-9efc-1051e1680366","text":"Hello","completed_at":"2014-09-12T10:31:38.310457446-07:00"}
```

**DELETE /todos/:id/complete**

Mark the todo as not complete.

```json
{"id":"faf24dae-ce0d-4b93-9efc-1051e1680366","text":"Hello","completed_at":null}
```
