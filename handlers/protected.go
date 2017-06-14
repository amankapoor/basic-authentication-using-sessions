package handlers

import (
	"net/http"

	"github.com/amankapoor/basic-authentication-using-sessions/common"
)

func Protected(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session-name")
	if session.IsNew {
		http.Redirect(w, r, "http://localhost:8080/", http.StatusPermanentRedirect)
		return
	}
	common.RenderTemplate(w, "templates/protected.html", nil)
}
