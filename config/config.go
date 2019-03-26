package config

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var (
	Logger      = createGlobalLogger()
	SECRET      = getSecret()
	CookieName  = "bang_token"
	ServerName  = "TheBang server"
	FrontentDst = getFrontDest()
	DefaultImg  = "default_img"
	// POS = "WORKPLACE" | "HEROKU"

	// DBUSER     = getDBUser()
	// DBPASSWORD = getDBPasswd()
	// DBNAME     = getDBNAme()
	DBSCHEMA = getDBschema()

	//connBDStr = " user=" + DBUSER + " dbname="+ DBNAME +" password=" + DBPASSWORD + " sslmode=disable"
	DB               *sql.DB = connectDB()
	RowsOnLeaderPage uint    = 6
	PORT                     = getPort()
)

func connectDB() *sql.DB {
	pos := os.Getenv("WORKPLACE")
	if pos == "HEROKU" {
		bd := connectDBHEROKU()
		return bd
	}

	db, err := sql.Open("sqlite3", "local_bd.db")
	if err != nil {
		Logger.Fatal(err.Error())
	}
	Logger.Info("SQL: SQLlite3")
	preRunSQLliteDB(db)
	Logger.Info("Database connected!")

	return db
}

func getDBschema() string {
	schema := os.Getenv("DBSCHEMA")
	if schema != "" {
		schema += `.`
		return schema
	}
	Logger.Warn("There is no DBSCHEMA!")
	return ""
}

func connectDBHEROKU() *sql.DB {
	Logger.Info("SQL: Postgres sql")
	DB, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		Logger.Fatal(err.Error())
	}
	Logger.Info("Database connected!")

	return DB
}

func getDBPasswd() string {
	bdpass := os.Getenv("DBPASSWORD")
	if bdpass == "" {
		Logger.Warn("There is no DBPASSWORD!")
		bdpass = "2017"
	}
	return bdpass
}

func getDBNAme() string {
	dbname := os.Getenv("DBNAME")
	if dbname == "" {
		Logger.Warn("There is no DBNAME!")
		dbname = "tp"
	}
	return dbname
}

func getDBUser() string {
	dbuser := os.Getenv("DBUSER")
	if dbuser == "" {
		Logger.Warn("There is no DBUSER!")
		dbuser = "postgres"
	}
	return dbuser
}

func getFrontDest() string {
	dst := os.Getenv("FrontentDst")
	if dst == "" {
		Logger.Warn("There is no FrontentDst!")
		dst = "http://localhost:3000"
	}
	return dst
}

func getSecret() []byte {
	secret := []byte(os.Getenv("SECRET"))
	if string(secret) == "" {
		Logger.Warn("There is no SECRET!")
		secret = []byte("secret")
	}

	return secret
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		Logger.Warn("There is no PORT!")
		port = "8090"
	}
	return port
}

func preRunSQLliteDB(db *sql.DB) {
	_, err := db.Exec(sqlCreateTableSQLlite)
	if err != nil {
		Logger.Fatalf("preRunDB, table: %v", err.Error())
	}

}

var sqlCreateTableSQLlite = `create table IF NOT EXISTS ` + DBSCHEMA + `users (
	id bigserial primary key,
	nickname citext unique not null,
	name citext null,
	surname citext null,
	dob date null,
	photo varchar(250) default 'default_img',
	score bigint default 0,
	passwd text not null
  )`

var sqlCreateSchema = `create schema project_bang`
