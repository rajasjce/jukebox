# Jukebox API 🎶

Jukebox is a RESTful API application for managing music albums and musicians. The API allows users to create, update, retrieve, and delete albums and musicians, as well as link albums to musicians. This project is implemented using Go, with Gorilla Mux for routing, and uses an SQLite database.
## Features
- **Albums**: Create, update, delete, and list albums.
- **Musicians**: Manage musicians and associate them with albums.
- **RESTful Endpoints**: Fully functional API to interact with the music catalog.

## Getting Started

### Prerequisites
- **Go** 1.18 or later
- **SQLite** 

### Setup

1. **Clone the Repo**:
   ```bash
   git clone https://github.com/rajasjce/jukebox.git
   cd jukebox

2. **Run the Server**:

   ```go run main.go```
   The server will run at http://localhost:8080. You can access this URL to check if the server is running.
   
4. **Install Dependencies**

   Use the Go tool to install required dependencies:
```go mod tidy```

5. **Set Up the Database**

   Ensure the database is properly set up. A sample database script is included in the repository to create the necessary tables.

   Run the following command to create the database and tables:
   ```sqlite3 jukebox.db < database/schema.sql```


## API Endpoints

- **Albums**:
  - `GET /albums` - Retrieve the list of music albums sorted by the date of release in ascending order (i.e., oldest first).
  - `POST /albums` - Create a new music album.
  - `PUT /albums/{id}` - Update an existing music album by ID.
  - `DELETE /albums/{id}` - Delete a music album by ID.
  - `GET /musicians/{id}/albums` - Retrieve the list of music albums for a specified musician sorted by price in ascending order (i.e., lowest first).

- **Musicians**:
  - `GET /musicians` - Retrieve all musician records.
  - `POST /musicians` - Create a new musician.
  - `PUT /musicians/{id}` - Update an existing musician by ID.
  - `DELETE /musicians/{id}` - Delete a musician by ID.
  - `GET /albums/{id}/musicians` - Retrieve the list of musicians for a specified music album sorted by musician's name in ascending order.


