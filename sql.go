package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct{ *sql.DB }

func (db *DB) init() {
	var err error
	db.DB, err = sql.Open("sqlite3", "sql.db")
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

func (db *DB) addUser(uid, name, rawpwd string) error {
	salt := newSalt()
	pwd := hashPassword(rawpwd, salt)

	sqlStmt := `INSERT INTO Users (uid, name, pwd, salt) VALUES (?, ?, ?, ?);`
	_, err := db.Exec(sqlStmt, uid, name, pwd, salt)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) queryUserName(uid string) string {
	sqlStmt := `SELECT name FROM Users WHERE uid = ?;`
	r := db.QueryRow(sqlStmt, uid)

	var name string
	if err := r.Scan(&name); err != nil {
		log.Println("query name", err)
	}
	return name
}

func (db *DB) queryUserPassword(uid, pwd string) bool {
	sqlStmt := `SELECT pwd, salt FROM Users WHERE uid = ?;`
	r := db.QueryRow(sqlStmt, uid)

	var hpwd, salt []byte
	if err := r.Scan(&hpwd, &salt); err != nil {
		log.Println("query password:", err)
		return false
	}

	return validPassword(pwd, hpwd, salt)
}
