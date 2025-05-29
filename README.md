# GO API Workshop (Based on BorntoDev Free Course)

This workshop was created as a practice project using the Go programming language.

I have prior experience with web development in JavaScript and TypeScript, which made it a bit easier to follow the course content.

Thanks to [BorntoDev](https://www.youtube.com/watch?v=fjEB75Xotxc) for providing such great and useful content for free!

## Scope of the Workshop

- Build an API and connect it to a MySQL database
- Retrieve all data from the database using the `GET` method
- Retrieve data by specific ID using `GET /:id`
- Use middleware and enable CORS

## Project Overview

This project is a simple API to manage `Worker` records, using `SQLite3` as the database to store information:

```go
type Worker struct {
    Id      int
    Name    string
    Email   string
    Phone   string
    Salary  int
}
````
## Getting Started

### Clone the repository

```bash
git clone https://github.com/NuttapolKhumdang/workshop-basic-go-api.git
cd borntodev-go-api-workshop
```

### Run the API server

```bash
go run .
```

The server will be running at: `http://localhost:8080`

---

## Project Structure

```txt
workshop-basic-go-api/
│
├── main.go          # Entry point of the application
├── handlers.go      # Functions to handle API requests
├── middleware.go    # Middleware for logging and CORS
├── logger.go        # Logger utility
│
├── web/             # Simple frontend to interact with the API
│   ├── index.html
│   └── app.js
│
└── test/            # Python scripts for API testing
    ├── __main__.py
    ├── handle.py
    └── body.json
```

## API Endpoints

| Method | Endpoint       | Description          | Request Body                             | Response |
| ------ | -------------- | -------------------- | ---------------------------------------- | -------- |
| `GET`  | `/workers`     | Fetch all workers    | -                                        | JSON     |
| `GET`  | `/workers/:id` | Fetch a worker by ID | -                                        | JSON     |
| `POST` | `/workers`     | Add a new worker     | `Id`, `Name`, `Email`, `Phone`, `Salary` | Status   |

---

## Middleware and CORS

This project uses middleware to log each incoming request's HTTP method and path. It also sets CORS headers:

```http
Access-Control-Allow-Origin: *
Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS
Access-Control-Allow-Headers: Accept, Content-Type, Content-Length
```

---

## Reflection

I spent 3–4 days completing this course and gained a lot of knowledge in the process.

From this project, I learned the fundamentals of the Go language, how to structure a basic web API, and how to interact with databases and middleware.

This was my first Go project — from zero to a working API!
