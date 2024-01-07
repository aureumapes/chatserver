package handler

import (
	"fmt"
	"github.com/aureumapes/chatserver/data"
	"github.com/gin-gonic/gin"
	"strings"
)

func LoadChat(context *gin.Context) {
	chatID := context.Request.URL.Query().Get("id")
	var messages []data.Message
	data.DB.Raw("SELECT * FROM messages WHERE chat = ?", chatID).Scan(&messages)
	var result []string
	var author data.User
	for _, message := range messages {
		data.DB.Raw("SELECT * FROM users WHERE id = ?", message.Author).First(&author)
		result = append(result, fmt.Sprintf("[%s] %s >> %s", author.Role, author.Username, message.Content))
	}
	context.Writer.WriteString(strings.Join(result, "\n"))
}

func reverse(input []data.Message) []data.Message {
	if len(input) == 0 {
		return input
	}
	return append(reverse(input[1:]), input[0])
}
