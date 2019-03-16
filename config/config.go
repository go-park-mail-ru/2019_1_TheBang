package config

import (
	"database/sql"
	"log"
)

var (
	SECRET      []byte
	CookieName  = "bang_token"
	ServerName  = "TheBang server"
	FrontentDst = "localhost:3000"
	DefaultImg  = "default_img"
	connBDStr = "user=postgres dbname=tp password=2017 sslmode=disable"
	BD *sql.DB = connectDB(connBDStr)
)

func connectDB(connStr string) *sql.DB {
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected!")

	return DB
}
