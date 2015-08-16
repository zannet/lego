package api

import (
	"github.com/gin-gonic/gin"
)

type User struct{

}

func (this *User) Get(c *gin.Context) {
	c.JSON(200, gin.H{"user_name": "Mojo Jojo"})
}