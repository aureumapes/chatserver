package handler

import (
	"encoding/base64"
	"github.com/aureumapes/chatserver/data"
	"github.com/aureumapes/goutil"
	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")
	token := base64.StdEncoding.EncodeToString([]byte(goutil.Hash(username+password, "", "")))
	var result data.User
	if r := data.DB.Raw("SELECT username FROM users WHERE username = ?", username).First(&result); r.RowsAffected != 0 {
		context.Writer.WriteString("Pick another Username")
		return
	}
	user := data.User{
		Username: username,
		Token:    token,
		Role:     data.MEMBER,
		Id:       goutil.GenerateId(12),
	}
	data.DB.Create(&user)
	context.Writer.WriteString("Successfully created your Account")
}
