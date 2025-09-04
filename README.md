# GO RESTful API

A simple RESTful API developed with GO Language.

## Pre-Requisites

- [GO](o.dev)

## Run

Use script bellow to run server.

```bash
go run .
```

## Routes

**GET /account/coins/?username=**

```bash
curl -X GET http://localhost:8000/account/coins/\?username\=alex -H "Content-Type:application/json" -H "Authorization:123ABC"
```
