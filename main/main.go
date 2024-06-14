package main

import (
	"hapemu/service"
	"log"
	"net/http"
)

func main() {
	// emailRequest := model.EmailRequest{
	// 	UserEmail:       "ard00243@gmail.com",
	// 	Recommendations: []string{"samsung S24", "iphone 5"},
	// }
	// requestBody, err := json.Marshal(emailRequest)
	// if err != nil {
	// 	fmt.Println("failed to Marshal")
	// }

	// // Create a new HTTP request
	// req, err := http.NewRequest("POST", "/email-recommendation", bytes.NewBuffer(requestBody))
	// if err != nil {
	// 	fmt.Println("failed to create request")
	// }
	// // Create a new HTTP recorder to capture the response
	// recorder := httptest.NewRecorder()

	// // Call the function to handle the HTTP request
	// service.EmailRecommendation(recorder, req)
	// service.TestEmail2()
	service.EmailRecommendation()

	// http.HandleFunc("/antutu", service.GetAntutuList)
	// http.HandleFunc("/dxomark", service.GetDxoMarkList)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
