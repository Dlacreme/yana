package session

import "github.com/Dlacreme/httpd/back/werror"

// SigninCtr :
type SigninCtr struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	Confirmation string `json:"confirmation"`
}

// Check make sure the input value are correct
func (ct *SigninCtr) Check() *werror.Error {
	return nil
}

// LoginCtr provide input for login
type LoginCtr struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Check make sure the input value are correct
func (ct *LoginCtr) Check() *werror.Error {
	return nil
}
