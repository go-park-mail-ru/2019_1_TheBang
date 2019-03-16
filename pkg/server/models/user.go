package models

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

func CreateUser(s *Signup) (Profile, int) {

	prof := Profile{
		Nickname: s.Nickname,
		Name:     s.Name,
		Surname:  s.Surname,
		DOB:      s.DOB,
	}
}

func SelectUser() {}

var SQLInsertUser = `insert into project_bang.users (nickname, name, surname, dob, passwd)
    values ($1, $2, $3, $4, $5);`

var SQLSeletUser = `select 
					nickname, name, surname, dob, photo, score	
					from project_bang.users
					where nickname = $1`
