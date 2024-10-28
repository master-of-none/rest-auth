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
    - Make sure to set up a .env file to get the values needed to be loaded.
3. **Set up your MongoDB database:**

    - Ensure that you have MongoDB installed and running.
    - Update the database connection details in the .env file.
    - Set up a collection named counters in the database to store the user ID counter sequentially.

        ```bash
            use yourDatabaseName
        ```

        ```bash
            db.counters.insertOne({
                _id: "postid",
                sequence_value: 0
            })
        ```

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
        {
            "message": "Hello Rest World"
        }
        ```

2.  **POST `/login`**

    -   **Description:** This endpoint allows users to log in by providing their credentials (username or email and password). It checks the provided credentials against the database and returns a JWT token if the credentials are valid. User can provide email or username to login.
    -   **Request Body:**

        ```json
        {
            "username": "exampleuser",
            "email": "example@mail.com",
            "password": "password123"
        }
        ```

3.  **GET `/dbcheck`**

    -   **Description:** This endpoint checks the connection to the MongoDB database. It returns a success message if the connection is established.
    -   **Response:**

        ```json
        {
            "message": "Database has been connected Successfully"
        }
        ```

4.  **POST `/register`**

    -   **Description:** This endpoint allows new users to register by providing their username and password. The details are stored in the

        **Request Body:**

        ```json
        {
            "username": "newuser",
            "email": "example@mail.com",
            "password": "password123"
        }
        ```

5.  **GET `/protected/dashboard`**
    -   **Description:** This endpoint is protected by JWT authentication. It returns a success message if the user is authenticated.
    -   **Response:**
        ```json
        {
            "message": "Welcome to the dashboard"
        }
        ```
6.  **POST `/refreshToken`**

    -   **Description:** This endpoint allows users to refresh their JWT token set from the Cookie since the tokens are stored in Cookie.
        New access Token is generated from the refresh token which is not expired and it is set in the cookie. If refresh token is expired it will not generate new access token.
    -   **Response:**

        ```json
        {
            "username": "username",
            "message": "New Access token generated successfully",
            "newAccessToken": "newToken"
        }
        ```

7.  **DELETE /admin/delete/:id**
    -   **Description:** This endpoint allows the admin to delete the user by providing the user ID. The user is removed from the database.
    -   **Response:**
        ```json
        {
            "message": "User deleted successfully"
        }
        ```

### Post Routes

-   These endpoints are used to manage posts. All routes under /posts are protected by JWT authentication and require a valid token.

1.  **POST `/posts`**

    -   **Description:** This endpoint allows users to create a new post. The post is stored in the database and associated with the user who created it.
    -   **Request Body:**

        ```json
        {
            "title": "Post Title",
            "content": "Post Content"
        }
        ```

2.  **GET `/posts`**

    -   **Description:** This endpoint returns all posts created by the user.
    -   **Response:**

        ```json
        [
            {
                "id": "postId1",
                "title": "Post Title 1",
                "content": "Content of post 1"
            },
            {
                "id": "postId2",
                "title": "Post Title 2",
                "content": "Content of post 2"
            }
        ]
        ```

3.  **PUT `/posts/:id`**

    -   **Description:** This endpoint allows users to update a post by providing the post ID. The updated post details are stored in the database.
    -   **Request Body:**

        ```json
        {
            "title": "Updated Post Title",
            "content": "Updated Post Content"
        }
        ```

4.  **DELETE `/posts/:id`**

    -   **Description:** This endpoint allows users to delete a post by providing the post ID. The post is removed from the database.
    -   **Response:**

        ```json
        {
            "message": "Post deleted successfully"
        }
        ```

## Docker

The API is now dockerized. First make sure the docker is installed and is up running. Add the .env file else it will give an error. Finally, it can be run using the following commands:

```bash
    docker-compose build
    docker-compose up
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.
