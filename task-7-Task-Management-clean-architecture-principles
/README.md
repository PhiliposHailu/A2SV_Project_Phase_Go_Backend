# Task Management API: Authentication & Authorization

This API is secured using **JWT (JSON Web Tokens)** and **BCRYPT** password hashing. Access to task resources is restricted based on user authentication and specific roles (Admin vs. Regular User).

## Security Features

- **Password Hashing:** All user passwords are encrypted using the `bcrypt` algorithm before being stored in MongoDB.
- **Stateless Authentication:** Uses JWT to verify user identity without storing sessions on the server.
- **Role-Based Access Control (RBAC):** Restricts high-privilege actions (like deleting tasks) to Admin users only.

---

## 🔑 Authentication Endpoints

### 1. User Registration

Creates a new user account.

- **URL:** `/register`
- **Method:** `POST`
- **Payload:**
  ```json
  {
    "username": "johndoe",
    "password": "securepassword123",
    "role": "user"
  }
  ```
- **Note:** Role defaults to `"user"` if not specified. Use `"admin"` to create an administrative account.

### 2. User Login

Authenticates credentials and returns a JWT access token.

- **URL:** `/login`
- **Method:** `POST`
- **Payload:**
  ```json
  {
    "username": "johndoe",
    "password": "securepassword123"
  }
  ```
- **Success Response (200 OK):**
  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "message": "Login successful"
  }
  ```

---

## 🛡 Protected Task Endpoints

All endpoints below require the following header:
`Authorization: Bearer <YOUR_JWT_TOKEN>`

| Endpoint     | Method   | Required Role  | Description                |
| :----------- | :------- | :------------- | :------------------------- |
| `/tasks`     | `GET`    | User / Admin   | View all tasks.            |
| `/tasks/:id` | `GET`    | User / Admin   | View a specific task.      |
| `/tasks`     | `POST`   | User / Admin   | Create a new task.         |
| `/tasks/:id` | `PUT`    | User / Admin   | Update a task.             |
| `/tasks/:id` | `DELETE` | **Admin Only** | Permanently remove a task. |

---

## ⚙️ How to use Protected Routes in Postman

1.  **Login:** Send a POST request to `/login` and copy the `token` value from the response.
2.  **Authorize:** In your Task request (e.g., `GET /tasks`), go to the **Authorization** tab.
3.  **Type:** Select **Bearer Token**.
4.  **Token:** Paste the copied JWT.
5.  **Send:** If the token is valid, you will receive the data. If missing or expired, you will receive `401 Unauthorized`. If a regular user tries to delete a task, they will receive `403 Forbidden`.

---

## 📁 Project Architecture

- `middleware/`: Contains the JWT validation bouncer.
- `utils/`: Contains JWT generation and verification logic.
- `data/`: Handles MongoDB operations and password hashing.
- `controllers/`: Orchestrates the request/response flow.

---
