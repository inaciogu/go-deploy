package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Go test deploy</h1>"))
	})
	serverName := os.Getenv("SERVER_NAME")

	fmt.Printf("%s is running on port 8080", serverName)
	http.ListenAndServe(":8080", nil)
}
