package stock

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/Dlacreme/httpd/back/qbuilder"
	"github.com/Dlacreme/httpd/back/wdb"
	"github.com/Dlacreme/httpd/back/werror"
)

// Stock allow to build product
type Stock struct {
	StockId       int    `db:"StockId"`
	Label         string `db:"Label"`
	Quantity      int    `db:"Quantity"`
	Size          int    `db:"Size"`
	TypeId        int    `db:"StockTypeId"`
	TypeLabel     string `db:"StockTypeLabel"`
	CategoryId    int    `db:"CategoryId"`
	CategoryLabel string `db:"CategoryLabel"`
	FormatId      int    `db:"StockFormatId"`
	FormatLabel   string `db:"StockFormatLabel"`
}

func QueryGet() string {
	return `SELECT s.StockId, s.Label, s.Quantity, s.Size,
	s.TypeId AS 'StockTypeId', st.Label AS 'StockTypeLabel',
	s.CategoryId, c.Label AS 'CategoryLabel',
	s.FormatId AS 'StockFormatId', sf.Label AS 'StockFormatLabel'
	FROM Stock s
	INNER JOIN Category c ON s.CategoryId = c.CategoryId
	INNER JOIN StockType st ON st.StockTypeId = s.TypeId
	INNER JOIN StockFormat sf ON sf.StockFormatId = s.FormatId`
}

// Get stock
func Get(db wdb.Connection, search string) (*[]Stock, *werror.Error) {
	res := []Stock{}

	q := QueryGet()

	err := db.Select(&res, q)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, werror.New(500, fmt.Sprintf("Cannot get stock : %s", err.Error()))
	}

	return &res, nil
}

// GetByUniq stock by uniq
func GetByUniq(db wdb.Connection, label string, typeId int, formatId int) (*Stock, *werror.Error) {

	item := Stock{}

	q := QueryGet() + fmt.Sprintf(`
	WHERE 1 = 1
	AND LOWER(s.Label) LIKE '%s'
	AND s.TypeId = %d
	AND s.FormatId = %d
	LIMIT 1`, strings.ToLower(label), typeId, formatId)

	err := db.Get(&item, q)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, werror.New(500, fmt.Sprintf("Cannot get stock : %s", err.Error()))
	}

	return &item, nil
}

func Update(db wdb.Connection, stock []*Stock) *werror.Error {

	if len(stock) == 0 {
		return nil
	}

	queries := make([]string, len(stock))
	for i := range stock {
		queries[i] = fmt.Sprintf(`UPDATE Stock
			SET Label = '%s'
			, Quantity = %d
			, Size = %d
			, TypeId = %d
			, CategoryId = %d
			, FormatId = %d
			WHERE StockId = %d
			`, stock[i].Label, stock[i].Quantity, stock[i].Size, stock[i].TypeId, stock[i].CategoryId, stock[i].FormatId, stock[i].StockId)
	}

	_, e := db.Exec(strings.Join(queries, ";"))
	if e != nil {
		return werror.New(500, fmt.Sprintf("Cannot update stocks : %s", e.Error()))
	}

	return nil
}

func Insert(db wdb.Connection, stock []*Stock) *werror.Error {
	if len(stock) == 0 {
		return nil
	}

	fields := []string{"Label", "Quantity", "Size", "TypeId", "CategoryId", "FormatId"}
	q := qbuilder.InsertInto("Stock", fields)

	for i := range stock {
		item := []string{qbuilder.FormatStr(stock[i].Label), strconv.Itoa(stock[i].Quantity), strconv.Itoa(stock[i].Size), strconv.Itoa(stock[i].TypeId), strconv.Itoa(stock[i].CategoryId), strconv.Itoa(stock[i].FormatId)}
		q += qbuilder.AddInto(q, item)
	}

	_, e := db.Exec(q)

	if e != nil {
		return werror.New(500, fmt.Sprintf("Cannot create new stock : %s", e.Error()))
	}

	return nil
}
