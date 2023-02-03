package main

import (
	"log"

	"backend/config"
	"backend/middlewares"
	ctrl "backend/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	if err := config.ConnectDB(); err != nil {
		log.Panic("Cannot connect database ", err)
	}
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(middlewares.CORSMiddleware())

	ctrl.Channels(router.Group("/channels"))
	ctrl.Template(router.Group("/template"))
	ctrl.Savety(router.Group("/savety"))
	ctrl.Saveunit(router.Group("/saveunit"))

	router.Run(":8888")
}
