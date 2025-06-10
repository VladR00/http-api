package main

import (
	"fmt"
	"log"
	"net/http"

	handlers "httpapi/internal/handlers"
)

func main() {
	http.HandleFunc("/task", handlers.HandlerTask)        // (POST & GET)(add/get) / OPTION(may be)
	http.HandleFunc("/task/", handlers.HandlerDELETETask) // DELETE

	fmt.Println("Server start at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	select {}
}
