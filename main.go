package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/scholzie/go-url-shortener/handler"
	"github.com/scholzie/go-url-shortener/store"
)

func main()  {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	r.POST("/create-short-url", func(ctx *gin.Context) {
		handler.CreateShortcutUrl(ctx)
	})

	r.GET("/:shortUrl", func(ctx *gin.Context) {
		handler.HandleShortUrlRedirect(ctx)
	})

	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}