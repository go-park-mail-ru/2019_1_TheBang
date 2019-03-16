package api

type Login struct {
	Nickname string `json:"nickname"`
	Passwd   string `json:"passwd"`
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

type Update struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	DOB     string `json:"dob"`
}
