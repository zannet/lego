package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/naikparag/lego/handlers"
)

var programHandler  *handlers.CCProgramHandler
type CCProgram struct{
	// User routes
}
func (this *CCProgram) Load() {
	fmt.Println("Program Api loaded: ")
	if programHandler == nil {
		programHandler = new(handlers.CCProgramHandler)
		programHandler.Load()
	}
}
func (this *CCProgram) Get(c *gin.Context) {
	this.Load()
	results := programHandler.FetchAllPrograms(c)
	fmt.Println("Results All: ", results)
	// fmt.Println(util.Glo_variable)
	c.JSON(200, results)
}

func (this *CCProgram) Post(c *gin.Context) {
	this.Load()	
	err := programHandler.AddProgram(c)
	if err != nil {
		fmt.Println("Data insertion failed ",err.Error())
		return
	}
	fmt.Println("Data Inserted Successfully: ")

}

func (this *CCProgram) Delete(c *gin.Context) {
	this.Load()
	programHandler.DeleteProgram(c)

}