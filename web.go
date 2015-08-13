package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"os"
	"./api"
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

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Y'all ready for this? \n\nOh no! They were ready for that.")
		//c.JSON(200, gin.H{"Pong": "Ping"})
	})

    r.GET("/user", api.GetUser)
    r.GET("/event", api.GetEvent)
    r.GET("/cards", api.GetCards)

    r.Run(bind)
}
