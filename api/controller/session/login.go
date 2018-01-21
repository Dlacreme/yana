package session

import (
	"encoding/json"
	"net/http"

	ct "yana/api/common/controller"
	"yana/api/controller/status"
	"yana/core/session"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/back/wesult"
)

// Me will check if user is logged and return basic information
func Me(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	ct.Final(w, session.Me(&c))
	return
}

// Login allow user to log in
func Login(w http.ResponseWriter, r *http.Request) {

	c := flight.Context(w, r)

	var ctr session.LoginCtr
	err := json.NewDecoder(r.Body).Decode(&ctr)

	if err != nil {
		status.Error500(w, r)
		return
	}

	inputErrors := ctr.Check()
	if inputErrors != nil {
		ct.Final(w, wesult.New(nil, inputErrors))
		return
	}

	res := session.Login(&c, &ctr, &w, r)

	ct.Final(w, res)
}

// Logout logout the user
func Logout(w http.ResponseWriter, r *http.Request) {

	c := flight.Context(w, r)

	session.Logout(&c, &w, r)

	ct.Final(w, wesult.New(nil, nil))
	return
}
