package hub

import (
	"2019_1_TheBang/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EditMessageHandle(c *gin.Context) {
	msg := api.ChatMessage{}
	c.BindJSON(&msg)

	err := EditMessage(msg)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}
}
