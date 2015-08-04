package event_model

type (
	// CCEvent contains information for single event.
	CCEvent struct {
		id 				string		`cc:"_id notnull"`
		title 			string		`cc:"title"`
		created_by 		string		`cc:"created_by"`
		created_on 		time.time 	`cc:"created_on"`
		participants 	int 		`cc:"participants"`
		volunteers 		int 		`cc:"volunteers"`
		time 			time.Time 	`cc:"time"`
		duration 		int 		`cc:"duration_in_mins"`
		location 		CCLocation 	`cc:"location"`
		image 			CCImage 	`cc:"image"`
		
	}

	CCImage struct{
		image_url 		string		`cc:"image_url"`
	}
	CCLocation struct{
		locality 		string		`cc:"locality"`
		city 			string		`cc:"city"`
		coordinates 	[]float64 	`cc:"coordinates"`
	}
)