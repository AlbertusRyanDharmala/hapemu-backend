package main

import (
	"hapemu/service"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	http.HandleFunc("/send-email", service.EmailRecommendations)

	handler := cors.Default().Handler(http.DefaultServeMux)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
