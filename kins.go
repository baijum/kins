package main

import (
	"github.com/urfave/negroni"

	"github.com/baijum/kins/auth"
	"github.com/baijum/kins/config"
	"github.com/baijum/kins/middleware"
	"github.com/baijum/kins/route"

	_ "github.com/baijum/kins/auth/google"
	_ "github.com/baijum/kins/web"
)

func main() {
	route.URT.PathPrefix("/api").Handler(
		negroni.New(negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext), negroni.Wrap(route.RT)))
	middleware.Run(config.Config.HTTPAddress)
}
