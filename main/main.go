package main

import (
	"hapemu/service"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/sendEmail", service.EmailRecommendations)

	// http.HandleFunc("/antutu", service.GetAntutuList)
	// http.HandleFunc("/dxomark", service.GetDxoMarkList)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
