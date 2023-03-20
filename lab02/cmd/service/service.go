package main

import (
	"github.com/gin-gonic/gin"
	"lab02/pkg/apis"
	"lab02/pkg/config"
)

func main() {
	router := gin.Default()

	router.GET(config.DefaultPath, apis.HomePage)
	router.GET(config.WeatherPath, apis.HandleRequest)

	router.Run(config.Localhost)
}
