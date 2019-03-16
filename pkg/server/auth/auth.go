package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-park-mail-ru/2019_1_TheBang/config"
	"log"
	"net/http"
)

func CheckTocken(r *http.Request) (token *jwt.Token, ok bool) {
	cookie, err := r.Cookie(config.CookieName)
	if err != nil {
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