package main

import (
	"fmt"
	"gopro/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Initializin")
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDB()
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})


	r.POST("/signup", controllers.SignUp){
		// Implemented the sign up logic here
		//get email and password of req body

	}
	r.Run() // listen and serve on 0.0.0.0:8080
}
