package handlers

import (
	"fmt"
	"net/http"

	cors "webrestapi/internal/cors"
)

func HandlerGETRandomQuote(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(w)
	fmt.Println("gigi")
}

func HandlerQuote(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(w)

	switch r.Method {
	case "GET":
		HGETQuote(w, r)
	case "POST":
		HPOSTQuote(w, r)
	case "DELETE":
		HDELETEQuote(w, r)
	}

}

func HPOSTQuote(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(w)

}

func HGETQuote(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(w)

}

func HDELETEQuote(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(w)

}
