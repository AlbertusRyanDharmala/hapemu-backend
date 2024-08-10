package service

import (
	"encoding/json"
	"hapemu/model"
	"net/http"
	"sort"

	_ "github.com/lib/pq"
)

// main function
func recommendSmartphone(smartphones []model.Smartphone, targetPhoneVec []float64) []model.SmartphoneSimilarity {
	var similarities []model.SmartphoneSimilarity

	for _, smartphone := range smartphones {
		similarity := cosineSimilarity(convertSmartphoneToVec(smartphone, targetPhoneVec), targetPhoneVec)
		similarities = append(similarities, model.SmartphoneSimilarity{Name: smartphone.Name, Similarity: similarity})
	}
	sort.Slice(similarities, func(i, j int) bool {
		return similarities[i].Similarity > similarities[j].Similarity
	})

	if len(similarities) > 5 {
		return similarities[:5]
	}
	return similarities
}

func RecommendSmartphones(w http.ResponseWriter, r *http.Request) {
	var recommendationsRequest model.RecommendationsRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&recommendationsRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var targetPhoneVec = convertRecommendationRequestToTargetVec(recommendationsRequest)
	var similarities = recommendSmartphone(getSmartphoneList(), targetPhoneVec)
	var recommendationsResponse model.RecommendationsResponse
	for _, similarity := range similarities {
		recommendationsResponse.Recommendations = append(recommendationsResponse.Recommendations, similarity.Name)
	}

	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(recommendationsResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
}
