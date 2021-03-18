package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	httperr "github.com/jason-adam/text-similarity/internal/api/http_err"
	"github.com/jason-adam/text-similarity/internal/api/models"
	"github.com/jason-adam/text-similarity/internal/api/services"
)

// CreateSimilarity calculates the similartiy metric between two input strings
func CreateSimilarity(c *gin.Context) {
	s := services.NewSimilarityService()
	var textPair models.TextPair
	_ = c.BindJSON(&textPair)

	if score, err := s.Calculate(&textPair); err != nil {
		httperr.NewHTTPError(c, http.StatusNotFound, errors.New("unable to calculate score"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, score)
	}
}
