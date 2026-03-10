---
# 🏗 Task Management API: Clean Architecture Refactor

This project is a complete refactor of the Task Management API, migrating from a traditional layered architecture to **Clean Architecture** (also known as Onion or Hexagonal Architecture).

The primary goal of this refactor is to achieve **Separation of Concerns**, **Dependency Inversion**, and **High Testability** by ensuring that core business logic remains entirely isolated from external frameworks, databases, and delivery mechanisms.

---

## 📐 Architectural Overview

The application is strictly divided into functional layers, with dependencies only pointing inward toward the Domain layer.

### Folder Structure & Responsibilities

```text
task-manager/
├── Delivery/           # Layer 4: External HTTP/Routing
│   ├── main.go         # Composition Root (Dependency Injector)
│   ├── controllers/    # Parses HTTP requests/JSON and calls Usecases
│   └── routers/        # Gin Engine setup and route definitions
├── Domain/             # Layer 1: The Core
│   └── domain.go       # Core Structs (Task, User) and Interface Contracts
├── Infrastructure/     # Layer 4: External Tools & Services
│   ├── auth_middleware.go # JWT request interception
│   ├── jwt_service.go     # Token generation and validation logic
│   └── password_service.go# Bcrypt hashing implementation
├── Repositories/       # Layer 3: Data Access
│   ├── task_repository.go # MongoDB specific queries and BSON mapping for Tasks
│   └── user_repository.go # MongoDB specific queries and BSON mapping for Users
└── Usecases/           # Layer 2: Business Logic
    ├── task_usecases.go   # Rules for creating/modifying tasks
    └── user_usecases.go   # Rules for registration and authentication
```

---

## 🔑 Key Clean Architecture Principles Applied

### 1. Dependency Inversion
The Usecase layer does not import the Repository or Infrastructure layers. Instead, the `Domain` defines **Interfaces** (e.g., `TaskRepository`, `PasswordService`). The outer layers implement these interfaces, and they are injected into the Usecases at runtime. 
* *Benefit:* The database (MongoDB) or hashing algorithm (Bcrypt) can be swapped out without altering a single line of business logic.

### 2. Domain Isolation (The Mapping Pattern)
The core `Domain.Task` and `Domain.User` models contain no database-specific tags (like `bson:"_id"`). 
* *Benefit:* The core logic relies on standard Go types (like `string` for IDs). The `Repositories` layer handles the translation between Go types and MongoDB's `primitive.ObjectID`.

### 3. Infrastructure Abstraction
Security functions like Password Hashing and JWT Generation are abstracted into interfaces within the Domain.
* *Benefit:* The `UserUsecase` manages the *flow* of authentication (e.g., "Hash the password before saving") without knowing the mathematical implementation of the hash.

---

## 🚀 How to Run the Application

### Prerequisites
*   Go 1.20+
*   MongoDB running locally on default port `27017`

### Startup
1. Clone the repository and navigate to the project root.
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the application from the Delivery folder:
   ```bash
   go run Delivery/main.go
   ```

The server will start on `http://localhost:8080`.

---

## 🧪 Testing

Because the architecture relies heavily on Dependency Injection, the business logic (`Usecases`) can be tested instantly using Mock implementations of the `Repositories` and `Infrastructure` interfaces, without requiring a live MongoDB connection.

*(To run tests, navigate to the Usecases directory and execute `go test -v`)*

---

## 🛣 API Endpoints

*(Note: All `/tasks` endpoints require a valid JWT Bearer Token in the Authorization header)*

### Authentication
*   `POST /register` - Register a new user (Requires `username`, `password`, optional `role`).
*   `POST /login` - Authenticate and receive a JWT.

### Task Management
*   `GET /tasks` - Retrieve all tasks.
*   `GET /tasks/:id` - Retrieve a specific task.
*   `POST /tasks` - Create a new task.
*   `PUT /tasks/:id` - Update an existing task.
*   **`DELETE /tasks/:id`** - Delete a task *(Requires `admin` role)*.