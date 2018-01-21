package unit

import (
	"database/sql"
	"fmt"

	"github.com/Dlacreme/httpd/back/wdb"
	"github.com/Dlacreme/httpd/back/werror"
)

// Unit db model
type Unit struct {
	UnitId int    `db:"UnitId"`
	Label  string `db:"Label"`
}

// GetUnits get
func GetUnits(db wdb.Connection) (*[]Unit, *werror.Error) {
	res := []Unit{}

	q := `SELECT u.UnitId, u.Label FROM Unit`

	err := db.Select(&res, q)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, werror.New(500, fmt.Sprintf("Cannot get Units : %s", err.Error()))
	}

	return &res, nil
}
