package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"hapemu/model"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "aws-0-ap-southeast-1.pooler.supabase.com"
	port     = 5432
	user     = "postgres.yovcppevikilglvpktzq"
	password = "HapemuPostgres123"
	dbname   = "postgres"
)

// region convert smartphone from database to vector
func getValueForPrice(price string) float64 {
	if strings.Compare(price, "essential") == 0 {
		return 1
	} else if strings.Compare(price, "mid") == 0 {
		return 2
	} else if strings.Compare(price, "high") == 0 {
		return 3
	}
	return 4
}

func getValueForProcessor(processor string) float64 {
	sTierProcessors := []string{
		"Snapdragon 8 Gen 3",
		"Snapdragon 8s Gen 3",
		"Snapdragon 8 Gen 2",
		"Dimensity 9300",
		"Dimensity 9200",
		"Dimensity 9000",
		"Exynos 2200",
		"Exynos 2100",
		"Kirin 9000",
		"Kirin 9000e",
		"Apple A17 Bionic",
		"Apple A16 Bionic",
		"Apple A15 Bionic",
	}

	aTierProcessors := []string{
		"Snapdragon 8+ Gen 1",
		"Snapdragon 8 Gen 1",
		"Snapdragon 7+ Gen 2",
		"Snapdragon 888+",
		"Snapdragon 888",
		"Snapdragon 870",
		"Snapdragon 865+",
		"Dimensity 8300",
		"Dimensity 8200",
		"Dimensity 8100",
		"Exynos 1080",
		"Exynos 990",
		"Exynos 9825",
		"Exynos 9820",
		"Kirin 990",
		"Kirin 985",
		"Kirin 980",
		"Apple A14 Bionic",
		"Apple A13 Bionic",
	}

	bTierProcessors := []string{
		"Snapdragon 7 Gen 3",
		"Snapdragon 7s Gen 2",
		"Snapdragon 7 Gen 1",
		"Snapdragon 782",
		"Snapdragon 780",
		"Snapdragon 778G",
		"Snapdragon 768",
		"Snapdragon 765G",
		"Snapdragon 750",
		"Snapdragon 732",
		"Dimensity 8050",
		"Dimensity 8020",
		"Dimensity 8000",
		"Dimensity 7200",
		"Dimensity 7000",
		"Exynos 980",
		"Exynos 9810",
		"Exynos 9611",
		"Exynos 9610",
		"Helio G99",
		"Helio G96",
		"Helio G95",
		"Helio G90T",
		"Kirin 970",
		"Kirin 810",
		"Kirin 820 5G",
		"Apple A12 Bionic",
		"Apple A11 Bionic",
	}

	cTierProcessors := []string{
		"Snapdragon 6 Gen 1",
		"Snapdragon 4 Gen 2",
		"Snapdragon 4 Gen 1",
		"Snapdragon 730G",
		"Snapdragon 730",
		"Snapdragon 720",
		"Snapdragon 712",
		"Snapdragon 710",
		"Snapdragon 695",
		"Snapdragon 690",
		"Snapdragon 685",
		"Snapdragon 680",
		"Snapdragon 675",
		"Snapdragon 670",
		"Snapdragon 665",
		"Snapdragon 662",
		"Snapdragon 660",
		"Dimensity 7020",
		"Dimensity 7030",
		"Dimensity 6100",
		"Dimensity 6080",
		"Exynos 850",
		"Exynos 8895",
		"Helio G88",
		"Helio G85",
		"Helio G80",
		"Helio G70",
		"Helio P90",
		"Helio P70",
		"Helio P65",
		"Helio P60",
		"Kirin 710",
		"Kirin 659",
		"Apple A10 Fusion",
	}

	for _, cur := range sTierProcessors {
		if strings.Contains(processor, cur) {
			return 4
		}
	}

	for _, cur := range aTierProcessors {
		if strings.Contains(processor, cur) {
			return 3
		}
	}
	for _, cur := range bTierProcessors {
		if strings.Contains(processor, cur) {
			return 2
		}
	}
	for _, cur := range cTierProcessors {
		if strings.Contains(processor, cur) {
			return 1
		}
	}
	return 4
}

func getValueForCamera(camera int) float64 {
	if camera <= 75 {
		return 1
	} else if camera <= 104 {
		return 2
	} else if camera <= 134 {
		return 3
	}
	return 4
}

func getValueForBattery(battery string) float64 {
	var batteryValue, err = strconv.Atoi(battery)
	if err != nil {
		fmt.Println("error converting battery string to integer")
	}
	if batteryValue < 4000 {
		return 1
	} else if batteryValue < 4500 {
		return 2
	} else if batteryValue < 5000 {
		return 3
	}
	return 4
}

func getVecValueFromRam(ram string) float64 {
	if ram == "1" || ram == "2" || ram == "4" {
		return 1
	} else if ram == "6" || ram == "8" {
		return 2
	} else if ram == "12" {
		return 3
	} else {
		return 4
	}
}

func getValueForRam(ram string, ramVec float64) float64 {
	var ramList = []string{"1", "2", "4", "6", "8", "12", "16", "32"}
	var minVec, maxVec float64
	for _, cur := range ramList {
		if strings.Contains(ram, cur) {
			minVec = getVecValueFromRam(cur)
			break
		}
	}
	for _, cur := range ramList {
		if strings.Contains(ram, cur) {
			maxVec = getVecValueFromRam(cur)
		}
	}
	if ramVec < minVec {
		return minVec
	} else if ramVec >= minVec && ramVec <= maxVec {
		return ramVec
	} else {
		return maxVec
	}
}

