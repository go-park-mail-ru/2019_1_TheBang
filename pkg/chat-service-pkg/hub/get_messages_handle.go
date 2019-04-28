package hub

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func MessagesHandle(c *gin.Context) {
	t := c.Query("timestamp")
	timestamp, err := strconv.Atoi(t)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	messages := GetMessages(timestamp)
	c.JSONP(http.StatusOK, messages)
}
