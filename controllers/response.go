package controllers

import (
	"github.com/PedroMartiniano/products-golang/models"
	"github.com/gin-gonic/gin"
)

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
