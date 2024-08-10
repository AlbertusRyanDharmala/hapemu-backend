package service

import (
	"encoding/json"
	"fmt"
	"hapemu/model"
	"net/http"
)

func CompareSmartphone(w http.ResponseWriter, r *http.Request) {
	var comparisonRequest model.ComparisonRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&comparisonRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(comparisonRequest.PhoneName1)
	fmt.Println(comparisonRequest.PhoneName2)

	var comparisonResponse model.ComparisonResponse
	if comparisonRequest.PhoneName1 == "" || comparisonRequest.PhoneName1 == "null" || comparisonRequest.PhoneName2 == "" || comparisonRequest.PhoneName2 == "null" {
		comparisonResponse.Similarity = 0
		response, err := json.Marshal(comparisonResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(response)
		return
	}

	var phone1 = getSmartphoneByName(comparisonRequest.PhoneName1)
	var phone2 = getSmartphoneByName(comparisonRequest.PhoneName2)
	var vectors = convertSmartphoneToVecForCompare(phone1, phone2)
	comparisonResponse.Similarity = cosineSimilarity(vectors[0], vectors[1])
	fmt.Printf("similarity %f\n", comparisonResponse.Similarity)

	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(comparisonResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
}
