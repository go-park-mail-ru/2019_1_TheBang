package logout

import (
	"net/http"
	"time"

	"2019_1_TheBang/config"

	"github.com/gin-gonic/gin"
)

func LogoutHandler(c *gin.Context) {
	token, err := c.Cookie(config.CookieName)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)

		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     config.CookieName,
		Value:    token,
		Expires:  time.Now().AddDate(0, 0, -1),
		HttpOnly: true,
	})
}
