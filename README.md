# A lightweight RESTful API built with Go for managing student records. This project serves as a practical implementation of RESTful principles using Go's standard libraries.


### API Reference

#### Create Student
http
```
POST /api/students
Content-Type: application/json
{
    "name": "John Doe",
    "age": 20,
    "grade": "A"
}
```

Response
```
{
    "id": "1",
    "name": "John Doe",
    "age": 20,
    "grade": "A",
    "created_at": "2024-02-01T10:00:00Z"
}
```

#### Get Student by ID
http
```
GET /api/students/{id}
```

Response
json
```
{
    "id": "1",
    "name": "John Doe",
    "age": 20,
    "grade": "A",
    "created_at": "2024-02-01T10:00:00Z"
}
```

#### List All Students
http
```
Get /api/students/
```

Response
json
```
[
    {
        "id": "1",
        "name": "John Doe",
        "age": 20,
        "grade": "A",
        "created_at": "2024-02-01T10:00:00Z"
    },
    {
        "id": "2",
        "name": "Jane Smith",
        "age": 22,
        "grade": "B",
        "created_at": "2024-02-01T11:00:00Z"
    }
]
```
