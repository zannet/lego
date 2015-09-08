package util
import (
	"fmt"
	// "encoding/json"
	"github.com/gin-gonic/gin"
)
var filter map[string]string
func ExtractProgramFilter(c *gin.Context)map[string]string {

	// filter = make(map[string]string)


	if c.Query("id") != "" {
		filter["id"] = c.Query("id")
	}
	if c.Query("title") != "" {
		filter["title"] = c.Query("title")

	}
	if c.Query("thread_name") != "" {
		filter["thread_name"] = c.Query("thread_name")
	}
	if c.Query("genre") != "" {
		filter["genre"] = c.Query("genre")
	}
	if c.Query("volunteers") != "" {
		filter["volunteers"] = c.Query("volunteers")
	}
	if c.Query("locality") != "" {
		filter["locality"] = c.Query("locality")
	}
	if c.Query("city") != "" {
		filter["city"] = c.Query("city")
	}
	if c.Query("image") != "" {
		filter["image"] = c.Query("image")
	}
	if c.Query("description") != "" {
		filter["description"] = c.Query("description")
	}
	fmt.Println("Filter - ", filter)
	return filter
}

func ExtractEventFilter(c *gin.Context)map[string]string {

	// filter = make(map[string]string)


	if c.Query("id") != "" {
		filter["id"] = c.Query("id")
	}
	if c.Query("title") != "" {
		filter["title"] = c.Query("title")

	}
	if c.Query("program_id") != "" {
		filter["program_id"] = c.Query("program_id")
	}
	if c.Query("created_by") != "" {
		filter["created_by"] = c.Query("created_by")
	}
	if c.Query("participants") != "" {
		filter["participants"] = c.Query("participants")
	}
	if c.Query("volunteers") != "" {
		filter["volunteers"] = c.Query("volunteers")
	}
	if c.Query("date") != "" {
		filter["date"] = c.Query("date")
	}
	if c.Query("duration") != "" {
		filter["duration"] = c.Query("duration")
	}
	if c.Query("location") != "" {
		filter["location"] = c.Query("location")
	}
	fmt.Println("Filter - ", filter)
	return filter

}