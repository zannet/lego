package models
import (
  "time"
)

type (
    // CCProgram contains information for single program.
    CCProgram struct {
      Id              string      `json:"id"`
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
   
)
