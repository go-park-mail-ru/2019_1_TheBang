package hub

import (
	"2019_1_TheBang/api"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteMessageHandle(c *gin.Context) {
	msg := api.ChatMessage{}
	c.BindJSON(&msg)

	err := DeleteMessage(msg.Timestamp, msg.Author)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	msg.Deleted = true
	bytes, err := json.Marshal(msg)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	HubInst.Broadcast <- bytes
}
