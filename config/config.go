package config

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

var (
	SECRET      []byte
	CookieName  = "bang_token"
	ServerName  = "TheBang server"
	FrontentDst = "localhost:3000"
	DefaultImg  = "default_img"
	connBDStr = "user=postgres dbname=tp password=2017 sslmode=disable"
	DB *sql.DB = connectDB(connBDStr)
	RowsOnLeaderPage uint = 6
)

func connectDB(connStr string) *sql.DB {
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected!")

	return DB
}
