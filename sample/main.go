package main

import (
	"log"
	"net/http"
	"runtime"
	"yana/core/model/enum"
	yanasession "yana/core/session"
	"yana/sample/controller"

	"github.com/Dlacreme/httpd/back/router"
	"github.com/Dlacreme/httpd/back/server"
	"github.com/Dlacreme/httpd/back/session"
	"github.com/Dlacreme/httpd/boot"
	"github.com/Dlacreme/httpd/config/env"
	"github.com/Dlacreme/httpd/middlewares/logrequest"
	"github.com/Dlacreme/httpd/middlewares/rest"
	"github.com/gorilla/context"
)

// Called by default before main. Set basic stuff
func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)
	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// First, always load env.json file
	config, err := env.LoadConfig("env.json")
	if err != nil {
		log.Fatalln(err)
	}

	// Load the controller routes
	controller.LoadRoutes()

	// Load web components
	db := boot.LoadDefaultServices(config)

	// Load enum
	enum.Init(db)

	session.RegisterLoggedUserBuilder(yanasession.GetLoggedUser)

	// Define query filters
	routing := loadMiddlewares(router.Instance())

	server.Run(
		routing,
		nil,
		config.Server,
	)
}

func loadMiddlewares(h http.Handler) http.Handler {
	return router.ChainHandler( // Chain middleware, top middleware runs first
		h, // Handler to wrap
		//setUpCSRF,            // Prevent CSRF
		rest.Handler,         // Support changing HTTP method sent via query string
		logrequest.Handler,   // Log every request
		context.ClearHandler, // Prevent memory leak with gorilla.sessions
	)
}
