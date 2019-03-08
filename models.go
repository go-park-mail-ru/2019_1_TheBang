package main

import (
	"sync"
)

// toDO заменить на бд
type AccountStorage struct {
	data map[string]string
	mu sync.Mutex
	count int // костыль для id
}

// toDO заменить на бд
func CreateAccountStorage() AccountStorage {
	acc := AccountStorage{}
	acc.data = make(map[string]string)

	return acc
}

// toDO заменить на бд
type ProfileStorage struct {
	data map[int]Profile
	mu sync.Mutex
	count int // костыль для id
}

//func (p *ProfileStorage)

// toDO заменить на бд
func CreateProfileStorage() ProfileStorage {
	prof := ProfileStorage{}
	prof.data = make(map[int]Profile)

	return prof
}

type Profile struct {
	Id int `json:"id"`
	Nickname string `json:"nickname"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	DOB string `json:"dob"`
	Photo string `json:"photo"`
}

type InfoText struct {
	Data string `json:"data"`
}