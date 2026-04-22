# task-api

Built to learn Go fundamentals and deepen understanding of REST API design. No frameworks — just `net/http`, `encoding/json`, and in-memory storage.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/tasks` | List all tasks |
| POST | `/tasks` | Create a task |
| GET | `/tasks/{id}` | Get a task |
| PUT | `/tasks/{id}` | Update a task |
| DELETE | `/tasks/{id}` | Delete a task |

## Run

```
go run main.go
```

## Example

```
curl -X POST http://localhost:8000/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Learn Go"}'
```
