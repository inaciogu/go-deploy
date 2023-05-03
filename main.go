package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("/etc/secrets/.env")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Go test deploy</h1>"))
	})
	serverName := os.Getenv("SERVER_NAME")

	fmt.Printf("%s is running on port 8080", serverName)
	http.ListenAndServe(":8080", nil)
}
