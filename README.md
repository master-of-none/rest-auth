# REST Auth

## Overview

`rest-auth` is a RESTful API built with Golang, using MongoDB for data storage, JWT for authentication, and custom middleware for enhanced security. This API provides a simple way to manage user authentication through registration and login functionalities.

## Features

-   User registration and login functionality
-   JWT-based authentication for secure access
-   MongoDB integration for data persistence

## Installation

1. **Clone the repository:**

    ```bash
    git clone https://github.com/yourusername/rest-auth.git
    cd rest-auth
    ```

2. **Install dependencies:**
    ```bash
    go mod tidy
    ```
3. **Set up your MongoDB database:**

    - Ensure that you have MongoDB installed and running.
    - Update the database connection details in the config/config.go file.

4. **Run the application:**

-   Start the application by executing:

    ```bash
    go run main.go
    ```

    A Better use is to check CompileDaemon to automatically compile and run the application. To do this install the compile daemon and follow it's documentation.

## Usage

Once the API is running, you can interact with the endpoints using tools like Postman or curl.

### API Endpoints

1.  **GET `/`**

    -   **Description:** A simple endpoint that returns a welcome message.
    -   **Response:**

        ```json
        `{
          "message": "Hello Rest World"
        }`
        ```

2.  **POST `/login`**

    -   **Description:** This endpoint allows users to log in by providing their credentials (username and password). It checks the provided credentials against the database and returns a JWT token if the credentials are valid.
    -   **Request Body:**

        ```json
        `{
          "username": "exampleuser",
          "password": "password123"
        }`
        ```

3.  **GET `/dbcheck`**

    -   **Description:** This endpoint checks the connection to the MongoDB database. It returns a success message if the connection is established.
    -   **Response:**

        ```json
        `{
            "message": "Database has been connected Successfully"
        }`
        ```

4.  **POST `/register`**

    -   **Description:** This endpoint allows new users to register by providing their username and password. The details are stored in the

    **Request Body:**

```json
{
    "username": "newuser",
    "password": "password123"
}
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.
