package main

import (
	"fmt"
	"log"
	"net/http"
)

func anotherPage(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome to Another Page"))
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.HandleFunc("/page", anotherPage)
	http.Handle("/", fileServer)

	fmt.Printf("Starting serving at port 8080")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}