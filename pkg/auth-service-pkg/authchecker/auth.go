package authchecker

import (
	"2019_1_TheBang/config"
	pb "2019_1_TheBang/pkg/public/protobuf"
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	Id       float64 `json:"id"`
	Nickname string  `json:"nickname"`
	PhotoURL string  `json:"photo_url"`

	jwt.StandardClaims
}

func InfoFromCookie(tokenStr string) (userInfo pb.UserInfo, err error) {
	var token *jwt.Token
	token, err = CheckTocken(tokenStr)
	if err != nil {
		config.Logger.Warnw("InfoFromCookie -> CheckTocken",
			"warn", err.Error())

		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userInfo.Id = claims["id"].(float64)
		userInfo.Nickname = claims["nickname"].(string)
		userInfo.PhotoUrl = claims["photo_url"].(string)
	} else {
		config.Logger.Warnw("InfoFromCookie",
			"warn", "Error with parsing token's claims")
	}

	return
}

func CheckTocken(tokenStr string) (token *jwt.Token, err error) {
	token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("CheckTocken: Unexpected signing method: %v", token.Header["alg"])
		}

		return config.SECRET, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		err = errors.New("CheckTocken: this is faked cookie")
		return nil, err
	}

	return token, nil
}
