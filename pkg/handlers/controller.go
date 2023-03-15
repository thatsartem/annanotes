package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	api := r.Group("/api")
	routes := api.Group("/notes")
	routes.POST("/", h.AddNote)
	routes.GET("/", h.GetNotes)
	routes.GET("/:id", h.GetNote)
	routes.PUT("/:id", h.UpdateNote)
	routes.DELETE("/:id", h.DeleteNote)
	user_router := api.Group("/users")
	user_router.POST("/", h.AddUser)
	user_router.GET("/", h.GetUsers)
	user_router.GET("/:id/notes", h.GetUserNotes)
}
