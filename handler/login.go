package handler

import (
	"encoding/base64"
	"github.com/aureumapes/chatserver/data"
	"github.com/aureumapes/goutil"
	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")
	token := base64.StdEncoding.EncodeToString([]byte(goutil.Hash(username+password, "", "")))
	var result data.User
	if r := data.DB.Raw("SELECT username,token FROM users WHERE token = ?", token).First(&result); r.RowsAffected != 1 {
		context.Writer.WriteString("Username/Password combination not registered in database")
		return
	}
	context.Writer.WriteString(token)
}
