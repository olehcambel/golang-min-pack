package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Web Server with GOlang")
	log.Println("/")
}

func main() {
	log.Println("Server has started!")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":6060", nil)
}
