package user

import (
	"net/http"

	"2019_1_TheBang/api"

	"github.com/gin-gonic/gin"
)

func MyProfileInfoUpdateHandler(c *gin.Context) {
	token := TokenFromCookie(c.Request)
	info, status := InfoFromCookie(token)
	if status == http.StatusInternalServerError {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	update := api.Update{}
	err := c.BindJSON(&update)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	update.DOB = "2018-01-01"

	profile, status := UpdateUser(info.Nickname, update)
	if status != http.StatusOK {
		c.AbortWithStatus(status)

		return
	}

	bytes, err := profile.MarshalJSON()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.Writer.Write(bytes)

	c.AbortWithStatus(http.StatusOK)
}
