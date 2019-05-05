package user

import (
	"net/http"
	"time"

	"2019_1_TheBang/api"
	"2019_1_TheBang/config"
	"2019_1_TheBang/config/mainconfig"
	"2019_1_TheBang/pkg/public/auth"

	"github.com/dgrijalva/jwt-go"
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

	ss, status := LoginAcount(signup.Nickname, signup.Passwd)
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
		Path:     "",
	}

	http.SetCookie(c.Writer, &cookie)

	prof, status := SelectUser(signup.Nickname)
	if status != http.StatusOK {
		c.AbortWithStatus(status)

		return
	}

	c.JSONP(http.StatusCreated, prof)
}

func LoginAcount(username, passwd string) (ss string, status int) {
	ok := CheckUser(username, passwd)
	if !ok {
		return ss, http.StatusUnauthorized
	}

	prof, _ := SelectUser(username)

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
