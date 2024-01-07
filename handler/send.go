package handler

import (
	"github.com/aureumapes/chatserver/data"
	"github.com/aureumapes/goutil"
	"github.com/gin-gonic/gin"
)

func Send(context *gin.Context) {
	token := context.PostForm("token")
	var user data.User
	r := data.DB.Raw("SELECT * FROM users WHERE token = ?", token).First(&user)
	if r.RowsAffected == 0 {
		context.Writer.WriteString("Unknown Token: " + token)
	}
	message := context.PostForm("message")
	messageObj := data.Message{
		Chat:    context.Request.URL.Query().Get("chat"),
		Content: message,
		Author:  user.Id,
		Id:      goutil.GenerateId(12),
	}
	data.DB.Create(&messageObj)
}
