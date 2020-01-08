package main

import (
	"github.com/joecomscience/oauth2-server/pkg/middleware"
	"github.com/joecomscience/oauth2-server/pkg/oauth"
	"gopkg.in/oauth2.v3/server"
	"log"
	"net/http"
)

const (
	Port = ":3000"
)

var (
	srv *server.Server
)

func main() {
	srv = oauth.NewOAuthServer()
	o := oauth.Handler{Server: srv}

	http.HandleFunc("/token", o.Token)
	http.HandleFunc("/c	redentials", o.Credentials)
	http.HandleFunc("/protected", middleware.ValidateToken(o.Protected, srv))

	if err := http.ListenAndServe(Port, nil); err != nil {
		log.Fatalf("start server error: %s\n", err.Error())
	}
}

func GetServer() *server.Server {
	return srv
}
