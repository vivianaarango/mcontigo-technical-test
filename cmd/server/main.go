package main

import (
	"os"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/cors"
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter/handler"
	"github.com/gin-gonic/gin"
)

// @contact.name                Grupo MContigo
// @title                       Newsletter API
// @version                     1.0
// @description                 Newsletter API
func main() {
	r := gin.Default()

	r.Use(cors.AllowCORS())

	newsletterHandler := handler.Build()

	libGroup := r.Group("/newsletter")
	libGroup.GET("", newsletterHandler.Get)

	err := r.Run(":" + os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}
