package store_test

import (
	"os"
	"testing"

	"github.com/rthornton128/login/crypt"
	"github.com/rthornton128/login/store"
)

var dataStore store.SqliteDB

func TestMain(m *testing.M) {
	dataStore.Init("test.db")
	result := m.Run()
	os.Remove("test.db")

	os.Exit(result)
}

func TestStore(t *testing.T) {
	password, salt := crypt.Hash("password")
	u := &store.User{
		UserID:   "uid",
		Name:     "Some Name",
		Password: password,
		Salt:     salt,
	}
	if err := u.Store(dataStore); err != nil {
		t.Fatal(err)
	}
}

func TestQueryAndValidate(t *testing.T) {
	password, salt := crypt.Hash("password")
	u := &store.User{
		UserID:   "uid",
		Name:     "Some Name",
		Password: password,
		Salt:     salt,
	}
	u.Store(dataStore) /* ignore errors and muscle through */
	u.UserID = "invalid"

	if u.Query(dataStore) == nil {
		t.Fatal("Expected error but not received")
	}

	u.UserID = "uid"
	err := u.Query(dataStore)
	if err != nil {
		t.Fatal(err)
	}

	if u.Name != "Some Name" {
		t.Fatal("Wrong user name:", u.Name)
	}
	if !crypt.Validate("password", u.Password, u.Salt) {
		t.Fatal("Password did not validate")
	}
}
