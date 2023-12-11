# Book Management System

This is a simple Book Management System implemented in Go (Golang). It provides RESTful API endpoints to manage books such as creating, reading, updating, and deleting book entries. The system utilizes Gorilla Mux for routing and GORM as the ORM (Object-Relational Mapping) for working with the MySQL database.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Interacting with API](#interacting-with-api)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/book-management-system.git
    ```

2. Navigate to the project directory:

    ```bash
    cd book-management-system
    ```

3. Set up a MySQL database and configure the connection details in the `config/Connect()` function.

4. Install dependencies:

    ```bash
    go mod tidy
    ```

## Usage

To run the application, execute the following command:

```bash
go run main.go
```

This will start the server at 127.0.0.1:3306.

## Endpoints
The following API endpoints are available:

- **POST /book/**: Create a new book entry.
- **GET /book/**: Retrieve all books.
- **GET /book/{bookID}**: Retrieve details of a specific book.
- **PUT /book/{bookID}**: Update details of a specific book.
- **DELETE /book/{bookID}**: Delete a specific book.

## Interacting with API
You can use tools like Postman or cURL to interact with the API endpoints. Here's an example of how to use Postman:

1. Creating a Book:
- Set the HTTP method to POST.
- Use the endpoint http://127.0.0.1:3306/book/.
- Include the book details in the request body as JSON.

2. Retrieving All Books:
- Set the HTTP method to GET.
- Use the endpoint http://127.0.0.1:3306/book/.

3. Retrieving a Specific Book:
- Set the HTTP method to GET.
- Use the endpoint http://127.0.0.1:3306/book/{bookID} (replace {bookID} with the actual book ID).

4. Updating a Book:
- Set the HTTP method to PUT.
- Use the endpoint http://127.0.0.1:3306/book/{bookID} (replace {bookID} with the actual book ID).
- Include the updated book details in the request body as JSON.

5. Deleting a Book:
- Set the HTTP method to DELETE.
- Use the endpoint http://127.0.0.1:3306/book/{bookID} (replace {bookID} with the actual book ID).

## Contributing
Contributions are welcome! If you have any suggestions, bug reports, or improvements, feel free to open an issue or create a pull request.

## License
This project is licensed under the MIT License.


You can use this Markdown content directly by copying and pasting it into your README.md file in your GitHub repository. Feel free to further customize or enhance it according to your project's specifications.
