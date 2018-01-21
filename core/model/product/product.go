package product

import (
	"fmt"
	"strconv"

	"github.com/Dlacreme/httpd/back/wdb"
	"github.com/Dlacreme/httpd/back/werror"
)

// ProductStock is the product composition details
type ProductStock struct {
	StockId           int
	ProductStockId    int
	Label             string
	RequiredQuantity  int
	AvailableQuantity int
	StockTypeId       int
}

// Product is directly available in the menu
type Product struct {
	ProductId   int     `db:"ProductId"`
	Label       string  `db:"Label"`
	Price       float64 `db:"Price"`
	TypeId      int     `db:"TypeId"`
	Type        string  `db:"Type"`
	Composition []ProductStock
}

func Get(db wdb.Connection, search string) ([]*Product, *werror.Error) {
	q := QueryGet()
	if search != "" {
		q += fmt.Sprintf(" WHERE p.Label LIKE '%%s%' OR c.Label LIKE '%%s%' OR s.Label LIKE '%%s%'", search, search, search)
	}

	tmp := []ProductTmp{}
	err := db.Select(&tmp, q)

	if err != nil {
		return nil, werror.New(500, fmt.Sprintf("Cannot get product : %s", err.Error()))
	}

	return processBuild(tmp)
}

func GetId(db wdb.Connection, id int) (*Product, *werror.Error) {

	q := QueryGet()
	q += " WHERE p.ProductId = " + strconv.Itoa(id)

	tmp := []ProductTmp{}
	err := db.Select(&tmp, q)

	if err != nil {
		return nil, werror.New(500, fmt.Sprintf("Cannot get product : %s", err.Error()))
	}

	res, e := processBuild(tmp)

	if len(res) > 0 {
		return res[0], e
	}

	return nil, e
}

func GetWhere(db wdb.Connection, where string) ([]Product, *werror.Error) {
	var res []Product

	q := QueryGet()
	q += where

	tmp := []ProductTmp{}
	err := db.Select(&tmp, q)

	if err != nil {
		return nil, werror.New(500, fmt.Sprintf("Cannot get product : %s", err.Error()))
	}

	return res, nil
}

func GetOrphan(db wdb.Connection) ([]Product, *werror.Error) {
	q := QueryGetOrphanProduct()
	res := []Product{}
	err := db.Select(&res, q)

	if err != nil {
		return nil, werror.New(500, fmt.Sprintf("Cannot get orphan products : %s", err.Error()))
	}

	return res, nil
}
