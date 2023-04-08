package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("web_resources")))
	fmt.Println("Server started on port 8085")
	http.ListenAndServe(":8085", nil)
}
