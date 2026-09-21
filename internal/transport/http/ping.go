package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "All good \n")
}
