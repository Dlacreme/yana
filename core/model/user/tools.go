package user

import (
	"fmt"

	"github.com/Dlacreme/httpd/back/wdb"
	"github.com/Dlacreme/httpd/back/werror"
)

func isExisting(db wdb.Connection, email string) (bool, *werror.Error) {

	var res int

	err := db.Get(&res, fmt.Sprintf(`
						SELECT COUNT(UserId)
						FROM %v 
						WHERE Email = ?
					`, table),
		email)

	if err != nil {
		return true, werror.New(500, fmt.Sprintf("Cannot check is user exists : %s", err.Error()))
	}

	if res > 0 {
		return true, nil
	}

	return false, nil
}

func IsStaff(db wdb.Connection, userId int) bool {
	var res bool
	q := fmt.Sprintf(`SELECT IF(UserId = 3, 0, 1) FROM User WHERE UserId = %d`, userId)
	err := db.Get(&res, q)
	if err != nil {
		return false
	}
	return res
}
