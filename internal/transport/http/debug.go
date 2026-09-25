package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Clear_lock(c *gin.Context){
	type Username struct{
		Username string `json:"username"`
	}
	var user Username
	if err := json.NewDecoder(c.Request.Body).Decode(&user); err != nil{
		Error(c, "Failed to ClearLocked", http.StatusBadRequest)
	}
	h.Service.Authorization.ClearLocked(user.Username)
}
