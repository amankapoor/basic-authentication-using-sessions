package handlers

import (
	"net/http"

	"github.com/gorilla/sessions"

	"log"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		if email == "aman.sk95@gmail.com" && password != "" {

			session, err := store.Get(r, "session-name")
			if err != nil {
				http.Redirect(w, r, "http://localhost:8080", http.StatusBadRequest)
				return
			}

			session.Values["foo"] = "bar"
			session.Values[42] = 43

			session.Save(r, w)

			log.Printf("&cookie gives: %v", &session)
			log.Printf("*cookie gives: %v", *session)
			log.Printf("cookie gives: %v", session)
		}

		http.Redirect(w, r, "http://localhost:8080/protected", http.StatusSeeOther)
	}
}
