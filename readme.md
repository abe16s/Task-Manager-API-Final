# Task Manager Backend & API Project

This project is a backend and API service for a task management application, built with Go (Golang). The application allows users to manage tasks, with features such as user authentication, task creation, updating, and deletion. The project is organized into several packages and folders, each with a specific responsibility, ensuring a clean and modular architecture.

## Features

- **User Authentication**: Secure login and registration with JWT tokens.
- **Task Management**: Create, update, delete, and retrieve tasks.
- **Role-Based Access Control**: Restrict access to certain actions based on user roles.
- **RESTful API**: Provides a RESTful API for interacting with the task management system.
- **Continuous Integration**: Automated testing and build process via GitHub Actions.

## Folder Structure

The project is organized into the following structure:

```plaintext
task_manager
│   .env
│   go.mod
│   go.sum
│   readme.md
│
├───.github
│   └───workflows
│           go.yml
│
├───delivery
│   │   main.go
│   │
│   ├───controllers
│   │       task_controller.go
│   │       user_controller.go
│   │
│   └───router
│           router.go
│
├───docs
│       api_documentation.md
│
├───domain
│       domain.go
│
├───infrastructure
│       auth_middleware.go
│       jwt_services.go
│       password_service.go
│
├───repositories
│       task_repository.go
│       user_repository.go
│
├───tests
│   │   auth_middleware_test.go
│   │   jwt_services_test.go
│   │   password_service_test.go
│   │   task_controller_test.go
│   │   task_usecase_test.go
│   │   user_controller_test.go
│   │   user_usecase_test.go
│   │
│   ├───mocks
│   │       JwtServiceInterface.go
│   │       PasswordServiceInterface.go
│   │       TaskRepoInterface.go
│   │       TaskServiceInterface.go
│   │       UserRepoInterface.go
│   │       UserServiceInterface.go
│   │
│   └───repository_tests
│           task_repository_test.go
│           user_repository_test.go
│
└───usecases
        jwt_service_interface.go
        password_service_interface.go
        task_repository_interface.go
        task_usecase.go
        user_repository_interface.go
        user_usecase.go
```

### File/Folder Descriptions

- **.env**: Configuration file that stores environment variables, such as database connection strings, JWT secret keys, and other sensitive information.

- **go.mod**: Go module file that defines the module’s path and its dependencies.

- **go.sum**: File that contains the expected cryptographic checksums of the dependencies listed in `go.mod`.

- **readme.md**: The project’s README file containing details about the application, its setup, and usage.

- ### `.github/`
  - **workflows/go.yml**: GitHub Actions workflow file that defines the CI/CD pipeline for automated testing and builds.

- ### `delivery/`
  - **main.go**: The entry point of the application, responsible for initializing the server and setting up routes.
  
  - #### `delivery/controllers/`
    - **task_controller.go**: Handles HTTP requests related to tasks, such as creating, updating, and deleting tasks.
    - **user_controller.go**: Manages HTTP requests related to user actions, such as registration and authentication.
    
  - #### `delivery/router/`
    - **router.go**: Sets up the routing for the application, mapping HTTP routes to the corresponding controllers.

- ### `docs/`
  - **api_documentation.md**: Documentation file that provides details on the API endpoints, request/response formats, and other relevant information.

- ### `domain/`
  - **domain.go**: Contains domain models and entities used throughout the application, representing core business objects like `User` and `Task`.

- ### `infrastructure/`
  - **auth_middleware.go**: Implements middleware for handling authentication and authorization using JWT tokens.
  - **jwt_services.go**: Contains services for generating and validating JWT tokens.
  - **password_service.go**: Provides utilities for hashing and verifying passwords.

- ### `repositories/`
  - **task_repository.go**: Responsible for interacting with the database to perform CRUD operations on tasks.
  - **user_repository.go**: Handles database interactions related to users, such as retrieving user information and storing new users.

