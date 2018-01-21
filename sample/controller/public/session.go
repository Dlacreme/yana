package public

import (
	"net/http"
	"yana/core/session"

	"github.com/Dlacreme/httpd/back/flight"
)

func SignIn(w http.ResponseWriter, r *http.Request) {

	c := flight.Context(w, r)

	// Validate with required fields
	if !c.FormValid("email", "password") {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Form values
	email := r.FormValue("email")
	password := r.FormValue("password")

	ctr := session.LoginCtr{email, password}

	res := session.Login(&c, &ctr, &w, r)
	if res.Error != nil {
		Index(w, r)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusTemporaryRedirect)
}

func LogOut(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	session.Logout(&c, &w, r)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	return
}
