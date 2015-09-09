package handlers

import (
"fmt"
"github.com/gin-gonic/gin"
"github.com/naikparag/lego/db_models"
"github.com/naikparag/lego/util"
"gopkg.in/mgo.v2/bson"
"github.com/oleiade/reflections"
)
type CCUserHandler struct{
	//dbSession *mgo.Session
}

func (this *CCUserHandler)Load() {
	// Load DB Object
	fmt.Println("Loads Db")
	session = util.GetDBSession()
	dbObj = session.DB(util.DB_Name).C(util.DB_User)
}
func (this *CCUserHandler)LoadModel(user models.CCUser)models.User {
	fields, _ := reflections.Fields(&models.User{})
//new instance
	userModel := models.User{}

	for i := 0; i < len(fields); i++ {
		has, _ := reflections.HasField(user, fields[i])
		if has == true {
			fmt.Println("Field Exist------",fields[i])
			
			value, err := reflections.GetField(user, fields[i])
			if err == nil {
				fmt.Println("Field Value------",value)
				setError := reflections.SetField(&userModel, fields[i], value)  // err != nil

				if setError != nil {
					panic(setError)
				}

			}

		}

	}
	fmt.Println("Data set to Model --- ", userModel)
	return userModel
}

func (this *CCUserHandler)AddUser(c *gin.Context)error{

	user := models.CCUser{}
	firstName 		:= c.Query("firstname")
	lastName 		:= c.Query("lastname")
	username		:= c.Query("username")
	password 		:= c.Query("password")
	email_id 		:= c.Query("email_id")
	image 			:= c.Query("image")
	
	password_enc 	:= util.Encrypt(password)

	fmt.Println("Data set to variables")

	id_str := util.GenarateId(16, "number")

	user = models.CCUser{Id:id_str,Firstname:firstName,Lastname:lastName,Username:username,Password:password_enc,Email_id:email_id,Image:image}
	fmt.Println("Data set to user model")

	fmt.Println("Db session open ")

	err := dbObj.Insert(user)

	if err != nil {
		fmt.Println("Data insertion failed ",err.Error())
		panic(err)
		return err
	}
	
	return nil	
}


func  (this *CCUserHandler)FetchUser(id string)(models.CCUser,error) {
	this.Load()
	fmt.Println("Fetching Particular User - ",id)
	result := models.CCUser{}
	err := dbObj.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		fmt.Println("Unable to fetch data - ",err.Error())
		return result,err
	}
	return result,err
}
func  (this *CCUserHandler)FetchAllUsers(c *gin.Context)[]models.User {
	fmt.Println("Fetching all users - ")
	filter := util.ExtractUserFilter(c)
	results := []models.CCUser{}
	err := dbObj.Find(filter).All(&results)
	if err != nil {
		fmt.Println("Unable to fetch data - ",err.Error())
	}


	var users []models.User
	for i := 0; i < len(results); i++ {
		userModel := this.LoadModel(results[i])
		users = append(users,userModel)
	}
	return users
}
func  (this *CCUserHandler)DeleteUser(c *gin.Context) {
	fmt.Println("Deleting  events - ")
	filter := util.ExtractEventFilter(c)
	err := dbObj.Remove(filter)
	if err != nil {
		fmt.Println("Unable to remove data - ",err.Error())
	}
}