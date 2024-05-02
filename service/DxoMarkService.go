package service

import (
	"encoding/json"
	"hapemu/model"
	"net/http"
)

func GetDxoMarkList(w http.ResponseWriter, r *http.Request) {
	// Access database and retrieve DXOMark data

	// Assuming Num1 and Num2 are fields of the model.Handphone struct
	decoder := json.NewDecoder(r.Body)
	var request model.Handphone
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Perform some calculation based on request data
	dxomark := request.Num1 - request.Num2

	// Create the response
	response := model.Handphone{
		// Set appropriate fields based on the calculation or database query
		// For example:
		Name:    "Example Handphone",
		DXOMark: dxomark,
		// Add more fields as needed
	}

	// Write the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
