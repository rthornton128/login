package main

import (
	"os"
	"testing"
)

var db DB

func init() {
	os.Remove("sql.db")
	db.init()
}

func TestAddUserAndQueryPassword(t *testing.T) {
	name := "test"
	pass := "myUberpwd123"

	db.addUser(name, name, pass)
	if ok := db.queryUserPassword(name, pass); !ok {
		t.FailNow()
	}
}

func TestAddUserTwice(t *testing.T) {
	name := "test"
	pass := "password"

	db.addUser(name, name, pass)
	if err := db.addUser(name, name, pass); err == nil {
		t.Fatal("Expected error from addUser but got nil")
	}
}

func TestQueryUserName(t *testing.T) {
	uid := "myuid"
	name := "myname"
	pass := "password"

	db.addUser(uid, name, pass)
	if res := db.queryUserName(uid); res != name {
		t.Log("Expected", name, "Got:", res)
		t.FailNow()
	}
}
