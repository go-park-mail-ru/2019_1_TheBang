package middleware

import (
	"2019_1_TheBang/config"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func CheckTocken(r *http.Request) (token *jwt.Token, ok bool) {
	cookie, err := r.Cookie(config.CookieName)
	if err != nil {
		config.Logger.Warnw("CheckTocken",
			"warn", err.Error())
		return nil, false
	}

	tokenStr := cookie.Value

	token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return config.SECRET, nil
	})
	if err != nil {
		log.Printf("Error with check tocken: %v", err.Error())

		return nil, false
	}

	if !token.Valid {
		log.Printf("%v use faked cookie: %v\n", r.RemoteAddr, err.Error())

		return nil, false
	}

	return token, true
}
