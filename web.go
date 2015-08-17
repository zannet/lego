package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/naikparag/lego/router"
	"os"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var HOST = os.Getenv("HOST")
	var PORT = os.Getenv("PORT")

	bind := fmt.Sprintf("%s:%s", HOST, PORT)

	r := gin.Default()
	router.Register(r)
	//util.InitDB()
	r.Run(bind)
}
