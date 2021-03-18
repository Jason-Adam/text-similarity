package services

import "github.com/jason-adam/text-similarity/internal/api/models"

// SimilarityService calculates the similarity between two text entries
type SimilarityService struct{}

// NewSimilarityService returns a new SimilarityService
func NewSimilarityService() *SimilarityService {
	return &SimilarityService{}
}

// Calculate returns the similarity score
func (s *SimilarityService) Calculate(t *models.TextPair) (*models.SimilarityScore, error) {
	// TODO: implement levenschtein distance
	return &models.SimilarityScore{
		Score: 0.5,
	}, nil
}