func getVecValueFromStorage(storage string) float64 {
	if storage == "32GB" || storage == "64GB" || storage == "128GB" {
		return 1
	} else if storage == "256GB" {
		return 2
	} else if storage == "512GB" {
		return 3
	} else {
		return 4
	}
}

func getValueForStorage(storage string, storageVec float64) float64 {
	var storageList = []string{"32GB", "64GB", "128GB", "256GB", "512GB", "1TB"}
	var minVec, maxVec float64
	for _, cur := range storageList {
		if strings.Contains(storage, cur) {
			minVec = getVecValueFromStorage(cur)
			break
		}
	}
	for _, cur := range storageList {
		if strings.Contains(storage, cur) {
			maxVec = getVecValueFromStorage(cur)
		}
	}
	if storageVec < minVec {
		return minVec
	} else if storageVec >= minVec && storageVec <= maxVec {
		return storageVec
	} else {
		return maxVec
	}
}

func convertSmartphoneToVec(smartphone model.Smartphone, targetVec []float64) []float64 {
	var smartphonesVecs []float64
	smartphonesVecs = append(smartphonesVecs, getValueForPrice(smartphone.SegmentPrice))            // price
	smartphonesVecs = append(smartphonesVecs, getValueForProcessor(smartphone.Processor))           // processor
	smartphonesVecs = append(smartphonesVecs, getValueForCamera(smartphone.DxomarkScore))           // camera
	smartphonesVecs = append(smartphonesVecs, getValueForBattery(smartphone.Battery))               // battery
	smartphonesVecs = append(smartphonesVecs, getValueForRam(smartphone.Ram, targetVec[4]))         // ram
	smartphonesVecs = append(smartphonesVecs, getValueForStorage(smartphone.Storage, targetVec[5])) // storage
	return smartphonesVecs
}

// endregion

// region apply cosine similarity algorithm
// Function to calculate dot product of two vectors
func dotProduct(vec1, vec2 []float64) float64 {
	var dotProduct float64
	for i := 0; i < len(vec1); i++ {
		dotProduct += vec1[i] * vec2[i]
	}
	return dotProduct
}

// Function to calculate magnitude of a vector
func magnitude(vec []float64) float64 {
	var sumSquares float64
	for _, val := range vec {
		sumSquares += val * val
	}
	return math.Sqrt(sumSquares)
}

// Function to calculate cosine similarity between two vectors
func cosineSimilarity(vec1, vec2 []float64) float64 {
	return dotProduct(vec1, vec2) / (magnitude(vec1) * magnitude(vec2))
}

// Function to recommend movies based on cosine similarity
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

// endregion

// region convert user quiz to vector
func getPriceValue(price string) float64 {
	if strings.Contains(price, "essensial") {
		return 1
	} else if strings.Contains(price, "midrange") {
		return 2
	} else if strings.Contains(price, "premium") {
		return 3
	}
	return 4
}

func getValue(str string) float64 {
	if strings.Contains(str, "tidak") {
		return 1
	} else if strings.Contains(str, "cukup") {
		return 2
	} else if strings.Contains(str, "penting") {
		return 3
	}
	return 4
}

func convertRecommendationRequestToTargetVec(request model.RecommendationsRequest) []float64 {
	var vec []float64

	vec = append(vec, getPriceValue(request.Price))
	vec = append(vec, getValue(request.Processor))
	vec = append(vec, getValue(request.Camera))
	vec = append(vec, getValue(request.Baterry))
	vec = append(vec, getValue(request.Ram))
	vec = append(vec, getValue(request.Storage))

	return vec
}

//endregion

// main function
func RecommendSmartphones(w http.ResponseWriter, r *http.Request) {
	var recommendationsRequest model.RecommendationsRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&recommendationsRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var targetPhoneVec = convertRecommendationRequestToTargetVec(recommendationsRequest)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var smartphones []model.Smartphone
	sqlStatement := `SELECT name, "segmentPrice", processor, "dxomarkScore", battery, ram, storage FROM "Smartphone"`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var smartphone model.Smartphone
		err := rows.Scan(&smartphone.Name, &smartphone.SegmentPrice, &smartphone.Processor, &smartphone.DxomarkScore, &smartphone.Battery, &smartphone.Ram, &smartphone.Storage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		smartphones = append(smartphones, smartphone)
	}

	var similarities = recommendSmartphone(smartphones, targetPhoneVec)
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

// func TestGetSmartphone() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		fmt.Println("failed on connecting to database " + err.Error())
// 		return
// 	}
// 	defer db.Close()

// 	var smartphones []model.Smartphone
// 	sqlStatement := `SELECT name, "segmentPrice", processor, "dxomarkScore", battery, ram, storage FROM "TestSmartphone"`
// 	rows, err := db.Query(sqlStatement)
// 	if err != nil {
// 		fmt.Println("failed on query " + err.Error())
// 		return
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var smartphone model.Smartphone
// 		err := rows.Scan(&smartphone.Name, &smartphone.SegmentPrice, &smartphone.Processor, &smartphone.DxomarkScore, &smartphone.Battery, &smartphone.Ram, &smartphone.Storage)
// 		if err != nil {
// 			fmt.Println("Error when scanning to go struct " + err.Error())
// 			return
// 		}
// 		smartphones = append(smartphones, smartphone)
// 	}

// 	for _, phone := range smartphones {
// 		fmt.Println(phone.Name)
// 		fmt.Println(phone.SegmentPrice)
// 		fmt.Println(phone.Processor)
// 		fmt.Println(phone.DxomarkScore)
// 		fmt.Println(phone.Battery)
// 		fmt.Println(phone.Ram)
// 		fmt.Println(phone.Storage)
// 	}
// }
