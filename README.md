# AnimeList

## Prerequisites

Before you begin, ensure you have the following:

- PostgreSQL installed on your system. If not, download and install it from the [official PostgreSQL website](https://www.postgresql.org/).
- Golang installed on your system. If not, download and install it from the [official Golang website](https://go.dev/doc/install).

## Steps to Create a PostgreSQL Database

1. **Open Command Line or Terminal**: Open your command line or terminal application.

2. **Connect to PostgreSQL**: Use the following command to connect to PostgreSQL:

   ```bash
   psql -U postgres

   CREATE DATABASE my_database;


## Setup Instructions

1. Clone the repository
2. Install dependencies: `go mod tidy`
3. Create your .env file based on .env.example
4. Start the server: `go run main.go`

## API Endpoint

- **POST /auth/sign-in**
- **POST /auth/sign-up**
- **GET /anime/:id**
- **PUT /anime/:id**
- **DELETE /anime/:id**
- **GET /anime**
- **POST /anime**

