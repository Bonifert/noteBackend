# About the Project
This project is a simple Go backend application designed for educational purposes. 
It showcases the use of the net/http package for handling HTTP requests and responses, 
alongside PostgreSQL for data storage and GORM for database interactions. The application includes 
essential features such as user registration, authentication using JWT tokens, and the capability 
for authenticated users to create notes. This project was undertaken to explore and learn Go, 
focusing on simplicity and practicality, especially since the primary backend language of 
the developer is Java.

## Project Status
This project is currently in the Work-In-Progress (WIP) phase. Initial development is underway, focusing on creating a backend system utilizing Go, PostgreSQL, and GORM. The goal is to implement user registration, authentication using JWT tokens, and note creation functionalities. While the core features are being developed, please note that the application is not yet complete and may undergo changes as development progresses.

## Technologies Use:
* Backend: Go
* HTTP Package: net/http for managing HTTP requests and responses
* Database: PostgreSQL
* ORM: GORM for database interactions
* Authentication: JWT (JSON Web Tokens) for secure user authentication
## Key Features: 
* User Registration: Enables new users to sign up and create an account.
* Authentication: Implements JWT-based authentication to ensure secure access to protected routes.
* Note Creation: Authenticated users can create notes, demonstrating basic CRUD operations.

## Getting Started
### Prerequisites:
1. Go installed on your machine.
2. PostgreSQL database set up and running.

### Installation
1. Clone the repository to your local machine.
2. Navigate to the project directory.
3. Ensure you have a PostgreSQL database running and accessible.
4. Fill in the missing fields in the .envtemplate with the data required for the database connection, and rename it to ".env".

### Usage:

Run the application using go run main.go.
