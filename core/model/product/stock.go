package product

import (
	"fmt"
	"strconv"

	"github.com/Dlacreme/httpd/back/qbuilder"
	"github.com/Dlacreme/httpd/back/wdb"
	"github.com/Dlacreme/httpd/back/werror"
)

func LinkToStock(db wdb.Connection, pdtId int, data []ProductStock) *werror.Error {

	// Disable previous
	disQ := fmt.Sprintf(`UPDATE ProductStock SET EndDate = NOW() WHERE ProductId = %d AND EndDate IS NULL`, pdtId)
	_, e := db.Exec(disQ)
	if e != nil {
		return werror.New(500, fmt.Sprintf("Cannot remove previous product stock links : %s", e.Error()))
	}

	// Create new link
	fields := []string{"ProductId", "StockId", "Quantity", "StartDate"}
	q := qbuilder.InsertInto("ProductStock", fields)

	for i := range data {
		item := []string{strconv.Itoa(pdtId), strconv.Itoa(data[i].StockId), strconv.Itoa(data[i].RequiredQuantity), "NOW()"}
		q += qbuilder.AddInto(q, item)
	}

	_, e = db.Exec(q)

	if e != nil {
		return werror.New(500, fmt.Sprintf("Cannot create new product stock links : %s", e.Error()))
	}

	return nil
}

func AddCompo(db wdb.Connection, pdtId int, data []ProductStock) *werror.Error {
	// Create new link
	fields := []string{"ProductId", "StockId", "Quantity", "StartDate"}
	q := qbuilder.InsertInto("ProductStock", fields)

	for i := range data {
		item := []string{strconv.Itoa(pdtId), strconv.Itoa(data[i].StockId), strconv.Itoa(data[i].RequiredQuantity), "NOW()"}
		q += qbuilder.AddInto(q, item)
	}

	_, e := db.Exec(q)

	if e != nil {
		return werror.New(500, fmt.Sprintf("Cannot create new product stock links : %s", e.Error()))
	}

	return nil
}

func RemoveCompo(db wdb.Connection, pdtStockId int) *werror.Error {
	// Disable previous
	disQ := fmt.Sprintf(`UPDATE ProductStock SET EndDate = NOW() WHERE ProductStockId = %d AND EndDate IS NULL`, pdtStockId)
	_, e := db.Exec(disQ)
	if e != nil {
		return werror.New(500, fmt.Sprintf("Cannot remove previous product stock links : %s", e.Error()))
	}
	return nil
}
