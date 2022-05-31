package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/ncroxas/questionaire_api/api/service"
)

func main() {
	var router *gin.Engine
	if ok, _ := strconv.ParseBool(os.Getenv("DEV_MODE")); !ok {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
		service.SetupLogger(router)
		service.SetupRLimit()
	} else {
		router = gin.Default()
	}

	service.SetupCors(router)
	service.SetupSecurity(router)
	service.SetupRoutes(router)

	router.Run(":" + os.Getenv("BACKEND_PORT"))
}
