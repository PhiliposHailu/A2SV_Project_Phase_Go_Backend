
---

# Unit Testing Documentation

This project implements comprehensive unit testing for the application's core business logic (the Usecase layer). By strictly adhering to Clean Architecture principles, the business logic is entirely decoupled from external dependencies such as the database (MongoDB) and the web framework (Gin).

## Testing Strategy

*   **Isolation via Dependency Injection:** External dependencies (`UserRepository`, `TaskRepository`, `PasswordService`, `JWTService`) are abstracted using Go interfaces.
*   **Mocking:** We utilize the `github.com/stretchr/testify/mock` library to generate mock implementations of these interfaces. This allows us to test the Usecase logic without requiring a live database connection.
*   **Table-Driven Tests:** Tests are written using the Table-Driven pattern, which groups multiple test scenarios (Success, Failure, Edge Cases) into a single, highly readable and easily scalable test function.

## Covered Components & Scenarios

### 1. Task Usecase (`usecase/task_usecase_test.go`)
Tests the business rules surrounding task creation and management.
*   **Create Task - Success:** Verifies that a valid task is successfully passed to the repository.
*   **Create Task - Empty Title (Failure):** Verifies that the business rule blocking tasks with empty titles functions correctly, ensuring the repository is never called.
*   **Create Task - Database Error (Failure):** Verifies that repository-level errors are properly propagated back to the caller.

### 2. User Usecase (`usecase/user_usecase_test.go`)
Tests the security and authentication rules.
*   **Registration - Success:** Verifies that a new username is checked for uniqueness, the password is encrypted via the PasswordService, and the data is saved.
*   **Registration - Username Exists (Failure):** Ensures the system rejects duplicate usernames before attempting to hash passwords or save data.
*   **Registration - Empty Fields (Failure):** Validates input constraints.
*   **Login - Success:** Verifies that existing users are found, passwords match the hashed value, and a JWT token is successfully generated.
*   **Login - User Not Found (Failure):** Verifies generic error response to prevent user enumeration.
*   **Login - Invalid Password (Failure):** Verifies rejection of incorrect credentials.

## How to Run the Tests

To execute the unit test suite, open a terminal in the root directory of the project and run the following command:

```bash
go test ./usecase/... -v
```

**Command Flags:**
*   `./usecase/...`: Directs the Go testing tool to run all tests inside the `usecase` directory and its subdirectories.
*   `-v`: Enables verbose output, displaying the exact name and status (PASS/FAIL) of every individual scenario executed in the table-driven tests.