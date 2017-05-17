package main

import (
	"github.com/baijum/kins/auth"
	_ "github.com/baijum/kins/auth/google"
	"github.com/baijum/kins/config"
	"github.com/baijum/kins/middleware"
	"github.com/baijum/kins/route"
	_ "github.com/baijum/kins/web"
	"github.com/urfave/negroni"
)

func main() {
	route.URT.PathPrefix("/api").Handler(
		negroni.New(negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext), negroni.Wrap(route.RT)))
	middleware.Run(config.Config.HTTPAddress)
}
