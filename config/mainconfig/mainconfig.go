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
	DBSCHEMA   = ""

	DB               *sql.DB = connectDB()
	RowsOnLeaderPage uint    = 6
	MAINPORT                 = getMainPort()
	POINTSPORT               = getPointsPort()
)

func connectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "local_bd.db")
	if err != nil {
		config.Logger.Fatal(err.Error())
	}
	config.Logger.Info("SQL: SQLlite3")
	preRunSQLliteDB(db)
	config.Logger.Info("Database connected!")

	return db
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

var sqlCreateTableSQLlite = `create table IF NOT EXISTS ` + `users (
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
