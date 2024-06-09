# BookStoreManagement

Project Overview
Project Name: Bookstore Management System

Technology Stack:

Programming Language: Go (Golang)
Web Framework: Gorilla Mux
Database: MySQL (using GORM for ORM)
Tools: Visual Studio Code for IDE, Postman for API testing, and Delve for debugging
Key Components
Main Application (main.go):

Initializes the HTTP server and routes.
Uses Gorilla Mux to handle routing.
Starts the server on a specified port.
Configuration (config/app.go):

Manages the database connection using GORM.
Contains functions to connect to the database and retrieve the DB instance.
Controllers (controllers/book-controller.go):

Handles HTTP requests for the API.
Functions for CRUD operations: CreateBook, GetBook, GetBookById, UpdateBook, DeleteBook.
Utilizes JSON for request and response payloads.
Models (models/book.go):

Defines the Book model and database schema.
Contains methods for interacting with the database (e.g., creating, retrieving, updating, and deleting records).
Routes (routes/bookstore-routes.go):

Registers the API endpoints and associates them with the corresponding controller functions.
Utilities (utils/utils.go):

Helper functions for common tasks, such as parsing the request body.
Explaining Your Project in an Interview
1. Project Overview:

"I developed a RESTful API for a Bookstore Management System using Go (Golang). The application allows users to manage book records in a MySQL database, supporting operations such as creating, reading, updating, and deleting books."
2. Key Technologies:

"I used Gorilla Mux for routing and handling HTTP requests, GORM as the ORM for interacting with a MySQL database, and Delve for debugging the Go code. The API handles JSON payloads for communication between the client and the server."
3. Project Structure:

"The project is organized into several components: main.go for server initialization, config for database configuration, controllers for handling HTTP requests, models for defining database schemas and operations, routes for registering API endpoints, and utils for utility functions."
4. Detailed Functionality:

"In the controllers, each function handles a specific HTTP method. For example, CreateBook handles POST requests to add a new book, GetBook handles GET requests to retrieve all books, GetBookById handles GET requests for a specific book, UpdateBook handles PUT requests to update a book, and DeleteBook handles DELETE requests to remove a book."
"The models package defines the Book struct, which includes fields like Name, Author, and Publication, and uses GORM to map these fields to database columns."
