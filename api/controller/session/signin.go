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

// SignIn create a new user
func SignIn(w http.ResponseWriter, r *http.Request) {

	var ctr session.SigninCtr

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

	c := flight.Context(w, r)
	ct.Final(w, session.SigninMember(&c, &ctr))
	return
}
