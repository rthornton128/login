package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/rthornton128/login/crypt"
	"github.com/rthornton128/login/middle"
	"github.com/rthornton128/login/session"
	"github.com/rthornton128/login/store"
)

var templates *template.Template
var db store.SqliteDB
var sm *session.Session

func serveRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.ServeFile(w, r, "html"+r.URL.Path)
		return
	}
	uid, err := sm.Query(r)
	u := &store.User{UserID: uid}
	if err != nil {
		log.Println("session manager query:", err)
	}
	u.Query(db)

	if err := templates.ExecuteTemplate(w, "index.html", u); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		u := &store.User{
			UserID: r.FormValue("UserID"),
		}
		u.Query(db)
		if !crypt.Validate(r.FormValue("Password"), u.Password, u.Salt) {
			log.Print("user name and password do not match")
			http.Error(w, "user name and password do not match",
				http.StatusUnauthorized)
			return
		}
		sm.Add(w, u.UserID)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		password, salt := crypt.Hash(r.FormValue("Password"))
		u := &store.User{
			UserID:   r.FormValue("UserID"),
			Name:     r.FormValue("Name"),
			Password: password,
			Salt:     salt,
		}
		if err := u.Store(db); err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sm.Add(w, u.UserID)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func main() {
	/* compile template(s) and exit on failure */
	templates = template.Must(template.ParseFiles("html/index.html"))

	/* setup flags */
	addr := flag.String("addr", "", "Server address")
	port := flag.String("port", "8080", "Server port")
	flag.Parse()

	/* initialize session manager */
	sm = session.New()

	/* initialize database */
	db.Init("sqlite_login.db")

	/* setup handlers */
	http.HandleFunc("/login", middle.LogAccess(handleLogin))
	http.HandleFunc("/register", middle.LogAccess(handleRegister))
	http.HandleFunc("/", serveRoot)

	/* start server */
	log.Println("Starting server...")
	if err := http.ListenAndServe(*addr+":"+*port, nil); err != nil {
		log.Fatal(err)
	}
	log.Println("Server shutdown")
}
