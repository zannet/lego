package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/naikparag/lego/handlers"
)
var userHandler *handlers.CCUserHandler

type CCUser struct{
	// User routes
}
func (this *CCUser) Load() {
	if eventHandler == nil {
		userHandler = new(handlers.CCUserHandler)
	}
	userHandler.Load()

}
func (this *CCUser) Get(c *gin.Context) {
	this.Load()
	results := userHandler.FetchAllUsers(c)
	fmt.Println("Results All: ", results)
	// fmt.Println(util.Glo_variable)
	c.JSON(200, results)	
}

func (this *CCUser) Post(c *gin.Context) {
	this.Load()
	err := userHandler.AddUser(c)
	if err != nil {
		fmt.Println("Data insertion failed ",err.Error())
		return
	}
	fmt.Println("User Created Successfully: ")

}

func (this *CCUser) Delete(c *gin.Context) {
	this.Load()
	userHandler.DeleteUser(c)
}