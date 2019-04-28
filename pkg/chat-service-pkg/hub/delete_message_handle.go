package hub

import (
	"2019_1_TheBang/api"
	"github.com/gin-gonic/gin"
	// 	"net/http"
)

func DeleteMessageHandle(c *gin.Context) {
	msg := api.ChatMessage{}
	c.BindJSON(&msg)
}
