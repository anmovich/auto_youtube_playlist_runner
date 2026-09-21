package handlers

import (
	"ypp/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *service.Service
}

func (h *Handler)InitRoute() *gin.Engine{
	router := gin.New()
	
	api := router.Group("/api")
	{
		api.POST("/", h.NewChannel)
		api.GET("/", h.GetChannel)
		api.PATCH("/:id", h.PatchChannel)
		api.DELETE("/:id", h.DeleteChannel)
	}
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SingUP)
		auth.POST("/sign-in", h.SingIN)
	}
	router.GET("/ping", h.PingHandler)
	return router
}
