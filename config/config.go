package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var (
	SECRET      = getSecret()
	CookieName  = "bang_token"
	ServerName  = "TheBang server"
	FrontentDst = getFrontDest()
	DefaultImg  = "default_img"
	// POS = "WORKPLACE" | "HEROKU"

	DBUSER     = getDBUser()
	DBPASSWORD = getDBPasswd()
	DBNAME     = getDBNAme()

	//connBDStr = " user=" + DBUSER + " dbname="+ DBNAME +" password=" + DBPASSWORD + " sslmode=disable"
	DB               *sql.DB = connectDB()
	RowsOnLeaderPage uint    = 6
	PORT                     = getPort()
)

func connectDB() *sql.DB {
	pos := os.Getenv("WORKPLACE")
	if pos == "HEROKU" {
		return connectDBHEROKU()
	}

	db, err := sql.Open("sqlite3", "local_bd.db")
	if err != nil {
		log.Fatalln(err.Error())
	}

	return db
}

func connectDBHEROKU() *sql.DB {
	DB, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Database connected!")

	return DB
}

func getDBPasswd() string {
	bdpass := os.Getenv("DBPASSWORD")
	if bdpass == "" {
		log.Println("There is no DBPASSWORD!")
		bdpass = "2017"
	}
	return bdpass
}

func getDBNAme() string {
	dbname := os.Getenv("DBNAME")
	if dbname == "" {
		log.Println("There is no DBNAME!")
		dbname = "tp"
	}
	return dbname
}

func getDBUser() string {
	dbuser := os.Getenv("DBUSER")
	if dbuser == "" {
		log.Println("There is no DBUSER!")
		dbuser = "postgres"
	}
	return dbuser
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
