package handler

import (
	"fmt"
	"github.com/aureumapes/chatserver/data"
	"github.com/gin-gonic/gin"
)

func Chats(context *gin.Context) {
	var chats []data.Chat
	data.DB.Raw("SELECT * FROM chats").Scan(&chats)
	result := "ID\t\tName\n"
	for _, chat := range chats {
		result += fmt.Sprintf("%s\t\t%s\n", chat.Id, chat.Name)
	}
	context.Writer.WriteString(result)
}
