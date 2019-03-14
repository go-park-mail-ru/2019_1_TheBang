package config

import "github.com/go-park-mail-ru/2019_1_TheBang/models"

var (
	StorageAcc models.AccountStorage
	StorageProf models.ProfileStorage
	SECRET      []byte
	CookieName  = "bang_token"
	ServerName  = "TheBang server"
	FrontentDst = "localhost:3000"
	DefaultImg  = "default_img"
)
