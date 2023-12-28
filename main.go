package main

import (
	"bookend/controllers"
	"bookend/inits"
	"bookend/middlewares"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	r := gin.Default()

	err := r.SetTrustedProxies(nil)
	if err != nil {
		return
	}

	r.Use(middlewares.CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Bookend API",
		})
	})

	r.POST("/v1/entry", controllers.CreateEntry)
	r.GET("/v1/entry", controllers.GetEntries)
	r.GET("/v1/entry/:id", controllers.GetEntry)

	runtimeErr := r.Run()
	if runtimeErr != nil {
		log.Fatalf("Error, failed to start gin server. %v", runtimeErr)
		return
	}
}
