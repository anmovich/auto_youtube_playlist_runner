package handlers

import (
	"encoding/json"
	"net/http"
	"ypp/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler)SingUP(c *gin.Context) {
	var input models.User
	if err := json.NewDecoder(c.Request.Body).Decode(&input); err != nil{
		Error(c, err.Error(), http.StatusBadRequest)
	}
	id, err := h.Service.Authorization.CreateUser(&input)
	if err != nil{
		Error(c, err.Error(), http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, map[string]any{
		"id":  id, 
	},)
}

func (h *Handler)SingIN(c *gin.Context) {
	var input models.UserSingInUsername
	if err := json.NewDecoder(c.Request.Body).Decode(&input); err != nil{
		Error(c, "Smth wrong", http.StatusBadRequest)
		return
	}
	err := h.Service.Authorization.SignIN(&input)
	if err != nil{
		Error(c, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, map[string]any{
		"message": "All good",
	})
}
