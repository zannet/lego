package api

import (
	"github.com/gin-gonic/gin"
	"github.com/naikparag/lego/util"
	"github.com/naikparag/lego/models"
	"fmt"
)

type User struct{
	//dbSession *mgo.Session
}

func (this *User) Get(c *gin.Context) {

	s := util.GetDBSession()
	defer s.Close()

	results := []models.User{}

	err := s.DB("lego").C("user").Find(nil).All(&results)
	if err != nil {
	    fmt.Println("some error ", err)
	}
	fmt.Println("Results All: ", results)
	fmt.Println(util.Glo_variable)
	c.JSON(200, results)
}

func (this *User) Post(c *gin.Context) {
	//c.JSON(200, gin.H{"user_name": "Mojo Jojo"})

	s := util.GetDBSession()
	defer s.Close()

	newUser := models.User{"Tin", "Man"}
	s.DB("lego").C("user").Insert(newUser)
}