- ### `tests/`
  - **auth_middleware_test.go**: Tests for the authentication middleware.
  - **jwt_services_test.go**: Tests for JWT services.
  - **password_service_test.go**: Tests for the password hashing and verification service.
  - **task_controller_test.go**: Tests for the task controller.
  - **task_usecase_test.go**: Tests for task use cases.
  - **user_controller_test.go**: Tests for the user controller.
  - **user_usecase_test.go**: Tests for user use cases.
  
  - #### `tests/mocks/`
    - **JwtServiceInterface.go**: Mock implementation for JWT service interface.
    - **PasswordServiceInterface.go**: Mock implementation for password service interface.
    - **TaskRepoInterface.go**: Mock implementation for task repository interface.
    - **TaskServiceInterface.go**: Mock implementation for task service interface.
    - **UserRepoInterface.go**: Mock implementation for user repository interface.
    - **UserServiceInterface.go**: Mock implementation for user service interface.

  - #### `tests/repository_tests/`
    - **task_repository_test.go**: Unit tests for the task repository.
    - **user_repository_test.go**: Unit tests for the user repository.

- ### `usecases/`
  - **jwt_service_interface.go**: Defines the interface for the JWT service.
  - **password_service_interface.go**: Defines the interface for the password service.
  - **task_repository_interface.go**: Defines the interface for the task repository.
  - **task_usecase.go**: Contains the business logic for tasks, coordinating between the repository and controllers.
  - **user_repository_interface.go**: Defines the interface for the user repository.
  - **user_usecase.go**: Encapsulates the business logic related to user actions, such as registration and authentication.

## Clean Architecture

The project follows the principles of **Clean Architecture** to ensure a robust, maintainable, and scalable codebase. This architecture emphasizes separation of concerns, enabling independent development, testing, and maintenance of different parts of the system.

### Design Decisions

1. **Separation of Layers**:
   - **Domain Layer**: Contains the core business logic and entities (`domain/`). It is independent of external frameworks and libraries.
   - **Use Cases Layer**: Implements application-specific business rules (`usecases/`). It orchestrates the flow of data to and from the entities and is independent of the external world.
   - **Interface Adapters Layer**: Comprises controllers and gateways that convert data between the domain and external formats (`delivery/controllers/`, `repositories/`).
   - **Frameworks & Drivers Layer**: Includes databases, web frameworks, and other external tools (`infrastructure/`, `delivery/router/`).

2. **Dependency Rule**:
   - The code dependencies point inwards. Outer layers can depend on inner layers, but inner layers are unaware of outer layers. This ensures that the core business logic remains unaffected by changes in external components.

3. **Interfaces for Abstraction**:
   - Interfaces are defined for repositories and services to decouple implementations from their contracts. This promotes flexibility and easier testing.

4. **Middleware and Services**:
   - Authentication and other cross-cutting concerns are handled in the `infrastructure/` layer, ensuring they do not pollute the business logic.

### Guidelines for Future Development

1. **Adhere to Layer Boundaries**:
   - Maintain the separation of concerns by ensuring that each new component or feature fits within the appropriate layer.

2. **Implement Interfaces**:
   - For any new external service or repository, define an interface in the inner layers and implement it in the outer layers. This promotes decoupling and testability.

3. **Maintain the Dependency Rule**:
   - Ensure that dependencies always point inward. Avoid importing packages from outer layers into inner layers.

4. **Testing**:
   - Write unit tests for each layer separately. Use mock implementations for external dependencies to test inner layers.

5. **Documentation**:
   - Update the `docs/` with any new API endpoints, architectural changes, or design decisions to keep the documentation in sync with the codebase.

6. **Code Reviews**:
   - Conduct regular code reviews focusing on adherence to the Clean Architecture principles.

## Getting Started

### Prerequisites

- **Go**: Make sure you have Go installed on your machine.
- **Database**: Ensure you have a running database instance and update the `.env` file with your database credentials.

### Installation

1. **Clone the repository**:
   ```sh
   git clone https://github.com/abe16s/Task-Manager-Clean-Architecture.git
   cd task_manager
   ```

2. **Install dependencies**:
   ```sh
   go mod tidy
   ```

3. **Run the application**:
   ```sh
   go run delivery/main.go
   ```

## API Documentation

For detailed information on the available API endpoints, refer to the [API Documentation](docs/api_documentation.md).