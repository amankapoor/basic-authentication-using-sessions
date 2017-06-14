package main

import (
	"net/http"

	"github.com/amankapoor/basic-authentication-using-sessions/handlers"
	"github.com/amankapoor/basic-authentication-using-sessions/middleware"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func main() {
	chain := alice.New(middleware.Middleware)

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Index)
	r.Handle("/login", chain.ThenFunc(handlers.Login))
	r.HandleFunc("/logout", handlers.Logout)

	r.HandleFunc("/protected", handlers.Protected)

	http.ListenAndServe(":8080", r)
}
