package login

import (
	"fmt"
	"net/http"
	"time"

	"2019_1_TheBang/api"
	"2019_1_TheBang/config"
	"2019_1_TheBang/config/mainconfig"
	"2019_1_TheBang/pkg/main-serivce-pkg/user"
	"2019_1_TheBang/pkg/public/auth"
	_ "encoding/json"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func LogInHandler(c *gin.Context) {
	login := api.Login{}
	err := c.BindJSON(&login)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	ss, status := LoginAcount(login.Nickname, login.Passwd)
	if status != http.StatusOK {
		c.AbortWithStatus(status)

		return
	}

	expiration := time.Now().Add(10 * time.Hour)
	cookie := http.Cookie{
		Name:     config.CookieName,
		Value:    ss,
		Expires:  expiration,
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, &cookie)

	prof, status := user.SelectUser(login.Nickname)
	if status != http.StatusOK {
		c.AbortWithStatus(status)

		return
	}

	c.JSONP(http.StatusOK, prof)
}

func LoginAcount(username, passwd string) (ss string, status int) {
	ok := user.CheckUser(username, passwd)
	if !ok {
		return ss, http.StatusUnauthorized
	}

	prof, _ := user.SelectUser(username)

	// логирование
	fmt.Println(prof)

	claims := auth.CustomClaims{
		prof.Id,
		prof.Nickname,
		prof.Photo,
		jwt.StandardClaims{
			Issuer: mainconfig.ServerName,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(config.SECRET)
	if err != nil {
		config.Logger.Warnw("LoginAcount",
			"Error with JWT tocken generation:", err.Error())

		return ss, http.StatusInternalServerError
	}

	return ss, http.StatusOK
}
