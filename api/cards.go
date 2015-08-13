package api

import (
	"fmt"
	"io/ioutil"
    "os"
    "encoding/json"
   	"net/http"
	"github.com/gin-gonic/gin"
)
type jsonobject struct {
    Cards []ObjectType
}

type ObjectType struct {
    Data CardType
}

type CardType struct {
    Text   string
    Image_url   string
    Author   AuthorType
}

type AuthorType struct {
    Name     string
    Message string
    Image_url    string
}

func GetCards(c *gin.Context) {
    file, e := ioutil.ReadFile("./Json/cards.json")
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))

    //m := new(Dispatch)
    //var m interface{}
    var jsontype jsonobject
    json.Unmarshal(file, &jsontype)
	c.JSON(http.StatusOK, jsontype)
}
