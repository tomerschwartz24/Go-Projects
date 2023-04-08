package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world"), func (w http.ResponseWriter, r http.Request)  {
		w.Write([])byteconst	
	
		
	}
	http.ListenAndServe(":8080", nil)
}
