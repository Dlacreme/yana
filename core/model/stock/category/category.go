package category

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Dlacreme/httpd/back/wdb"
	"github.com/Dlacreme/httpd/back/werror"
)

// Category db model
type Category struct {
	CategoryId int    `db:"CategoryId"`
	Label      string `db:"Label"`
	UnitId     int8   `db:"UnitId"`
	Unit       string `db:"UnitLabel"`
}

// Get get all category matchin the query. if search = nul, get all
func Get(db wdb.Connection, search *string) (*[]Category, *werror.Error) {
	res := []Category{}

	q := `SELECT c.CategoryId, c.Label, c.UnitId, u.Label AS 'UnitLabel'
		FROM Category c
		INNER JOIN CategoryUnit u ON c.UnitId = u.CategoryUnitId`
	if search != nil {
		q += fmt.Sprintf(" WHERE LOWER(c.Label) LIKE '%%s%'", strings.ToLower(*search))
	}

	err := db.Select(&res, q)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, werror.New(500, fmt.Sprintf("Cannot get categories : %s", err.Error()))
	}

	return &res, nil
}

// GetByLabels From labels
func GetByLabels(db wdb.Connection, labels []string) (*[]Category, *werror.Error) {
	res := []Category{}

	if len(labels) == 0 {
		return nil, nil
	}

	q := fmt.Sprintf(`SELECT c.CategoryId, c.Label, c.UnitId, u.Label AS 'UnitLabel'
			FROM Category c
			INNER JOIN CategoryUnit u ON c.UnitId = u.CategoryUnitId
			WHERE LOWER(c.Label) IN ('%s')`, strings.Join(labels, "','"))

	err := db.Select(&res, q)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, werror.New(500, fmt.Sprintf("Cannot get categories : %s", err.Error()))
	}

	return &res, nil
}

// InsertCategories will create new categories
func InsertCategories(db wdb.Connection, categories []Category) (map[string]int, *werror.Error) {
	return nil, nil
}
