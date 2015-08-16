package main

import (
	"github.com/naikparag/lego/router"
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

	r := gin.Default()	
	router.Register(r)

	r.Run(bind)
}
