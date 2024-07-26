package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Fallback(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
}