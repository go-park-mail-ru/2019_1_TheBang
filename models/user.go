package models

type Signup struct {
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	DOB      string `json:"dob"`
	Passwd   string `json:"passwd"`
}

var InsertUser = `insert into project_bang.users (nickname, name, surname, dob, passwd)
    values ($1, $2, $3, $4, $5);`
