package main

import (
	"github.com/gin-gonic/gin"
	"mooc/chap22/proto"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/moreJSON", moreJSON)
	router.GET("/someProtoBuf", returnProto)

	router.Run(":8083")
}

func returnProto(c *gin.Context) {
	course := []string{"python", "go", "java"}
	user := &proto.Teacher{
		Name:   "aa",
		Course: course,
	}
	c.ProtoBuf(http.StatusOK, user)
}

func moreJSON(c *gin.Context) {
	var msg struct {
		name    string `json:"user"`
		Message string
		Number  int
	}
	msg.name = "msg_name"
	msg.Message = "msg_Message"
	msg.Number = 20

	c.JSON(http.StatusOK, msg)
}
