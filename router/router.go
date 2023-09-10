package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Initialize(uri string) *gin.Engine {
	router := gin.Default()

	// Configurar o middleware CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"} // Permitir qualquer origem
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	corsConfig.AllowHeaders = []string{
		"X-CSRF-Token",
		"X-Requested-With",
		"Accept",
		"Accept-Version",
		"Content-Length",
		"Content-MD5",
		"Content-Type",
		"Date",
		"X-Api-Version",
		"Authorization",
	}

	router.Use(cors.New(corsConfig)) // Use o middleware CORS configurado
	initializeRoutes(router, uri)

	return router
}
