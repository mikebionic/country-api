package internal

import (
	"countryApi/pkg/countryLib"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitApp() *gin.Engine {
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
		file, _ := os.Create("gin.log")
		gin.DefaultWriter = io.MultiWriter(file)
	}

	router := gin.New()
	router.SetTrustedProxies(nil)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"*"},
		MaxAge:          12 * time.Hour,
	}))
	router.Use(func(ctx *gin.Context) {
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	})

	log.SetOutput(gin.DefaultWriter)
	countryLib.CountryLib(router)

	return router
}
