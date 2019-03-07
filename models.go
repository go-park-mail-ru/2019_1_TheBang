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

// toDO заменить на бд
func CreateProfileStorage() ProfileStorage {
	prof := ProfileStorage{}
	prof.data = make(map[int]Profile)

	return prof
}

type Profile struct {
	Id int `json:"user_id, string"`
	Nickname string
	Name string
	Surname string
	DOB string //toDo нужно заменить на time.Time
	Photo string
}

type InfoText struct {
	Data string
}