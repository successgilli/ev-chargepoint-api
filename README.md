# EV Charge Point API

A simple RESTful API server in Go for managing and querying a database of Electric Vehicle (EV) charge points.

## Overview

This service exposes endpoints to create, retrieve, update, and delete EV charge point records, backed by a SQLite database.

## Tech Stack

- Go 1.26
- Gin
- SQLite
- Follow golang standard project layout design https://github.com/golang-standards/project-layouthttps://github.com/golang-standards/project-layout

## Getting Started

### Prerequisites

- Go 1.26 installed

### Running the server

```bash
go run ./cmd/server
```

### Generating swagger docs
```sh
go tool swag init -g cmd/server/main.go -o api --parseDependency --parseInternal
```

### View swagger docs
- Generate docs
- Visit configured path [http://127.0.0.1:8080/api/v1/swagger/index.html](http://127.0.0.1:8080/api/v1/swagger/index.html)

### Generating mocks for test
This repo uses mockery for generating mocks as part of its tools.
Run the below to regenerate mocks after updates to signatures.
```sh
go tool mockery
```

### Running tests
```bash
go test ./...
```

## API Endpoints

| Method | Endpoint              | Description                  |
|--------|-----------------------|-------------------------------|
| GET    | /chargepoints         | List all charge points        |
| GET    | /chargepoints/{id}    | Get a single charge point     |
| POST   | /chargepoints         | Create a new charge point     |

## Time Spent

Approximately [X] hours.

## LLM Usage
TODO: Mention LLM usage as it relates to this work.

## TODO
- Add remaining two endpoints
- Add shutdown handling
- Add catchall error route/middleware
- Add clean and add docs for itnerfaces and structs
- Check validation for lat/long. 0 is rejected and might be valid.
