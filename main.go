package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Life Heatmap server is running buddy!")
}

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
