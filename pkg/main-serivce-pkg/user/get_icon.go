package user

import (
	"io/ioutil"
	"net/http"

	"os"

	"github.com/gin-gonic/gin"
)

func GetIconHandler(c *gin.Context) {
	c.Header("Content-Type", "image/jpeg")

	filename := c.Param("filename")

	root, _ := os.Getwd()

	filepath := root + "/tmp/" + filename
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	_, err = c.Writer.Write(data)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}
}
