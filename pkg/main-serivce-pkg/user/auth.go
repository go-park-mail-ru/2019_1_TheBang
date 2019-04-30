package user

import (
	"2019_1_TheBang/config"
	"2019_1_TheBang/config/mainconfig"
	"2019_1_TheBang/pkg/public/auth"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	Id       float64 `json:"id"`
	Nickname string  `json:"nickname"`
	PhotoURL string  `json:"photo_url"`

	jwt.StandardClaims
}

func TokenFromCookie(r *http.Request) *jwt.Token {
	cookie, _ := r.Cookie(config.CookieName)
	tokenStr := cookie.Value
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return config.SECRET, nil
	})
	return token
}

func InfoFromCookie(token *jwt.Token) (userInfo UserInfo, status int) {
	userInfo = UserInfo{}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userInfo.Id = claims["id"].(float64)
		userInfo.Nickname = claims["nickname"].(string)
		userInfo.PhotoURL = claims["photo_url"].(string)
	} else {
		status = http.StatusInternalServerError
		config.Logger.Warnw("NicknameFromCookie",
			"warn", "Error with parsing token's claims")

		return userInfo, status
	}

	return userInfo, http.StatusOK
}

func GetFreashToken(username string) string {
	prof, _ := SelectUser(username)
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
	ss, _ := token.SignedString(config.SECRET)

	return ss
}
