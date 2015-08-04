package model

type event struct {
	id 				string
	title 			string
	created_by 		string
	created_on 		time.Time
	participants 	int
	volunteers 		int
	time 			time.Time
	duration 		int
	location 		location
	image 			image
}

type location struct{
	locality string
	city string
	coordinates coordinates
}
type coordinates struct{
	latitude float
	longitude float
}
type image struct{
	image_url string
}

type (
	// CCEvent contains information for single event.
	CCEvent struct {
	id 				string		`cc:"_id notnull"`
	title 			string		`cc:"wind_speed_milehour"`
	created_by 		string		`cc:"wind_speed_milehour"`
	created_on 		time.time 	`cc:"wind_speed_milehour"`
	participants 	int 		`cc:"wind_speed_milehour"`
	volunteers 		int 		`cc:"wind_speed_milehour"`
	time 			time.Time 	`cc:"wind_speed_milehour"`
	duration 		int 		`cc:"wind_speed_milehour"`
	location 		location 	`cc:"wind_speed_milehour"`
	image 			image 		`cc:"wind_speed_milehour"`
		
	}

	// BuoyLocation contains the buoy's location.
	BuoyLocation struct {
		Type        string    `cc:"type"`
		Coordinates []float64 `cc:"coordinates"`
	}

	// BuoyStation contains information for an individual station.
	BuoyStation struct {
		ID        cc.ObjectId `cc:"_id,omitempty"`
		StationId string        `cc:"station_id"`
		Name      string        `cc:"name"`
		LocDesc   string        `cc:"location_desc"`
		Condition BuoyCondition `cc:"condition"`
		Location  BuoyLocation  `cc:"location"`
	}
)