package handlers

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/naikparag/lego/db_models"
	"github.com/naikparag/lego/util"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session
var dbObj *mgo.Collection

type CCProgramHandler struct{
	
}

func (this *CCProgramHandler)Load() {
	// Load DB Object
	fmt.Println("Loads Db")
	session = util.GetDBSession()
	dbObj = session.DB(util.DB_Name).C(util.DB_Program)
}
func (this *CCProgramHandler)AddProgram(c *gin.Context) error{

	program := models.CCProgram{}
	title 			:= c.Query("title")
	thread_name 	:= c.Query("thread_name")
	genre			:= c.Query("genre")
	volunteers 		:= c.Query("volunteers")
	locality 		:= c.Query("locality")
	city 			:= c.Query("city")
	image 			:= c.Query("image")
	description 	:= c.Query("description")
		// Int variables
	

	fmt.Println("Data set to variables")

		// Generate unique id
	id_str := util.GenarateId(16, "number")

	program = models.CCProgram{
		Id:id_str,Creation_date:time.Now(),Title:title,Thread_name:thread_name,Volunteers:volunteers,Genre:genre,Locality:locality,City:city,Image:image,Description:description,
	}
	fmt.Println("Data set to program model")
	fmt.Println("Db session open ")
	err := dbObj.Insert(program)

	if err != nil {
		fmt.Println("Data insertion failed ",err.Error())
		panic(err)
		return err
	}

	return nil	
}


func  (this *CCProgramHandler)FetchProgram(id string)(models.CCProgram,error) {
	this.Load()
	fmt.Println("Fetching Particular Program - ",id)
	result := models.CCProgram{}

	err := dbObj.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		fmt.Println("Unable to fetch data - ",err.Error())
		return result,err
	}
	
	return result,err

}
func  (this *CCProgramHandler)FetchAllPrograms(c *gin.Context)[]models.CCProgram {
	fmt.Println("Fetching all events - ")
	filter := util.ExtractProgramFilter(c)
	var err error

	results := []models.CCProgram{}
	if filter != nil {
		fmt.Println("Reached - ")
		err = dbObj.Find(filter).All(&results)
	}else{
		fmt.Println("Reached to All- ")
		err = dbObj.Find(nil).All(&results)
	}
	
	if err != nil {
		fmt.Println("Unable to fetch data - ",err.Error())
	}
	
	return results

}

func (this *CCProgramHandler)DeleteProgram(c *gin.Context) {
	fmt.Println("Deleting  Programs - ")
	session := util.GetDBSession()
	dbObj := session.DB(util.DB_Name).C(util.DB_Program)

	filter := util.ExtractProgramFilter(c)
	err := dbObj.Remove(filter)
	if err != nil {
		fmt.Println("Unable to remove data - ",err.Error())
	}
}













