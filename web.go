package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"os"
)

func main() {

	var HOST = os.Getenv("HOST")
	var PORT = os.Getenv("PORT")

	if (PORT == "") {
		PORT = "8888"
	}

	bind := fmt.Sprintf("%s:%s", HOST, PORT)
	fmt.Printf("listening on %s...", bind)

	r := gin.Default()

    r.GET("/test", func(c *gin.Context) {
    	c.JSON(200, gin.H{"Pong": "Ping"})
    })

    r.Run(bind)
}
