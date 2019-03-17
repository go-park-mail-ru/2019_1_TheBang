package models

import (
	"github.com/go-park-mail-ru/2019_1_TheBang/api"
	"github.com/go-park-mail-ru/2019_1_TheBang/config"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func hashPasswd(passwd string) string {
	pw := []byte(passwd)

	hash, err := bcrypt.GenerateFromPassword(pw, 0)
	if err != nil {
		log.Fatal(err.Error())
	}

	return string(hash)
}

func CreateUser(s *api.Signup) (profile api.Profile, status int) {
	s.Passwd = hashPasswd(s.Passwd)

	_, err := config.DB.Query(SQLInsertUser,
		s.Nickname,
		s.Name,
		s.Surname,
		s.DOB,
		s.Passwd)
	if err != nil {
		log.Printf("CreateUser: %v", err.Error())
		return profile, http.StatusConflict
	}

	profile = api.Profile{
		Nickname: s.Nickname,
		Name:     s.Name,
		Surname:  s.Surname,
		DOB:      s.DOB,
	}
	profile.Photo = config.DefaultImg

	return profile, http.StatusCreated
}

func SelectUser(nickname string) (p api.Profile, status int) {
	rows, err := config.DB.Query(SQLSeletUser,
		nickname)
	if err != nil {
		log.Printf("SelectUser: %v", err.Error())
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

func UpdateUser(nickname string, u api.Update) (p api.Profile, status int) {
	_, err := config.DB.Query(SQLUpdateUser,
		u.Name,
		u.Surname,
		u.DOB,
		nickname)
	if err != nil {
		log.Printf("UpdateUser: %v", err.Error())

		return p, http.StatusBadRequest
	}

	p, status = SelectUser(nickname)
	if status != http.StatusOK {
		return p, http.StatusInternalServerError
	}

	return p, http.StatusOK
}

func CheckUser(nickname, passwd string) bool {
	row, err := config.DB.Query(SQLCheckUser,
		nickname)
	if err != nil {
		return false
	}

	if !row.Next() {
		return false
	}

	var hash string

	if err := row.Scan(
		&hash);
	err != nil {
		log.Printf("ProfileHandler: %v\n", err.Error())

		return false
	}

	if err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwd)); err != nil {
		log.Printf("ProfileHandler: wrong passwd: %v\n", err.Error())


		return false
	}

	return true
}

func UpdateUserPhoto(nickname, photo string) bool {
	_, err := config.DB.Query(SQLUpdatePhoto,
		photo, nickname)
	if err != nil {
		log.Printf("UpdateUserPhoto: %v\n")

		return false
	}

	return true
}

func DeleteUser(nickname string) bool {
	_, err := config.DB.Query(SQLDeleteUser, nickname)
	if err != nil {
		log.Printf("DeleteUser: %v\n")

		return false
	}

	return true
}

var SQLInsertUser = `insert into project_bang.users
 						(nickname, name, surname, dob, passwd)
    					values ($1, $2, $3, $4, $5)`

var SQLSeletUser = `select 
					nickname, name, surname, dob, photo, score	
					from project_bang.users
					where nickname = $1`

var SQLUpdateUser = `update project_bang.users 
						set (name, surname, dob) = ($1, $2, $3)
						where nickname = $4`

var SQLCheckUser = `select 
					passwd	
					from project_bang.users
					where nickname = $1`

var SQLUpdatePhoto = `update project_bang.users 
						set photo = $1
						where nickname = $2`

var SQLDeleteUser = `delete from project_bang.users
						where nickname = $1`

