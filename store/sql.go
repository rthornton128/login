package store

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // for database/sql driver
)

// SqliteDB is a wrapper for the Sqlite3 database store
type SqliteDB struct{ *sql.DB }

// Init opens the database and sets up the tables if not already created
func (db *SqliteDB) Init(filename string) {
	var err error
	db.DB, err = sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatalln("open database:", err)
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS Users (uid TEXT PRIMARY KEY NOT NULL UNIQUE, 
	name TEXT, pwd TEXT, salt TEXT);`

	if _, err := db.Exec(sqlStmt); err != nil {
		log.Fatalln("creating table", err)
	}
}

// AddUser inserts user data into the database
func (db SqliteDB) AddUser(u *User) (err error) {
	sqlStmt := `INSERT INTO Users (uid, name, pwd, salt) VALUES (?, ?, ?, ?);`
	_, err = db.Exec(sqlStmt, u.UserID, u.Name, u.Password, u.Salt)
	return err
}

// QueryUser queries the database for the user identified by the UserID
// field
func (db SqliteDB) QueryUser(u *User) (err error) {
	sqlStmt := `SELECT name, pwd, salt FROM Users WHERE uid = ?;`
	r := db.QueryRow(sqlStmt, u.UserID)

	return r.Scan(&u.Name, &u.Password, &u.Salt)
}
