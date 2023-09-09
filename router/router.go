package router

import "github.com/gin-gonic/gin"

func Initialize(uri string) {
	router := gin.Default()

	initializeRoutes(router, uri)
	router.Run(":8080")
}