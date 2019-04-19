package mainconfig

import (
	"2019_1_TheBang/config"
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var (
	ServerName = "TheBang server"
	DefaultImg = "default_img"
	DBSCHEMA   = getDBschema()

	DB               *sql.DB = connectDB()
	RowsOnLeaderPage uint    = 6
	MAINPORT                 = getMainPort()
	POINTSPORT               = getPointsPort()
)

func connectDB() *sql.DB {
	pos := os.Getenv("WORKPLACE")
	if pos == "HEROKU" {
		bd := connectDBHEROKU()

		return bd
	}

	db, err := sql.Open("sqlite3", "local_bd.db")
	if err != nil {
		config.Logger.Fatal(err.Error())
	}
	config.Logger.Info("SQL: SQLlite3")
	preRunSQLliteDB(db)
	config.Logger.Info("Database connected!")

	return db
}

func getDBschema() string {
	schema := os.Getenv("DBSCHEMA")
	if schema != "" {
		schema += `.`
		return schema
	}
	config.Logger.Warn("There is no DBSCHEMA!")

	return ""
}

func connectDBHEROKU() *sql.DB {
	config.Logger.Info("SQL: Postgres sql")
	DB, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		config.Logger.Fatal(err.Error())
	}
	config.Logger.Info("Database connected!")

	return DB
}

func getDBPasswd() string {
	bdpass := os.Getenv("DBPASSWORD")
	if bdpass == "" {
		config.Logger.Warn("There is no DBPASSWORD!")
		bdpass = "2017"
	}

	return bdpass
}

func getDBNAme() string {
	dbname := os.Getenv("DBNAME")
	if dbname == "" {
		config.Logger.Warn("There is no DBNAME!")
		dbname = "tp"
	}

	return dbname
}

func getDBUser() string {
	dbuser := os.Getenv("DBUSER")
	if dbuser == "" {
		config.Logger.Warn("There is no DBUSER!")
		dbuser = "postgres"
	}

	return dbuser
}

func getMainPort() string {
	port := os.Getenv("MAINPORT")
	if port == "" {
		config.Logger.Warn("There is no MAINPORT!")
		port = "8001"
	}

	return port
}

func getPointsPort() string {
	port := os.Getenv("POINTSPORT")
	if port == "" {
		config.Logger.Warn("There is no POINTSPORT!")
		port = "50052"
	}

	return port
}

func preRunSQLliteDB(db *sql.DB) {
	_, err := db.Exec(sqlCreateTableSQLlite)
	if err != nil {
		config.Logger.Fatalf("preRunDB, table: %v", err.Error())
	}

}

var sqlCreateTableSQLlite = `create table IF NOT EXISTS ` + DBSCHEMA + `users (
	id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
	nickname citext unique not null,
	name citext null,
	surname citext null,
	dob date null,
	photo varchar(250) default 'default_img',
	score bigint default 0,
	passwd text not null
  )`

var sqlCreateSchema = `create schema project_bang`
