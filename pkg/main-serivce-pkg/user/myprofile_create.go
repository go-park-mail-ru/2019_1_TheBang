package user

import (
	"net/http"

	"2019_1_TheBang/api"

	"github.com/gin-gonic/gin"
)

func MyProfileCreateHandler(c *gin.Context) {
	signup := &api.Signup{}
	err := c.BindJSON(signup)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	signup.DOB = "2018-01-01"

	status := CreateUser(signup)
	if status != http.StatusCreated {
		c.AbortWithStatus(status)

		return
	}

	c.Status(http.StatusCreated)
}
