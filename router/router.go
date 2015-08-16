package router

import (
	"github.com/gin-gonic/gin"
	"github.com/naikparag/lego/api"
)

func Register(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Y'all ready for this? \n\nOh no! They were ready for that.")
		//c.JSON(200, gin.H{"Pong": "Ping"})
	})

	// User routes

	userApi := new(api.User)
	r.GET("/user", userApi.Get)

	// Event routes

	r.GET("/event", api.GetEvent)
	r.POST("/event", api.PostEvent)
	
}