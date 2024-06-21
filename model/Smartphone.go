package model

type RecommendationsRequest struct {
	Price     string `json:"price"`
	Processor string `json:"processor"`
	Camera    string `json:"camera"`
	Baterry   string `json:"battery"`
	Ram       string `json:"ram"`
	Storage   string `json:"storage"`
}

// Struct to hold an item and its similarity score
type SmartphoneSimilarity struct {
	Name       string
	Similarity float64
}

type RecommendationsResponse struct {
	Recommendations []string `json:"recommendations"`
}

type Smartphone struct {
	Name         string
	SegmentPrice string
	Processor    string
	DxomarkScore int
	Battery      string
	Ram          string
	Storage      string
}
