package models
import (
  "time"
)

type (
    // CCProgram contains information for single program.
    CCProgram struct {
      Id              string      `cc:"_id notnull"`
      Creation_date   time.Time   `cc:"creation_date"`
      Title           string      `cc:"title"`
      Thread_name     string      `cc:"thread_name"`
      Volunteers      string         `cc:"volunteers"`
      Genre           string      `cc:"genre"`
      Locality        string      `cc:"locality"`
      City            string      `cc:"city"`
      Image           string      `cc:"image"`
      Description     string      `cc:"description"`
    }
   
)
