package chatconfig

import (
	"2019_1_TheBang/config"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var (
	CHATPORT = getPort()

	DB       *sql.DB = connectDB()
	DBSCHEMA         = getDBschema()
)

func getPort() string {
	port := os.Getenv("CHATPORT")
	if port == "" {
		return "8003"
	}

	return port
}

func getDBschema() string {
	return ""
}

func connectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "local_chat_db.db")
	if err != nil {
		config.Logger.Fatal(err.Error())
	}

	config.Logger.Info("SQL: SQLlite3")
	preRunSQLliteDB(db)
	config.Logger.Info("Database connected!")

	return db
}

func preRunSQLliteDB(db *sql.DB) {
	_, err := db.Exec(sqlCreateTableSQLlite)
	if err != nil {
		config.Logger.Fatalf("preRunDB, table: %v", err.Error())
	}

}

var sqlCreateTableSQLlite = `create table IF NOT EXISTS ` + DBSCHEMA + `messages (
	id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
	author varchar(250) NOT NULL,
	timestamp INTEGER NOT NULL,
	message text NOT NULL,
	photo_url varchar(250) NOT NULL
  )`
