package services

import (
	"github.com/jason-adam/text-similarity/internal/api/models"
)

// SimilarityService calculates the similarity between two text entries
type SimilarityService struct{}

// NewSimilarityService returns a new SimilarityService
func NewSimilarityService() *SimilarityService {
	return &SimilarityService{}
}

// Calculate returns the normalized (0-1) similarity score
func (s *SimilarityService) Calculate(t *models.TextPair) (*models.SimilarityScore, error) {

	edits := make([][]int, len(t.SecondText)+1)
	for y := range edits {
		edits[y] = make([]int, len(t.FirstText)+1)
		for x := range edits[y] {
			edits[y][x] = x
		}
	}

	for i := 1; i < len(t.SecondText)+1; i++ {
		edits[i][0] = edits[i-1][0] + 1
	}

	for i := 1; i < len(t.SecondText)+1; i++ {
		for j := 1; j < len(t.FirstText)+1; j++ {
			if t.SecondText[i-1] == t.FirstText[j-1] {
				edits[i][j] = edits[i-1][j-1]
			} else {
				edits[i][j] = 1 + min(edits[i-1][j-1], edits[i-1][j], edits[i][j-1])
			}
		}
	}

	score := edits[len(t.SecondText)][len(t.FirstText)]
	normalizedScore := normalize(score, t)

	return &models.SimilarityScore{
		Score: normalizedScore,
	}, nil
}

func min(args ...int) int {
	curr := args[0]
	for _, num := range args {
		if curr > num {
			curr = num
		}
	}
	return curr
}

func max(args ...int) int {
	curr := args[0]
	for _, num := range args {
		if curr < num {
			curr = num
		}
	}
	return curr
}

func normalize(s int, t *models.TextPair) float32 {
	maxDistance := max(len(t.FirstText), len(t.SecondText))
	normalizedScore := 1.0 - (float32(s) / float32(maxDistance))
	return normalizedScore
}
