package models

// TextPair holds the incoming text pairs for comparison
type TextPair struct {
	FirstText  string `json:"first_text"`
	SecondText string `json:"second_text"`
}

// SimilarityScore holds the normalized similarity score
type SimilarityScore struct {
	Score float32 `json:"similarity_score"`
}
