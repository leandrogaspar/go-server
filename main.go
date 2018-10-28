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
	http.ListenAndServe(GetPort(), nil)
}

// Get the port environment variable
func GetPort() string {
	log.Println("Loading env port...")
	var port = os.Getenv("PORT")
	log.Println("Env port: " + port)
	if port == "" {
		port = "3000"
		log.Println("No PORT variable was set. Using default port: " + port)
	}
	return ":" + port
}
