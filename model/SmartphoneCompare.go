package model

// Struct to hold an item and its similarity score
type ComparisonRequest struct {
	PhoneName1 string `json:"phoneName1"`
	PhoneName2 string `json:"phoneName2"`
}

type ComparisonResponse struct {
	Similarity float64 `json:"similarity"` // in percentage
}
