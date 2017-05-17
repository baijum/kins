package middleware

import (
	"github.com/baijum/kins/route"
	"github.com/baijum/kins/web"
	"github.com/urfave/negroni"
)

// MW is the Negroni middleware
var MW *negroni.Negroni

// Run starts the server
func Run(addr string) {
	MW.Run(addr)
}

func init() {
	MW = negroni.New(negroni.NewRecovery(), negroni.NewLogger(), negroni.NewStatic(web.AssetFS()))
	MW.UseHandler(route.URT)
}
