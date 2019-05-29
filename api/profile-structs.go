package api

//easyjson:json
type ProfileList []Profile

type Profile struct {
	Id       uint   `json:"id"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	DOB      string `json:"dob"`
	Photo    string `json:"photo"`
	Score    int    `json:"score"`
}
