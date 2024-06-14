package service

import (
	"encoding/json"
	"fmt"
	"hapemu/model"
	"log"
	"net/http"

	"github.com/MakMoinee/go-mith/pkg/email"
)

func EmailRecommendation() {
	emailService := email.NewEmailService(587, "smtp.gmail.com", "onlyforancile@gmail.com", "zhqwpzuzjhivnvsx")

	var messages = `Here are top 5 recommendations for you based on the quiz. We hope you find them useful:
	1. smartphone 1
	2. smartphone 2
	3. smartphone 3
	4. smartphone 4
	5. smartphone 5`
	sent, err := emailService.SendEmail("yosuajayapura@gmail.com", "Email recommendations from hapemu", messages)

	if err != nil {
		log.Fatalf("Error when sending email: %s", err)
	}

	if sent {
		log.Println("Email sent successfully")
	} else {
		log.Println("fail to send email")
	}
}

func EmailRecommendations(w http.ResponseWriter, r *http.Request) {
	var emailRequest model.EmailRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&emailRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	emailService := email.NewEmailService(587, "smtp.gmail.com", "onlyforancile@gmail.com", "zhqwpzuzjhivnvsx")

	var message = `Here are top 5 recommendations for you based on the quiz. We hope you find them useful:
	1. %s
	2. %s
	3. %s
	4. %s
	5. %s`

	var recommendations = emailRequest.Recommendations
	formattedMessage := fmt.Sprintf(message, recommendations[0], recommendations[1], recommendations[2], recommendations[3], recommendations[4])
	sent, err := emailService.SendEmail(emailRequest.UserEmail, "Email recommendations from hapemu", formattedMessage)

	if err != nil {
		log.Fatalf("Error when sending email: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")

	// Set the response body to a JSON string

	var emailResponse model.EmailResponse
	if sent {
		emailResponse.Success = true
		emailResponse.Message = "Email sent successfully"
	} else {
		emailResponse.Success = false
		emailResponse.Message = "Fail to send email!!"
	}
	response, err := json.Marshal(emailResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
}
