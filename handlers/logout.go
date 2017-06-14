package handlers

import (
	"net/http"

	"log"

	"github.com/amankapoor/new/common"
)

func Logout(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session-name")
	if session.IsNew == false {
		session.Options.MaxAge = -1
		session.Save(r, w)
		common.RenderTemplate(w, "templates/logout.html", nil)
		return
	}

	w.Write([]byte("You are not logged in."))
	log.Printf("Not logged in with session: %v", session)
}
