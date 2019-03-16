package models

import (
	"github.com/go-park-mail-ru/2019_1_TheBang/config"
	"log"
	"net/http"
)

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

func CreateUser(s *Signup) (profile Profile, status int) {
	_, err := config.DB.Query(SQLInsertUser,
		s.Nickname,
		s.Name,
		s.Surname,
		s.DOB,
		s.Passwd)
	if err != nil {
		return profile, http.StatusConflict
	}

	profile = Profile{
		Nickname: s.Nickname,
		Name:     s.Name,
		Surname:  s.Surname,
		DOB:      s.DOB,
	}

	return profile, http.StatusCreated
}

func SelectUser(nickname string) (p Profile, status int) {
	rows, err := config.DB.Query(SQLSeletUser,
		nickname)
	if err != nil {
		return p, http.StatusBadRequest
	}

	for rows.Next() {
		if err := rows.Scan(&p.Nickname,
			&p.Name,
			&p.Surname,
			&p.DOB,
			&p.Photo,
			&p.Score);
		err != nil {
			log.Printf("ProfileHandler: %v\n", err.Error())

			return p, http.StatusInternalServerError
		}
	}

	return p, http.StatusOK
}

var SQLInsertUser = `insert into project_bang.users
 						(nickname, name, surname, dob, passwd)
    					values ($1, $2, $3, $4, $5)`

var SQLSeletUser = `select 
					nickname, name, surname, dob, photo, score	
					from project_bang.users
					where nickname = $1`
