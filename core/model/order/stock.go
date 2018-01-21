package order

import (
	"fmt"
	"strconv"

	"github.com/Dlacreme/httpd/back/qbuilder"
	"github.com/Dlacreme/httpd/back/wdb"
	"github.com/Dlacreme/httpd/back/werror"
)

func AddProduct(db wdb.Connection, orderId int, products []OrderProduct, createdId int) *werror.Error {
	fields := []string{"ProductId", "OrderId", "ExtraOrderProductId", "CreatedBy", "CreatedAt"}
	q := qbuilder.InsertInto("OrderProduct", fields)

	for i := range products {
		item := []string{strconv.Itoa(products[i].ProductId), strconv.Itoa(orderId), "NULL", strconv.Itoa(createdId), "NOW()"}
		q += qbuilder.AddInto(q, item)
	}
	_, e := db.Exec(q)
	if e != nil {
		return werror.New(500, fmt.Sprintf("Cannot add products to order : %s", e.Error()))
	}
	return nil
}
