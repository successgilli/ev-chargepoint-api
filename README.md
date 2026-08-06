# EV Charge Point API

A simple RESTful API server in Go for managing and querying a database of Electric Vehicle (EV) charge points.

## Overview

This service exposes endpoints to create, retrieve, update, and delete EV charge point records, backed by a SQLite database.

## Tech Stack

- Go 1.26
- Gin
- SQLite

## Getting Started

### Prerequisites

- Go 1.26 installed

### Running the server

```bash
go run main.go
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
| PUT    | /chargepoints/{id}    | Update a charge point         |
| DELETE | /chargepoints/{id}    | Delete a charge point         |

## Time Spent

Approximately [X] hours.

## LLM Usage
TODO: Mention LLM usage as it relates to this work.
