package session

import (
	"yana/core/middleware/acl"

	"github.com/Dlacreme/httpd/back/router"
)

// Load set routing
func Load() {

	router.Get("/me", Me)
	router.Post("/session/signin", SignIn, acl.DisallowAuth)
	router.Post("/session/login", Login, acl.DisallowAuth)
	router.Post("/session/logout", Logout)

}
