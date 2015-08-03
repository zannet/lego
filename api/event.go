package api

import (
	"github.com/gin-gonic/gin"
)

func GetEvent(c *gin.Context) {
	c.JSON(200, gin.H{"event_name": "Mojo Jojo evil party"})
}