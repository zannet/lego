package util

import (
	"gopkg.in/mgo.v2"
	"github.com/joho/godotenv"
	"fmt"
	"os"
	"log"
)

//var session *mgo.Session
var session *mgo.Session

func GetDBSession() *mgo.Session {

//	if session == nil {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if session != nil {
		fmt.Printf("DB: reusing same session")
		return Session
	}

	session, err = mgo.Dial(os.Getenv("MONGO_URL"))

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	fmt.Printf("DB: creating new db session")

	return session.Clone()
}



