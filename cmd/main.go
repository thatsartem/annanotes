package main

import (
	"annanotes/pkg/db"
	"annanotes/pkg/handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	db := db.Init("test.db")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ping": "pong",
		})
	})

	handlers.RegisterRoutes(r, db)

	log.Fatal(http.ListenAndServe(":8080", r))
}
