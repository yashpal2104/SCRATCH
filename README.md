# RSS Aggregator

## Overview
This project is an RSS Aggregator that allows users to create and manage RSS feeds. It provides several endpoints for creating users, creating feeds, and fetching feeds.

## Prerequisites
- Go 1.23.5 or later
- PostgreSQL

## Installation
1. Clone the repository:
    ```sh
    git clone https://github.com/yashpal2104/SCRATCH.git
    cd SCRATCH
    ```

2. Set up the environment variables by creating a `.env` file in the root directory:
    ```sh
    touch .env
    ```

    Add your PostgreSQL connection string and other required variables to the `.env` file.

3. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

### Running the Server
To run the server, use the following command:
```sh
go run main.go
```
Endpoints
User Endpoints
Create User

HTTP
POST /users
Request Body:

JSON
{
  "name": "username"
}
Response:

JSON
{
  "id": "user-id",
  "name": "username",
  "created_at": "timestamp",
  "updated_at": "timestamp"
}
Get User

HTTP
GET /users/{id}
Response:

JSON
{
  "id": "user-id",
  "name": "username",
  "created_at": "timestamp",
  "updated_at": "timestamp"
}
Feed Endpoints
Create Feed

HTTP
POST /feeds
Request Body:

JSON
{
  "name": "feed name",
  "url": "feed url"
}
Response:

JSON
{
  "id": "feed-id",
  "name": "feed name",
  "url": "feed url",
  "created_at": "timestamp",
  "updated_at": "timestamp",
  "user_id": "user-id"
}
Get Feeds

HTTP
GET /feeds
Response:

JSON
[
  {
    "id": "feed-id",
    "name": "feed name",
    "url": "feed url",
    "created_at": "timestamp",
    "updated_at": "timestamp",
    "user_id": "user-id"
  }
]
Project Structure
cmd/ - Entry points for the application.
internal/ - Internal packages.
auth/ - Authentication related code.
database/ - Database interactions.
sql/ - SQL queries and schema.
queries/ - SQL query files.
schema/ - Database schema files.
Contributing
Fork the repository.
Create a new branch (git checkout -b feature-branch).
Commit your changes (git commit -am 'Add new feature').
Push to the branch (git push origin feature-branch).
Create a new Pull Request.
License
This project is licensed under the MIT License.

Code

You can customize further as per your project's specifics and additional functionalities.
