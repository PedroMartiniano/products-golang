package controllers

import "github.com/gin-gonic/gin"

func sendError(c *gin.Context, code int, err string) {
	c.JSON(code, gin.H{"success": false, "message": err})
}

func sendSuccess(c *gin.Context, code int, data any) {
	c.JSON(code, gin.H{"success": true, "data": data})
}
