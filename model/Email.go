package model

type EmailRequest struct {
	UserEmail       string   `json:"email"`
	Recommendations []string `json:"recommendations"`
}

type EmailResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
