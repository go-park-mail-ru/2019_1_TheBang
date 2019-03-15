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

type Profile struct {
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	DOB      string `json:"dob"`
	Photo    string `json:"photo"`
	Score    int    `json:"score"`
}