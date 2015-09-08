package handlers

import (
"fmt"
"time"
"github.com/gin-gonic/gin"
"github.com/naikparag/lego/models"
"github.com/naikparag/lego/util"
"gopkg.in/mgo.v2/bson"
)
type CCEventHandler struct{
	//dbSession *mgo.Session
}

func (this *CCEventHandler)Load() {
	// Load DB Object
	fmt.Println("Loads Db")
	session = util.GetDBSession()
	dbObj = session.DB(util.DB_Name).C(util.DB_Event)
}

func (this *CCEventHandler)AddEvent(c *gin.Context)error{

	event := models.CCEvent{}
	title 			:= c.Query("title")
	program_id 		:= c.Query("program_id")
	created_by 		:= c.Query("created_by")
	participants	:= c.Query("participants")
	volunteers 		:= c.Query("volunteers")
	date 			:= c.Query("date")
	duration 		:= c.Query("duration")
	location 		:= c.Query("location")
	image 			:= c.Query("image")
	layout 			:= "2006-01-02"
	tm_Date, _ 		:= time.Parse(layout, date)
	
	fmt.Println("Data set to variables")

	id_str := util.GenarateId(16, "number")

	event = models.CCEvent{Id:id_str,Program_Id:program_id,Title:title,Created_by:created_by,Created_on:time.Now(),Participants:participants,Volunteers:volunteers,Date:tm_Date,Duration:duration,Location:location,Image:image}
	fmt.Println("Data set to event model")

	fmt.Println("Db session open ")

	err := dbObj.Insert(event)

	if err != nil {
		fmt.Println("Data insertion failed ",err.Error())
		panic(err)
		return err
	}
	
	return nil	
}


func  (this *CCEventHandler)FetchEvent(id string)models.CCEvent {
	fmt.Println("Fetching Particular event - ",id)
	result := models.CCEvent{}
	err := dbObj.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		fmt.Println("Unable to fetch data - ",err.Error())
	}
	
	return result
}
func  (this *CCEventHandler)FetchAllEvent(c *gin.Context)[]models.CCEvent {
	fmt.Println("Fetching all events - ")
	filter := util.ExtractEventFilter(c)
	results := []models.CCEvent{}
	err := dbObj.Find(filter).All(&results)
	if err != nil {
		fmt.Println("Unable to fetch data - ",err.Error())
	}
	
	return results
}
func  (this *CCEventHandler)DeleteEvent(c *gin.Context) {
	fmt.Println("Deleting  events - ")
	filter := util.ExtractEventFilter(c)
	err := dbObj.Remove(filter)
	if err != nil {
		fmt.Println("Unable to remove data - ",err.Error())
	}
}






