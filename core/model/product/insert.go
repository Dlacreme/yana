package product

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Dlacreme/httpd/back/qbuilder"
	"github.com/Dlacreme/httpd/back/wdb"
	"github.com/Dlacreme/httpd/back/werror"
)

func InsertMultiple(db wdb.Connection, pdts []*Product) *werror.Error {

	if len(pdts) == 0 {
		return nil
	}

	fields := []string{"Label", "TypeId"}
	q := qbuilder.InsertInto("Product", fields)

	for i := range pdts {
		item := []string{qbuilder.FormatStr(pdts[i].Label), strconv.Itoa(pdts[i].TypeId)}
		q += qbuilder.AddInto(q, item)
	}

	r, e := db.Exec(q)

	if e != nil {
		return werror.New(500, fmt.Sprintf("Cannot create new Product : %s", e.Error()))
	}

	lastId, _ := r.LastInsertId()

	for i := range pdts {
		pdts[i].ProductId = int(lastId)
		lastId++

		UpdateProductPrice(db, pdts[i].ProductId, pdts[i].Price)
		if len(pdts[i].Composition) > 0 {
			LinkToStock(db, pdts[i].ProductId, pdts[i].Composition)
		}
	}

	return nil
}

func DisableMultiple(db wdb.Connection, ids []int) *werror.Error {
	id := make([]string, len(ids))
	for i := range ids {
		id[i] = strconv.Itoa(ids[i])
	}
	q := fmt.Sprintf(`UPDATE Product SET IsDisabled = 1 WHERE ProductId IN (%s)`, strings.Join(id, ","))
	_, err := db.Exec(q)

	if err != nil {
		return werror.New(500, fmt.Sprintf("Cannot disable products : %s", err.Error()))
	}
	return nil
}

func (pdt *Product) Save(db wdb.Connection) *werror.Error {
	q := fmt.Sprintf(`UPDATE Product SET Label = '%s', TypeId = %d WHERE ProductId = %d`, pdt.Label, pdt.TypeId, pdt.ProductId)

	_, err := db.Exec(q)

	if err != nil {
		return werror.New(500, err.Error())
	}

	e := UpdateProductPrice(db, pdt.ProductId, pdt.Price)
	if e != nil {
		return e
	}

	return nil
}
