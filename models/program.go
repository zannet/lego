package models
import (
  "time"
)

type (
    // CCProgram contains information for single program.
    CCProgram struct {
      id              string      `cc:"_id notnull"`
      creation_date   time.Time   `cc:"creation_date"`
      title           string      `cc:"title"`
      thread_name     string      `cc:"thread_name"`
      volunteers      int         `cc:"volunteers"`
      genre           string      `cc:"genre"`
      location        CCLocation  `cc:"location"`
      image           string      `cc:"image"`
      description     string      `cc:"description"`
    }
    CCLocation struct{
      locality        string      `cc:"locality"`
      city            string      `cc:"city"`
      coordinates     []float64   `cc:"coordinates"`
    }
)
