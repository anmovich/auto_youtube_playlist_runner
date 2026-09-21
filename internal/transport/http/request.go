package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
)

type errorResponse struct{
	Message string `json:"message"`
}

func Error(c *gin.Context, message string, status int) {
	log.Println(message)
	c.AbortWithStatusJSON(status, gin.H{"error": message})
}
