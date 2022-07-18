# ENV
```
DB_NAME = ./animal
ENVIRONMENT = dev
```
# DB
```sqlite3
CREATE TABLE IF NOT EXISTS animals (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   name TEXT NOT NULL UNIQUE,
   class TEXT NOT NULL,
   legs INT NOT NULL
);
```
# How to use
1. `go run main.go`
2. access in `http://localhost:8080/api/v1/`

# Endpoint
## /api/v1/get
query
```
page: [int]
limit: [int]
```
## /api/v1/get/:id
## /api/v1/delete
body
```
{
  "id": int,
  "name": string,
  "class": string,
  "legs": int,
}
```
## /api/v1/add
body
```
{
  "name": string,
  "class": string,
  "legs": int,
}
```
## /api/v1/update
body
```
{
  "id": int,
  "name": string,
  "class": string,
  "legs": int,
}
```
