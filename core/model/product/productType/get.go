package productType

import (
	"fmt"

	"github.com/Dlacreme/httpd/back/wdb"
	"github.com/Dlacreme/httpd/back/werror"
)

type ProductType struct {
	ProductTypeId int    `json:"ProductTypeId"`
	Label         string `json:"Label"`
}

func QueryGet() string {
	return "SELECT * FROM ProductType ORDER BY Label"
}

func Get(db wdb.Connection, search string) ([]ProductType, *werror.Error) {
	q := QueryGet()
	if search != "" {
		q += fmt.Sprintf(" WHERE Label LIKE '%%s%'", search)
	}
	types := []ProductType{}
	err := db.Select(&types, q)

	if err != nil {
		return nil, werror.New(500, fmt.Sprintf("Cannot get product types : %s\n", err.Error()))
	}

	return types, nil
}
