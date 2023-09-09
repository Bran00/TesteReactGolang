package router

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize(uri string) {
	router := gin.Default()
	port := os.Getenv("PORT")
	
	initializeRoutes(router, uri)
	router.Run(fmt.Sprintf(":%s", port))
}