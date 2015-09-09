package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/naikparag/lego/handlers"
)
var eventHandler *handlers.CCEventHandler
type CCEvent struct{
	// User routes
}
func (this *CCEvent) Load() {
	if eventHandler == nil {
		eventHandler = new(handlers.CCEventHandler)
	}
	eventHandler.Load()
}
func (this *CCEvent) Get(c *gin.Context) {
	this.Load()
	results := eventHandler.FetchAllEvent(c)
	fmt.Println("Results All: ", results)
	// fmt.Println(util.Glo_variable)
	c.JSON(200, results)	
}

func (this *CCEvent) Post(c *gin.Context) {
	this.Load()
	program_id 	:= c.Query("program_id")
	user_id 	:= c.Query("created_by")

	if user_id != "" {
		_,err := userHandler.FetchUser(user_id)
		if err != nil {
			fmt.Println("User id not exist: ")
			panic(err)
			return
		}
	}
	if program_id != "" {
		_,err := programHandler.FetchProgram(program_id)
		if err != nil {
			fmt.Println("Program id not exist: ")
			panic(err)
			return
		}
	}
	err := eventHandler.AddEvent(c)
	if err != nil {
		fmt.Println("Data insertion failed ",err.Error())
		return
	}
	fmt.Println("Data Inserted Successfully: ")

}

func (this *CCEvent) Delete(c *gin.Context) {
	this.Load()
	eventHandler.DeleteEvent(c)
}