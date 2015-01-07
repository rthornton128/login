// Package session creates a new session handler. This simple package
// uses an in-memory database (mapping) to manage sessions and generates a
// session cookie. It is not designed to be extensible or anything fancy.
// The UUIDs remain valid forever and are never removed while the server
// remains running but are not persistant between restarts
package session

import (
	"errors"
	"net/http"
	"time"

	"github.com/rthornton128/login/uuid"
)

// Session maps UUIDs to internal uids
type Session struct {
	m map[string]string
}

var errInvalidUUID = errors.New("Invalid or expired UUID")

// New creates a new in-memory session manager
func New() *Session {
	return &Session{m: make(map[string]string)}
}

// Add requests a new UUID/GID for the given user, adds it to the session
// manager and sends the newly generated ID to w in a cookie
func (s *Session) Add(w http.ResponseWriter, uid string) {
	gid := uuid.NewVersion4()
	s.m[gid] = uid

	cookie := &http.Cookie{
		Name:     "login_uuid_cookie",
		Value:    gid,
		Expires:  time.Now().Add(time.Minute * 5),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
}

// Query returns the user ID of the connected client based on the UUID stored
// in the session cookie
func (s *Session) Query(req *http.Request) (string, error) {
	c, err := req.Cookie("login_uuid_cookie")
	if err != nil {
		return "", err
	}
	if uid, ok := s.m[c.Value]; ok {
		return uid, nil
	}
	return "", errInvalidUUID
}
