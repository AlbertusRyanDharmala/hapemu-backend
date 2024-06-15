package service

import (
	"encoding/json"
	"fmt"
	"hapemu/model"
	"log"
	"net/http"
	"net/smtp"
)

// EmailService represents the email service with configuration
type EmailService struct {
	smtpPort int
	smtpHost string
	username string
	password string
}

// NewEmailService creates a new EmailService
func NewEmailService(port int, host, username, password string) *EmailService {
	return &EmailService{
		smtpPort: port,
		smtpHost: host,
		username: username,
		password: password,
	}
}

// SendEmail sends an email with the given recipient, subject, and body
func (s *EmailService) SendEmail(to, subject, body string) (bool, error) {
	auth := smtp.PlainAuth("", s.username, s.password, s.smtpHost)
	from := s.username

	// Set up headers and body for HTML email
	headers := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n", from, to, subject)
	htmlBody := fmt.Sprintf(`<html><body style="font-family: 'Arial', sans-serif; font-size: 16px;">%s</body></html>`, body)
	msg := []byte(headers + htmlBody)

	// Send the email
	err := smtp.SendMail(fmt.Sprintf("%s:%d", s.smtpHost, s.smtpPort), auth, from, []string{to}, msg)
	if err != nil {
		return false, err
	}
	return true, nil
}

// for testing
// func EmailRecommendation() {
// 	emailService := NewEmailService(587, "smtp.gmail.com", "hapemu.id@gmail.com", "fdvtmvobhemhxvvi")

// 	var messages = `<p>Here are top 5 recommendations for you based on the quiz. We hope you find them useful:</p>
//     <ol>
//         <li>%s</li>
//         <li>smartphone 2</li>
//         <li>smartphone 3</li>
//         <li>smartphone 4</li>
//         <li>smartphone 5</li>
//     </ol>`
// 	formattedMessage := fmt.Sprintf(messages, "Samsung S24")
// 	sent, err := emailService.SendEmail("ard00243@gmail.com", "Email recommendations from hapemu", formattedMessage)

// 	if err != nil {
// 		log.Fatalf("Error when sending email: %s", err)
// 	}

// 	if sent {
// 		log.Println("Email sent successfully")
// 	} else {
// 		log.Println("fail to send email")
// 	}
// }

func EmailRecommendations(w http.ResponseWriter, r *http.Request) {
	var emailRequest model.EmailRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&emailRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	emailService := NewEmailService(587, "smtp.gmail.com", "hapemu.id@gmail.com", "fdvtmvobhemhxvvi")

	var message = `<p>Here are top 5 recommendations for you based on the quiz. We hope you find them useful:</p>
    <ol>
        <li>%s</li>
        <li>%s</li>
        <li>%s</li>
        <li>%s</li>
        <li>%s</li>
    </ol>`

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
