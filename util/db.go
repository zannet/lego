package util

import (
	"gopkg.in/mgo.v2"
	"github.com/joho/godotenv"
	"fmt"
	// "os"
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
		return session
	}

	session, err = mgo.Dial("localhost:27017")

	if err != nil {
		log.Fatal("Error Connecting Mongo")
	}

	session.SetMode(mgo.Monotonic, true)

	fmt.Printf("DB: creating new db session")

	return session
}



