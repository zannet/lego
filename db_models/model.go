package models

import (
	"time"
)
type (
    // jsonProgram contains information for single program.
	Program struct {
		Id              string      `json:"_id"`
		Creation_date   time.Time   `json:"creation_date"`
		Title           string      `json:"title"`
		Thread_name     string      `json:"thread_name"`
		Volunteers      string      `json:"volunteers"`
		Genre           string      `json:"genre"`
		Locality        string      `json:"locality"`
		City            string      `json:"city"`
		Image           string      `json:"image"`
		Description     string      `json:"description"`
	}

	User struct {
		Id 				string 		`json:"id"`
		Firstname 		string 		`json:"firstname"`
		Lastname 		string 		`json:"lastname"`
		Email_id		string 		`json:"email_id"`
		Image			string 		`json:"image"`
	}

	Event struct {
		Id 				string		`json:"id"`
		Program_Id 		string		`json:"program_id"`
		Title 			string		`json:"title"`
		Created_by 		User		`json:"created_by"`
		Created_on 		time.Time 	`json:"created_on"`
		Participants 	string 		`json:"participants"`
		Volunteers 		string 		`json:"volunteers"`
		Date 			time.Time 	`json:"date"`
		Duration 		string 		`json:"duration_in_mins"`
		Location 		string 		`json:"location"`
		Image 			string 		`json:"image"`
		
	}

)