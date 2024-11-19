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
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	log.SetOutput(gin.DefaultWriter)
	countryLib.CountryLib(router)

	return router
}
