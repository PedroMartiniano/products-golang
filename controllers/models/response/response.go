package response

import (
	"errors"
	"github.com/PedroMartiniano/products-golang/config"
	"github.com/PedroMartiniano/products-golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func httpError(err error) (code int, message string) {
	var myErr config.Error

	if errors.As(err, &myErr) {
		message = myErr.AppError().Error()

		switch myErr.TypeError() {
		case config.ErrInternalServer:
			code = http.StatusInternalServerError
		case config.ErrNotFound:
			code = http.StatusNotFound
		case config.ErrBadRequest:
			code = http.StatusBadRequest
		}
	}

	return
}

func sendError(c *gin.Context, code int, err string) {
	c.JSON(code, gin.H{"success": false, "message": err})
}

func sendSuccess(c *gin.Context, code int, data any) {
	c.JSON(code, gin.H{"success": true, "data": data})
}

//swagger responses

// error response
type errorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// products
type productSuccessResponse1 struct {
	Success bool           `json:"success"`
	Data    models.Product `json:"data"`
}

type productSuccessResponse2 struct {
	Success bool             `json:"success"`
	Data    []models.Product `json:"data"`
}
