package main

import (
	"sync"
)

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
	data  map[int]Profile
	mu    sync.Mutex
	count int // костыль для id
}

// toDO заменить на бд
func CreateProfileStorage() ProfileStorage {
	prof := ProfileStorage{}
	prof.data = make(map[int]Profile)

	//toDO убрать чудо админа
	admin := Profile{
		Id:       0,
		Nickname: "admin",
		Name:     "admin",
		Surname:  "admin",
		DOB:      "0.0.0.0",
		Photo:    defaultImg,
	}
	prof.data[0] = admin
	prof.count++

	return prof
}

type Profile struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	DOB      string `json:"dob"`
	Photo    string `json:"photo"`
}

type Login struct {
	Nickname string `json:"nickname"`
	Passwd   string `json:"passwd"`
}

type Signup struct {
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	DOB      string `json:"dob"`
	Passwd   string `json:"passwd"`
}

type Update struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	DOB     string `json:"dob"`
}

type InfoText struct {
	Data string `json:"data"`
}
