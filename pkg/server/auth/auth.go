package auth

import (
	"fmt"
	"log"
	"net/http"

	"2019_1_TheBang/config"

	"github.com/dgrijalva/jwt-go"
)

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

func NicknameFromCookie(token *jwt.Token) (nickname string, status int) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		nickname = claims["nickname"].(string)
	} else {
		status = http.StatusInternalServerError
		config.Logger.Warnw("NicknameFromCookie",
			"warn", "Error with parsing token's claims")

		return nickname, status
	}

	return nickname, http.StatusOK
}

func CheckTocken(r *http.Request) (token *jwt.Token, ok bool) {
	cookie, err := r.Cookie(config.CookieName)
	if err != nil {
		log.Printf("CheckTocken: %v", err.Error())
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
