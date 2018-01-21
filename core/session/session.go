package session

import (
	"net/http"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/back/werror"
	"github.com/Dlacreme/httpd/back/wesult"
	sess "github.com/blue-jay/core/session"

	"yana/core/model/user"
)

//// ME

// Me will provide basic info
func Me(c *flight.Info) wesult.Result {
	if c.Sess.Values["UserId"] != nil {
		usr, _ := user.FindById(c.DB, c.Sess.Values["UserId"].(int))
		return wesult.New(usr, nil)
	}
	return wesult.New(nil, werror.New(http.StatusUnauthorized, "no auth"))
}

// SigninAdmin will create a new session
func SigninAdmin(c *flight.Info, input *SigninCtr) wesult.Result {
	id, err := user.CreateAdmin(c.DB, input.Email, "Default name", input.Password)

	if err != nil {
		return wesult.New(nil, err)
	}

	user, err := user.FindById(c.DB, id)

	if err != nil {
		return wesult.New(nil, err)
	}

	return wesult.New(user, nil)
}

// SigninStaff will create a new session
func SigninStaff(c *flight.Info, input *SigninCtr) wesult.Result {
	id, err := user.CreateStaff(c.DB, input.Email, "Default name", input.Password)

	if err != nil {
		return wesult.New(nil, err)
	}

	user, err := user.FindById(c.DB, id)

	if err != nil {
		return wesult.New(nil, err)
	}

	return wesult.New(user, nil)
}

// SigninMember will create a new session
func SigninMember(c *flight.Info, input *SigninCtr) wesult.Result {
	id, err := user.CreateMember(c.DB, input.Email, "Default name", input.Password)

	if err != nil {
		return wesult.New(nil, err)
	}

	user, err := user.FindById(c.DB, id)

	if err != nil {
		return wesult.New(nil, err)
	}

	return wesult.New(user, nil)
}

// Login will process the user login
func Login(c *flight.Info, input *LoginCtr, w *http.ResponseWriter, r *http.Request) wesult.Result {
	usr, err := user.FindByCredentials(c.DB, input.Email, input.Password)

	if usr == nil {
		return wesult.New(nil, werror.New(http.StatusUnauthorized, "Invalid credentials"))
	}
	if err != nil {
		return wesult.New(nil, err)
	}
	sess.Empty(c.Sess)
	c.Sess.Values["UserId"] = usr.UserId
	c.Sess.Values["UserTypeId"] = usr.TypeId
	c.Sess.Values["Email"] = usr.Email
	c.Sess.Values["Name"] = usr.Name
	c.Sess.Save(r, *w)

	return wesult.New(usr, nil)
}

// Logout a user
func Logout(c *flight.Info, w *http.ResponseWriter, r *http.Request) {
	sess.Empty(c.Sess)
	c.Sess.Save(r, *w)
}
