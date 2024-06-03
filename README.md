# Simple Calculator

## Description

This project is a simple web application that generates Simple Calculator from 2 input. Users can enter first number and sacent and put Calculated for display result.

## Authors

<a href="https://learn.zone01oujda.ma/git/ikazbat">
  <img src="https://learn.zone01oujda.ma/git/avatars/a2d348a39c50555da8ad3cc337e6a349?size=870" title="Idriss Kazbat" width="40"height="40"/>
</a>

## Usage

To run the application:

1. Ensure you have [Go](https://golang.org/dl/) installed.
2. Clone the repository.
3. Navigate to the project directory.
4. Run the server using the command: `go run main.go`.
5. Open your browser and go to `http://localhost:8080`.

## Implementation details

### Endpoints

- `GET /`: Renders the main page with the form.
- `POST /`: Accepts form input and generates Calculator.

### HTTP Status Codes

- `200 OK`: Request was successful.
- `400 Bad Request`: Missing or incorrect parameters in the request.
- `404 Not Found`: Resource not found.
- `500 Internal Server Error`: An error occurred on the server.

