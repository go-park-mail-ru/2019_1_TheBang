package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-park-mail-ru/2019_1_TheBang/config"
	"sync"
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

type Signup struct {
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	DOB      string `json:"dob"`
	Passwd   string `json:"passwd"`
}

// toDO заменить на бд
type AccountStorage struct {
	data  map[string]string
	mu    sync.Mutex
	count int // костыль для id
}

// toDO заменить на бд
func CreateAccountStorage() AccountStorage {
	acc := AccountStorage{}
	acc.data = make(map[string]string)

	//toDo убрать эту чудо запись
	acc.data["admin"] = "admin"

	return acc
}

// toDO заменить на бд
type ProfileStorage struct {
	data  map[string]Profile
	mu    sync.Mutex
	count int // костыль для id
}

// toDO заменить на бд
func CreateProfileStorage() ProfileStorage {
	prof := ProfileStorage{}
	prof.data = make(map[string]Profile)

	//toDO убрать чудо админа
	admin := Profile{
		Nickname: "admin",
		Name:     "admin",
		Surname:  "admin",
		DOB:      "0.0.0.0",
		Photo:    config.DefaultImg,
		Score:    1000,
	}
	prof.data[admin.Nickname] = admin

	return prof
}