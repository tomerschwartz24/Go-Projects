package main

import (
	"fmt"
	"net/http"
	"strings"
)

func contactForm() {
	//Handle POST requests to Contact form
	http.HandleFunc("/user-message", func(w http.ResponseWriter, r *http.Request) {
		// Check if the request came from the same host and if the referer header is set
		if r.Referer() == "" || r.Host != strings.Split(r.Referer(), "/")[2] {
			http.Error(w, "Direct access to resources isn't allowed!", http.StatusForbidden)
			return
		}
		if r.Method == "POST" {
			// Parse the form data
			if returnedData := r.ParseForm(); returnedData != nil {
				http.Error(w, "Failed to parse form", http.StatusBadRequest)
				return
			}
			// Get the name & email & message from the form
			name := r.FormValue("name")
			email := r.FormValue("email")
			message := r.FormValue("message")

			if name == "" || email == "" || message == "" {
				http.Error(w, "Please fill the entire form!", http.StatusBadRequest)
				return
			}
			// Print the message to console
			fmt.Printf("Your name is  %s\n your email is %s\n and your message is %s", name, email, message)

			// Return a response to the user
			fmt.Fprintf(w, "Your details have been sent!")

		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func main() {

	// Handle HTTP requests to "/" and serve the web resources inside the webapp folder ( index.html / images and so on)
	http.Handle("/", http.FileServer(http.Dir("web_resources")))
	contactForm()
	//Error Handling
	webappListener := http.ListenAndServe(":8080", nil)
	if webappListener != nil {
		fmt.Printf("Unable to bind port  %v\n", webappListener)
	} else {
		fmt.Println("Webserver is now listening on port 8080")
	}
}
