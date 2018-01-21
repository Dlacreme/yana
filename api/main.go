package main

import (
	"log"
	"net/http"
	"runtime"

	"yana/api/controller"
	"yana/core/model/enum"
	yanasession "yana/core/session"

	"github.com/Dlacreme/httpd/back/router"
	"github.com/Dlacreme/httpd/back/server"
	"github.com/Dlacreme/httpd/back/session"
	"github.com/Dlacreme/httpd/boot"
	"github.com/Dlacreme/httpd/config/env"
	"github.com/Dlacreme/httpd/middlewares/logrequest"
	"github.com/Dlacreme/httpd/middlewares/rest"
	"github.com/gorilla/context"
)

// init sets runtime settings.
func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	config, err := env.LoadConfig("env.json")
	if err != nil {
		log.Fatalln(err)
	}

	controller.LoadRoutes()

	// Register the services
	db := boot.LoadDefaultServices(config)

	enum.Init(db)

	session.RegisterLoggedUserBuilder(yanasession.GetLoggedUser)

	handler := loadMiddlewares(router.Instance())

	// Start the HTTP and HTTPS listeners
	server.Run(
		handler,       // HTTP handler
		handler,       // HTTPS handler
		config.Server, // Server settings
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
