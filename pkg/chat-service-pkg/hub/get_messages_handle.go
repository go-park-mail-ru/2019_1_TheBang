package hub

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"fmt"
)

func MessagesHandle(c *gin.Context) {
	t := c.Query("timestamp")
	timestamp, err := strconv.Atoi(t)
	if err != nil {
		fmt.Println(t)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	messages := GetMessages(timestamp)
	fmt.Println(messages)
	c.JSONP(http.StatusOK, messages)
}
