---

# Task Management API Documentation

This API provides a set of endpoints to manage a collection of tasks. It supports full CRUD (Create, Read, Update, Delete) operations using JSON payloads.

**Base URL:** `http://localhost:3000`

---

## Endpoints

### 1. Get All Tasks
Retrieves a list of all tasks currently stored in the system.

*   **URL:** `/tasks`
*   **Method:** `GET`
*   **Success Response:**
    *   **Code:** `200 OK`
    *   **Content:**
        ```json
        {
          "message": [
            {
              "id": "1",
              "title": "Learn Go",
              "description": "Master Gin framework",
              "due_date": "2025-02-25",
              "status": "In Progress"
            }
          ]
        }
        ```

---

### 2. Get Task by ID
Retrieves the details of a single task based on its unique ID.

*   **URL:** `/tasks/:id`
*   **Method:** `GET`
*   **URL Params:** `id=[string]`
*   **Success Response:**
    *   **Code:** `200 OK`
    *   **Content:**
        ```json
        {
          "message": {
            "id": "1",
            "title": "Learn Go",
            "description": "Master Gin framework",
            "due_date": "2025-02-25",
            "status": "In Progress"
          }
        }
        ```
*   **Error Response:**
    *   **Code:** `404 NOT FOUND`
    *   **Content:** `{"message": "task not found"}`

---

### 3. Create a New Task
Adds a new task to the system.

*   **URL:** `/tasks`
*   **Method:** `POST`
*   **Request Body:**
    ```json
    {
      "id": "2",
      "title": "Buy Milk",
      "description": "Get almond milk from the store",
      "due_date": "2025-02-20",
      "status": "Pending"
    }
    ```
*   **Success Response:**
    *   **Code:** `201 CREATED`
    *   **Content:** `{"message": "Task created successfully."}`

---

### 4. Update an Existing Task
Updates the information of an existing task identified by its ID.

*   **URL:** `/tasks/:id`
*   **Method:** `PUT`
*   **URL Params:** `id=[string]`
*   **Request Body:**
    ```json
    {
      "title": "Buy Milk",
      "description": "Get almond milk and eggs",
      "due_date": "2025-02-21",
      "status": "In Progress"
    }
    ```
*   **Success Response:**
    *   **Code:** `200 OK`
    *   **Content:** `{"message": "updated successfully."}`
*   **Error Response:**
    *   **Code:** `400 BAD REQUEST`
    *   **Content:** `{"message": "task not found"}`

---

### 5. Delete a Task
Removes a task from the system permanently.

*   **URL:** `/tasks/:id`
*   **Method:** `DELETE`
*   **URL Params:** `id=[string]`
*   **Success Response:**
    *   **Code:** `200 OK`
    *   **Content:** `{"message": "deleted successfully."}`
*   **Error Response:**
    *   **Code:** `400 BAD REQUEST`
    *   **Content:** `{"message": "task not found"}`

---
