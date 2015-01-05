/* Package login/store acts as the persistant storage for the login example
 * program. It uses sqlite but any other SQL based storage would work with
 * minimal work
 * This package needs to accomplish three tasks:
 * 1) Store a User's ID, Name, Password and Salt
 * 2) Retrieve the Password and Salt hashes for verification
 * 3) Retrieve the Name of the User */
package store

type User struct {
	UserID   string
	Name     string
	Password []byte
	Salt     []byte
}

type Storage interface {
	AddUser(*User) error
	QueryUser(*User) error
}

/* Store user data */
func (u *User) Store(s Storage) error {
	return s.AddUser(u)
}

/* Query searches the storage for the user ID stored in u.UserID and
 * sets the other fields to the data discovered. If the user ID is invalid
 * or doesn't exist then the other fields are zero'd. */
func (u *User) Query(s Storage) error {
	u.Name, u.Password, u.Salt = "", []byte{}, []byte{}
	return s.QueryUser(u)
}
