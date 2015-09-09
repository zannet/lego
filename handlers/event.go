package handlers

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/naikparag/lego/db_models"
	"github.com/naikparag/lego/util"
	"gopkg.in/mgo.v2/bson"
    "github.com/oleiade/reflections"

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

func (this *CCEventHandler)LoadModel(event models.CCEvent)models.Event {
	fields, _ := reflections.Fields(&models.Event{})
//new instance
	eventModel := models.Event{}

	for i := 0; i < len(fields); i++ {
		has, _ := reflections.HasField(event, fields[i])
		if has == true {
			fmt.Println("Field Exist------",fields[i])
			if fields[i] == "Created_by" {
				var userObj CCUserHandler
				var usermodel  models.CCUser
				value, err := reflections.GetField(event, fields[i])

				if err == nil && value != nil {
					str, _ := value.(string)
					usermodel,err = userObj.FetchUser(str)
					if err == nil  {
						responseModel := userObj.LoadModel(usermodel)
						setError := reflections.SetField(&eventModel, fields[i], responseModel)  // err != nil
						if setError != nil {
						panic(setError)
						}
					}
				}
			}else{
				value, err := reflections.GetField(event, fields[i])
				if err == nil {
					fmt.Println("Field Value------",value)
				setError := reflections.SetField(&eventModel, fields[i], value)  // err != nil

				if setError != nil {
					panic(setError)
				}
			}
		}

	}
	
}
fmt.Println("Data set to Model --- ", eventModel)
return eventModel
}

func (this *CCEventHandler)AddEvent(c *gin.Context)error{
	this.Load()
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


func  (this *CCEventHandler)FetchEvent(id string)models.Event {
	this.Load()
	fmt.Println("Fetching Particular event - ",id)
	result := models.CCEvent{}
	err := dbObj.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		fmt.Println("Unable to fetch data - ",err.Error())
	}
	eventModel := this.LoadModel(result)
	return eventModel
}
func  (this *CCEventHandler)FetchAllEvent(c *gin.Context)[]models.Event {
	this.Load()
	fmt.Println("Fetching all events - ")
	filter := util.ExtractEventFilter(c)
	results := []models.CCEvent{}
	err := dbObj.Find(filter).All(&results)
	if err != nil {
		fmt.Println("Unable to fetch data - ",err.Error())
	}
	var events []models.Event
	for i := 0; i < len(results); i++ {
		eventModel := this.LoadModel(results[i])
		events = append(events,eventModel)
	}
	return events
}
func  (this *CCEventHandler)DeleteEvent(c *gin.Context) {
	this.Load()
	fmt.Println("Deleting  events - ")
	filter := util.ExtractEventFilter(c)
	err := dbObj.Remove(filter)
	if err != nil {
		fmt.Println("Unable to remove data - ",err.Error())
	}
}






