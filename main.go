package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

// Get the port environment variable
func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Println("No PORT variable was set. Using default port: " + port)
	}
	return ":" + port
}
