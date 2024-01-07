package main

import (
	"github.com/aureumapes/chatserver/data"
	"github.com/aureumapes/chatserver/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	data.InitDB()
	r := gin.New()
	r.Handle(http.MethodPost, "/register", handler.Register)
	r.Handle(http.MethodPost, "/login", handler.Login)
	r.Handle(http.MethodPost, "/send", handler.Send)
	r.Handle(http.MethodGet, "/load", handler.LoadChat)
	r.Handle(http.MethodGet, "/chats", handler.Chats)
	r.Run("0.0.0.0:8080")
}
