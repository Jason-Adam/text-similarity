package httperr

import "github.com/gin-gonic/gin"

// HTTPError holds a generic error response
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// NewHTTPError creates and JSON-ifies an HTTPError struct
func NewHTTPError(c *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	c.JSON(status, er)
}
