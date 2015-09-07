package handlers

import (
"fmt"
"time"
"github.com/gin-gonic/gin"
"github.com/naikparag/lego/models"
"github.com/naikparag/lego/util"
"gopkg.in/mgo.v2/bson"
"gopkg.in/mgo.v2"
)

func Load(c *gin.Context,session *mgo.Session) models.CCEvent{

	event := models.CCEvent{}

	if c.Query("id") != "" {
		id := c.Query("id")
		fmt.Println("Id exist in query ")
		event = FetchEvent(id,session)
		fmt.Println("Data fetched")
	}else{
		title 			:= c.Query("title")
		created_by 		:= c.Query("created_by")
		participants	:= c.Query("participants")
		volunteers 		:= c.Query("volunteers")
		date 			:= c.Query("date")
		duration 		:= c.Query("duration")
		location 		:= c.Query("location")
		image 			:= c.Query("image")
		layout 			:= "2006-01-02"
		tm_Date, _ 		:= time.Parse(layout, date)
		// Int variables
		var duration_Int,participants_int,volunteers_int int

		if duration != "" {
			duration_Int = util.StringToInt(duration)
		}
		if participants != "" {
			participants_int = util.StringToInt(participants)
		}
		if volunteers != "" {
			volunteers_int = util.StringToInt(volunteers)
		}

		fmt.Println("Data set to variables")

		events := FetchAllEvent(session)
		var id_str = "1"
		if len(events) > 0{
			id_str = util.Increment(len(events))
		}

		event = models.CCEvent{Id:id_str,Title:title,Created_by:created_by,Created_on:time.Now(),Participants:participants_int,Volunteers:volunteers_int,Date:tm_Date,Duration:duration_Int,Location:location,Image:image}
		fmt.Println("Data set to event model")
	}
	return event	
}


func  FetchEvent(id string,session *mgo.Session)models.CCEvent {
	fmt.Println("Fetching Particular event - ",id)
	// Collection Event
	dbObj := session.DB("lego").C("ccevent")


	result := models.CCEvent{}
	err := dbObj.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		fmt.Println("Unable to fetch data - ",err.Error())
	}
	
	return result

}
func  FetchAllEvent(session *mgo.Session)[]models.CCEvent {
	fmt.Println("Fetching all events - ")
	// Collection Event
	dbObj := session.DB("lego").C("ccevent")

	results := []models.CCEvent{}
	err := dbObj.Find(nil).All(&results)
	if err != nil {
		fmt.Println("Unable to fetch data - ",err.Error())
	}
	
	return results

}