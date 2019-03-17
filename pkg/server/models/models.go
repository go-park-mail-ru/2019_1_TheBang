package models

import (
	"github.com/dgrijalva/jwt-go"
)

type InfoText struct {
	Data string `json:"data"`
}

type CustomClaims struct {
	Nickname string `json:"nickname"`
	jwt.StandardClaims
}
