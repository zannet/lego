package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/naikparag/lego/util"
	"github.com/naikparag/lego/handlers"
)
type CCEvent struct{
	//dbSession *mgo.Session
}
func (this *CCEvent) Get(c *gin.Context) {

	s := util.GetDBSession()
	results := handlers.FetchAllEvent(s)
	fmt.Println("Results All: ", results)
	// fmt.Println(util.Glo_variable)
	c.JSON(200, results)

	
}

func (this *CCEvent) Post(c *gin.Context) {

	// Insert Datas
	session := util.GetDBSession()
	
	newEvent := handlers.Load(c,session)
	fmt.Println("Data loaded to model: ")

	// Collection Event
	dbObj := session.DB("lego").C("ccevent")

	fmt.Println("Db session open ")

	err := dbObj.Insert(newEvent)

	if err != nil {
		fmt.Println("Data insertion failed ",err.Error())
	}
	
	fmt.Println("Data Inserted Successfully: ")

}

func (this *CCEvent) Delete(c *gin.Context) {
	session := util.GetDBSession()
	
	session.DB("lego").C("ccevent").RemoveAll(nil)

}