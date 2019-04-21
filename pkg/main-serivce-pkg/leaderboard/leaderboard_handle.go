package leaderboard

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LeaderbordHandler(c *gin.Context) {
	per := c.Param("page")

	page, err := strconv.Atoi(per)
	if page == 0 || err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	profs, status := LeaderPage(uint(page))
	if status != http.StatusOK {
		c.AbortWithStatus(status)
		fmt.Println("HETE")

		return
	}

	c.JSONP(http.StatusOK, profs)
}
