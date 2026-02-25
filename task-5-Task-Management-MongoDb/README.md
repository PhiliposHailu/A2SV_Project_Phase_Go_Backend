
---

# Task Management API (MongoDB Edition)

This is a RESTful API for a Task Management system built with Go, the Gin Web Framework, and MongoDB for persistent data storage. It supports full CRUD operations using JSON payloads.

## Prerequisites & Database Setup

To run this API, you must have MongoDB installed and running on your host machine.

1. **Install MongoDB:** Download MongoDB Community Server from the [official website](https://www.mongodb.com/try/download/community) or install via your system's package manager (e.g., `apt` for Debian/Ubuntu based systems).
2. **Start the Database Service:** Ensure the MongoDB service is running on the default port `27017`.
   ```bash
   sudo systemctl start mongod
   ```
3. **Run the API:**
   ```bash
   go mod tidy
   go run main.go
   ```
*Note: The API will automatically create the database (`taskdb`) and collection (`tasks`) upon the first insertion.*

---

## API Endpoints

**Base URL:** `http://localhost:3000`

### 1. Get All Tasks
Retrieves a list of all tasks currently stored in the system.
* **Method:** `GET`
* **URL:** `/tasks`
* **Success Response (200 OK):**
  ```json
  {
    "message": [
      {
        "id": "65d4a1b2c3d4e5f67890abcd",
        "title": "Learn Go",
        "description": "Master Gin framework",
        "due_date": "2025-02-25",
        "status": "In Progress"
      }
    ]
  }
  ```

### 2. Get Task by ID
Retrieves the details of a single task based on its unique ObjectID.
* **Method:** `GET`
* **URL:** `/tasks/:id`
* **URL Params:** `id=[string]`
* **Success Response (200 OK):**
  ```json
  {
    "message": {
      "id": "65d4a1b2c3d4e5f67890abcd",
      "title": "Learn Go",
      "description": "Master Gin framework",
      "due_date": "2025-02-25",
      "status": "In Progress"
    }
  }
  ```
* **Error Response (404 Not Found):** `{"message": "mongo: no documents in result"}`
* **Error Response (400 Bad Request):** `{"message": "invalid ID format"}`

### 3. Create a New Task
Adds a new task to the database. The `id` will be generated automatically by MongoDB.
* **Method:** `POST`
* **URL:** `/tasks`
* **Request Body:**
  ```json
  {
    "title": "Buy Milk",
    "description": "Get almond milk from the store",
    "due_date": "2025-02-20",
    "status": "Pending"
  }
  ```
* **Success Response (201 Created):** `{"message": { ...task details with new ID... }}`
* **Error Response (400 Bad Request):** `{"message": "Error with given task"}`

### 4. Update an Existing Task
Updates specific fields of an existing task identified by its ID. Only the provided fields will be modified.
* **Method:** `PUT`
* **URL:** `/tasks/:id`
* **URL Params:** `id=[string]`
* **Request Body:**
  ```json
  {
    "status": "Completed"
  }
  ```
* **Success Response (200 OK):** `{"message": "updated successfully."}`
* **Error Response (404 Not Found):** `{"message": "mongo: no documents in result"}`

### 5. Delete a Task
Removes a task from the system permanently.
* **Method:** `DELETE`
* **URL:** `/tasks/:id`
* **URL Params:** `id=[string]`
* **Success Response (200 OK):** `{"message": "Deleted."}`
* **Error Response (404 Not Found):** `{"message": "mongo: no documents in result"}`

---

## Status Codes Reference

| Status Code | Description |
| :--- | :--- |
| `200 OK` | The request was successful. |
| `201 Created` | The task was successfully created. |
| `400 Bad Request` | The request payload was invalid or missing required fields. |
| `404 Not Found` | The requested endpoint or resource (Task ID) does not exist in the database. |
| `500 Internal Server Error` | An unexpected error occurred on the server or database. |

---
