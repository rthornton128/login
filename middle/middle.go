// Package middle is an example of what middleware might look like
package middle

import (
	"log"
	"net/http"
)

// LogAccess logs the user ID of the person trying to login along with their
// IP address
func LogAccess(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			log.Println("Attempted login by:", r.FormValue("UserID"),
				"("+r.RemoteAddr+")")
		}
		f(w, r)
	}
}
