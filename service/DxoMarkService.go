package service

import (
	"encoding/json"
	"fmt"
	"hapemu/model"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func InsertDxoMarkToDatabase() []model.DxoMark {
	// Command to execute Python script with arguments
	fmt.Println("test1")
	cmd := exec.Command("python", "C:\\Kuliah\\Skripsi\\hapemu-scrape\\dxomark.py")
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error running Python script: %v", err)
	}

	// Path to the output file
	fmt.Println("test2")
	outputFilePath := "C:\\Kuliah\\Skripsi\\hapemu\\dxomark.json"
	fmt.Println("Output JSON file saved at:", outputFilePath)

	// Read the JSON file
	fmt.Println("test3")
	jsonFile, err := os.Open(outputFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer jsonFile.Close()

	// Decode the JSON file into a struct
	fmt.Println("test4")
	var dxomarkList []model.DxoMark
	err = json.NewDecoder(jsonFile).Decode(&dxomarkList)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}

	// Access the data as needed
	fmt.Println("test5")
	fmt.Println(len(dxomarkList))
	fmt.Println("dxomarkList:", dxomarkList)
	return dxomarkList
}

func GetDxoMarkList(w http.ResponseWriter, r *http.Request) {
	// Access database and retrieve DXOMark data

	// decoder := json.NewDecoder(r.Body)
	// var request model.DxoMarkRequest
	// err := decoder.Decode(&request)
	// if err != nil {
	// 	http.Error(w, "Invalid request body", http.StatusBadRequest)
	// 	return
	// }

	// // Write the response as JSON
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(response)
}
