package main

import (
	"github.com/naikparag/lego/api"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {

	var HOST = os.Getenv("HOST")
	var PORT = os.Getenv("PORT")

	if PORT == "" {
		PORT = "8888"
	}

	bind := fmt.Sprintf("%s:%s", HOST, PORT)
	fmt.Printf("listening on %s...", bind)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Y'all ready for this? \n\nOh no! They were ready for that.")
		//c.JSON(200, gin.H{"Pong": "Ping"})
	})

	r.GET("/user", api.GetUser)
	r.GET("/event", api.GetEvent)

	r.Run(bind)
}
