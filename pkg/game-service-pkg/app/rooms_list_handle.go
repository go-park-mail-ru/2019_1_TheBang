package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoomsListHandle(c *gin.Context) {
	rooms := AppInst.WrappedRoomsList()
	c.JSONP(http.StatusOK, rooms)
}
