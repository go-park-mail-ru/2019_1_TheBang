package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MyProfileInfoHandler(c *gin.Context) {
	token := TokenFromCookie(c.Request)
	info, status := InfoFromCookie(token)
	if status == http.StatusInternalServerError {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	profile, status := SelectUser(info.Nickname)
	if status != http.StatusOK {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	bytes, err := profile.MarshalJSON()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.Writer.Write(bytes)

	c.AbortWithStatus(http.StatusOK)
}
