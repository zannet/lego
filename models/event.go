package models

import (
	"time"
	// "labix.org/v2/mgo/bson"
)
	// CCEvent contains information for single event.

type CCEvent struct {
		Id 				string		`json:"id"`
		Program_Id 		string		`json:"program_id"`
		Title 			string		`json:"title"`
		Created_by 		string		`json:"created_by"`
		Created_on 		time.Time 	`json:"created_on"`
		Participants 	string 		`json:"participants"`
		Volunteers 		string 		`json:"volunteers"`
		Date 			time.Time 	`json:"date"`
		Duration 		string 		`json:"duration_in_mins"`
		Location 		string 		`json:"location"`
		Image 			string 		`json:"image"`
		
	}



