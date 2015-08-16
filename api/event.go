package api

import (
	"github.com/gin-gonic/gin"
	"github.com/naikparag/lego/models"
)

var events = []models.User{
		models.User{"Mojo", "Jojo"},
		models.User{"Gur", "Victor"},
		models.User{"Johny", "Bravo"},
	}

func GetEvent(c *gin.Context) {
	c.JSON(200, gin.H{"users": events})
}

func PostEvent(c *gin.Context) {
	events = append(events, models.User{"Money", "man"})
	c.JSON(200, gin.H{"yes":"Done"})
}