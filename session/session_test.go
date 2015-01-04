package session_test

import (
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/rthornton128/login/session"
)

/* only determines if a cookie was set, not whether it is correct in any way */
func TestSessionAdd(t *testing.T) {
	var s session.Session
	s.Init()

	w := httptest.NewRecorder()
	s.Add(w, "user_id")

	cookie_str := w.HeaderMap.Get("Set-Cookie")
	if cookie_str == "" {
		t.Fail()
	}
}

func TestSessionQuery(t *testing.T) {
	/* setup basic client */
	cj, _ := cookiejar.New(nil)
	client := &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           cj,
		Timeout:       time.Second * 10,
	}

	/* setup session manager */
	var s session.Session
	s.Init()

	/* setup test server */
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			s.Add(w, "user_id")
		}))
	defer ts.Close()

	/* test for uuid query */
	r, err := client.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	cookies := r.Cookies()
	if len(cookies) != 1 {
		t.Log("Expected 1 cookie, got:", len(cookies))
		t.Fail()
	}
	r.Request.AddCookie(cookies[0])
	uid, err := s.Query(r.Request)
	if err != nil {
		t.Fatal(err)
	}
	if uid != "user_id" {
		t.Log("Expected 'user_id', got:", uid)
	}
}