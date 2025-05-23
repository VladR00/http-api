package main

import (
	"fmt"
	"log"
	"net/http"

	handlers "webrestapi/internal/handlers"
)

func main() {
	http.HandleFunc("/quotes", handlers.HandlerQuote)                 // POST & GET & OPTION (add/get/delete)
	http.HandleFunc("/quotes/random", handlers.HandlerGETRandomQuote) // GET (getrand)
	http.HandleFunc("/quotes/", handlers.HandlerDELETEQuote)          // DELETE

	fmt.Println("Server start at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	select {}
}
