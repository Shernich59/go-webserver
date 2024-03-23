package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
// helloHandler handles GET requests to "/hello" path.
// It responds with "Hello!" if the request method is GET
// and the URL path is "/hello".
*/

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "Invalid path to helloHandler", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
	}

	fmt.Fprintf(w, "hello!")
}

/*
// formHandler handles POST requests to "/form" path.
// It parses the form data and responds with a success message
// along with the name and address submitted in the form.
*/
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("ParseForm() err: %v", err), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {
	// Static file serving
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Route handling
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	// Start the HTTP server
	log.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
