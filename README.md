# Book Store API Project

This is a simple REST API project for managing books and employees using Golang. The project utilizes MongoDB as the database, and it is structured into various folders to maintain modularity and scalability.

### **Folder Structure**

```
book_store/
├── main.go
├── config/
│   └── database.go
├── models/
│   ├── book.go
│   └── employee.go
├── controllers/
│   ├── book_controller.go
│   └── employee_controller.go
├── routes/
│   └── routes.go
├── dist/
└── .env
```

### **Project Overview**

- **Golang**: The main programming language used for the backend REST API.
- **MongoDB**: Used as the NoSQL database to store book and employee data.
- **Folder Structure**: The project is organized into separate folders for configuration, models, controllers, routes, and the main entry point, `main.go`.

### **Current Status**

- The project was developed with the assistance of AI to speed up the implementation of various features.
- MongoDB connection is currently not working, and further troubleshooting is needed to fix the database connectivity issue.
- The project is intended to be deployed in the future, so the deployment plan will be added once all functionality is properly tested and verified.

### **How to Run the Project**

1. Clone the repository:
   ```sh
   git clone <repository_url>
   cd book_store
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Create a `.env` file in the root of the project and add the following:
   ```env
   MONGODB_URI=mongodb+srv://<username>:<password>@cluster0.mongodb.net/?retryWrites=true&w=majority
   MONGODB_DATABASE=bookstore
   ```
   Replace `<username>` and `<password>` with your MongoDB credentials.

4. Run the server:
   ```sh
   go run main.go
   ```

5. The server will start on port `8080`. You can access the API at `http://localhost:8080`.

### **Endpoints**

#### **Books**
- **GET /books**: Retrieve all books.
- **GET /books/{id}**: Retrieve a specific book by ID.
- **POST /books**: Create a new book.
- **PUT /books/{id}**: Update an existing book.
- **DELETE /books/{id}**: Delete a book.

#### **Employees**
- **GET /employees**: Retrieve all employees.
- **POST /employees**: Create a new employee.

### **Deployment Plan**

- The project is intended to be deployed using a cloud provider (e.g., AWS, GCP, or Azure) once all core functionalities are properly implemented and tested.

### **Acknowledgment**

- This project was developed with assistance from AI to streamline the development process, especially in setting up the basic structure and implementation.
