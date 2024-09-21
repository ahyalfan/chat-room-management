# Golang Base Project

This is a base project for implementasi simple chat app using Golang with the hibiken/asynq for server redis queue. godotenv Configuration is managed via an env file.

## Features

- **Redis queue server**
- **Send Email**: Send email service
- **Env File**: Simple configuration management using environment variables.

## Getting Started

### Prerequisites

Ensure you have the following installed:

- [Golang](https://golang.org/dl/)
- [Git](https://git-scm.com/)

### Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/shellrean/golang-base-project-clean-directory.git
   cd golang-base-project-clean-directory
   ```

2. **Install dependencies:**

   ```sh
   go mod tidy
   ```

3. **Create and configure `.env` file:**

   Create a `.env` file in the root directory and add your configuration variables.

   ```env

   SERVER_HOST=localhost
   SERVER_PORT=8700


    MAIL_HOST=smtp.gmail.com
    MAIL_PORT=587
    MAIL_USERNAME=example@example.com
    MAIL_PASSWORD=

    REDIS_ADDR=localhost:6379
    REDIS_PASSWORD=
    REDIS_DB=0
   ```

````

### Running the Application

Start the application with the following command:

```sh
go run main.go
```
````
