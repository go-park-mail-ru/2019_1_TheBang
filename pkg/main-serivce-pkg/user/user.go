package user

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"2019_1_TheBang/api"
	"2019_1_TheBang/config"
	"2019_1_TheBang/config/mainconfig"

	"golang.org/x/crypto/bcrypt"
)

type UserInfo struct {
	Id       float64 `json:"id"`
	Nickname string  `json:"nickname"`
	PhotoURL string  `json:"photo_url"`
}

func hashPasswd(passwd string) string {
	pw := []byte(passwd)

	hash, err := bcrypt.GenerateFromPassword(pw, 0)
	if err != nil {
		log.Fatal(err.Error())
	}

	return string(hash)
}

func CreateUser(s *api.Signup) (status int) {
	s.Passwd = hashPasswd(s.Passwd)
	_, err := mainconfig.DB.Exec(sqlInsertUser,
		s.Nickname,
		s.Name,
		s.Surname,
		s.DOB,
		s.Passwd)
	if err != nil {
		config.Logger.Warnw("CreateUser",
			"warn", err.Error())

		return http.StatusConflict
	}

	return http.StatusCreated
}

func SelectUser(nickname string) (p api.Profile, status int) {
	rows, err := mainconfig.DB.Query(SQLSeletUser,
		nickname)
	if err != nil {
		config.Logger.Warnw("SelectUser",
			"warn", err.Error())

		return p, http.StatusBadRequest
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&p.Id,
			&p.Nickname,
			&p.Name,
			&p.Surname,
			&p.DOB,
			&p.Photo,
			&p.Score); err != nil {
			config.Logger.Warnw("SelectUser",
				"warn", err.Error())

			return p, http.StatusInternalServerError
		}
	}

	return p, http.StatusOK
}

func UpdateUser(nickname string, u api.Update) (p api.Profile, status int) {
	fmt.Println(u)

	res, err := mainconfig.DB.Exec(SQLUpdateUser,
		u.Name,
		u.Surname,
		u.DOB,
		nickname)
	if err != nil {
		config.Logger.Warnw("UpdateUser",
			"warn", err.Error())

		return p, http.StatusBadRequest
	}

	rows, _ := res.RowsAffected()
	if rows != 1 {
		config.Logger.Warnw("UpdateUser",
			"warn", err.Error())

		return p, http.StatusInternalServerError
	}

	p, status = SelectUser(nickname)
	if status != http.StatusOK {
		return p, http.StatusInternalServerError
	}

	return p, http.StatusOK
}

func CheckUser(nickname, passwd string) bool {
	row, err := mainconfig.DB.Query(SQLCheckUser,
		nickname)
	if err != nil {
		return false
	}
	defer row.Close()

	if !row.Next() {
		return false
	}

	var hash string

	if err := row.Scan(
		&hash); err != nil {
		config.Logger.Warnw("CheckUser",
			"warn", err.Error())

		return false
	}

	if err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwd)); err != nil {
		config.Logger.Warnw("CheckUser",
			"warn", err.Error())

		return false
	}

	return true
}

func DeletePhoto(filename string) {
	if filename == mainconfig.DefaultImg {
		return
	}

	err := os.Remove("tmp/" + filename)
	if err != nil {
		config.Logger.Warnw("GetIconHandler",
			"warn", err.Error())

		return
	}
}

func UpdateUserPhoto(nickname, photo string) bool {
	res, err := mainconfig.DB.Exec(SQLUpdatePhoto,
		photo, nickname)
	if err != nil {
		config.Logger.Warnw("UpdateUserPhoto",
			"warn", err.Error())

		return false
	}

	rows, _ := res.RowsAffected()
	if rows != 1 {
		config.Logger.Warnw("UpdateUserPhoto",
			"warn", "wrong count affected rows")

		return false
	}

	return true
}

func DeleteUser(nickname string) bool {
	res, err := mainconfig.DB.Exec(SQLDeleteUser, nickname)
	if err != nil {
		config.Logger.Warnw("DeleteUser",
			"warn", err.Error())

		return false
	}

	rows, _ := res.RowsAffected()
	if rows != 1 {
		config.Logger.Warnw("DeleteUser",
			"warn", "wrong count affected rows")

		return false
	}

	return true
}

func UpdateUserScore(id float64, points int32) bool {
	res, err := mainconfig.DB.Exec(SQLUpdateUserScore, points, id)
	if err != nil {
		config.Logger.Warnw("UpdateUserScore",
			"warn", err.Error())

		return false
	}

	rows, _ := res.RowsAffected()
	if rows != 1 {
		config.Logger.Warnw("UpdateUserScore",
			"warn", "wrong count affected rows")

		return false
	}

	return true
}

var sqlInsertUser = `insert into ` + mainconfig.DBSCHEMA + `users (nickname, name, surname, dob, passwd) values ($1, $2, $3, $4, $5)`

var SQLSeletUser = `select 
					id, nickname, name, surname, dob, photo, score	
					from ` + mainconfig.DBSCHEMA + `users
					where nickname = $1`

var SQLUpdateUser = `update ` + mainconfig.DBSCHEMA + `users 
						set (name, surname, dob) = ($1, $2, $3)
						where nickname = $4`

var SQLCheckUser = `select 
					passwd	
					from ` + mainconfig.DBSCHEMA + `users
					where nickname = $1`

var SQLUpdatePhoto = `update ` + mainconfig.DBSCHEMA + `users 
						set photo = $1
						where nickname = $2`

var SQLDeleteUser = `delete from ` + mainconfig.DBSCHEMA + `users
						where nickname = $1`

var SQLUpdateUserScore = `update ` + mainconfig.DBSCHEMA + `users 
						set score = score + $1
						where id = $2`
