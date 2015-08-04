package model

type program struct{
  id            string
  creation_date time.Time
  title         string
  thread_name   string
  volunteers    int
  genre         string
  location      location
  image         image
  description   string
}
