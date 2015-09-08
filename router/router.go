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
	r.POST("/user", userApi.Post)

	// Event routes
	eventApi := new(api.CCEvent)
	r.GET("/event", eventApi.Get)
	r.POST("/event", eventApi.Post)	
	r.DELETE("/event", eventApi.Delete)

	// Program routes
	programApi := new(api.CCProgram)
	r.GET("/program", programApi.Get)
	r.POST("/program", programApi.Post)	
	r.DELETE("/program", programApi.Delete)
}