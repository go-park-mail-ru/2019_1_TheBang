package config

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
	"os"
)

var (
	SECRET       = getSecret()
	CookieName  = "bang_token"
	ServerName  = "TheBang server"
	FrontentDst  = getFrontDest()
	DefaultImg  = "default_img"
	connBDStr = "user=postgres dbname=tp password=2017 sslmode=disable"
	DB *sql.DB = connectDB(connBDStr)
	RowsOnLeaderPage uint = 6
	PORT = getPort()
)

func connectDB(connStr string) *sql.DB {
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
	}
	log.Println("Database connected!")

	return DB
}

func getFrontDest() string {
	dst := os.Getenv("FrontentDst")
	if dst == "" {
		log.Println("There is no FrontentDst!")
		dst = "http://localhost:3000"
	}
	return dst
}

func getSecret() []byte {
	secret := []byte(os.Getenv("SECRET"))
	if string(secret) == "" {
		log.Println("There is no SECRET!")
		secret = []byte("secret")
	}

	return secret
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("There is no PORT!")
		port = "8090"
	}
	return port
}


