# My Go Web Server Project

This is a simple web server build using Go programming language.

## Description

This project is a beginner-friendly web server implemented in Go. It serves as a starting point for understanding the basics of web development in Go such as routing, handling GET and POST requests, serving static files, and error handling.

## Features

- HTTP Server: Sets up a basic HTTP server to handle incoming requests.
* Routing: Defines routes to handle different URL paths.
+ Request Handling: Shows how to parse form data and handle different HTTP methods.
- Static File Serving: Serves static files from the static directory.

### Function details

#### 'helloHandler'
This function handles GET requests to the "/hello" path. 
It checks if the request method is GET and the URL path is "/hello". 
If the conditions are met, it responds with "Hello!". 
Otherwise, it returns an appropriate error response.

#### 'formHandler'
This function handles POST requests to the "/form" path. 
It parses the form data submitted in the request. 
If parsing is successful, it responds with a success message and prints the submitted name and address. 
If there's an error during form parsing, it returns a bad request error response.

## Usage
Once the server is running, you can access it by navigating to http://localhost:8080 in your web browser.

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.

## License
This project is licensed under the MIT License.

## Acknowledgements
- [Go Documentation](https://golang.org/doc/)
- [Gorilla Mux](https://github.com/gorilla/mux) for routing



Additional Informations:
https://pkg.go.dev/net/http

