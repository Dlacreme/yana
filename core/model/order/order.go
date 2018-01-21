package order

import (
	"fmt"
	"strconv"
	"time"

	"yana/core/model/enum"
	"yana/core/model/product"

	"github.com/Dlacreme/httpd/back/qbuilder"
	"github.com/Dlacreme/httpd/back/wdb"
	"github.com/Dlacreme/httpd/back/werror"
)

type OrderProductComposition struct {
}

type OrderProduct struct {
	OrderProductId int
	ProductId      int
	Product        *product.Product
}

type Order struct {
	OrderId            int       `db:"OrderId"`
	TableId            int       `db:"TableId"`
	TableLabel         string    `db:"TableLabel"`
	StaffId            int       `db:"StaffId"`
	StaffName          string    `db:"StaffName"`
	MemberId           int       `db:"MemberId"`
	MemberName         string    `db:"MemberName"`
	IsPaid             bool      `db:"IsPaid"`
	TotalPrice         float64   `db:"TotalPrice"`
	AmountPaid         float64   `db:"AmountPaid"`
	PaymentMethodId    int       `db:"PaymentMethodId"`
	PaymentMethodLabel string    `db:"PaymentMethodLabel"`
	InvoiceId          string    `db:"InvoiceId"`
	CreatedBy          int       `db:"CreatedBy"`
	UpdatedBy          int       `db:"UpdatedBy"`
	CreatedAt          time.Time `db:"CreatedAt"`
	UpdatedAt          time.Time `db:"UpdatedAt"`
	Products           []OrderProduct
}

func New(tableId int, createdBy int, fromMember bool, products []int) Order {

	tId, tLabel := enum.GetKey(enum.Table, tableId)
	var staffId int = 0
	var memberId int = 0
	if fromMember {
		memberId = createdBy
	} else {
		staffId = createdBy
	}
	no := Order{0, tId, tLabel, staffId, "UNDEFINED STAFF", memberId, "UNDEFINED MEMBER", false, 0, 0, 0, "No payment defined", "", createdBy, createdBy, time.Now(), time.Now(), make([]OrderProduct, len(products))}

	for i := range products {
		no.Products[i] = OrderProduct{0, products[i], nil}
	}

	return no
}

func Create(db wdb.Connection, newOrder *Order) *werror.Error {
	fields := []string{"TableId", "MemberId", "StaffId", "IsPaid", "TotalPrice", "AmountPaid", "PaymentMethodId", "InvoiceId", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt"}
	q := qbuilder.InsertInto("`Order`", fields)
	item := []string{strconv.Itoa(newOrder.TableId),
		qbuilder.ToStringOrNull(newOrder.MemberId),
		qbuilder.ToStringOrNull(newOrder.StaffId),
		qbuilder.BoolAsIntString(newOrder.IsPaid),
		strconv.FormatFloat(newOrder.TotalPrice, 'G', -1, 64),
		strconv.FormatFloat(newOrder.AmountPaid, 'G', -1, 64),
		qbuilder.ToStringOrNull(newOrder.PaymentMethodId),
		qbuilder.FormatStr(newOrder.InvoiceId),
		strconv.Itoa(newOrder.CreatedBy),
		strconv.Itoa(newOrder.UpdatedBy),
		"NOW()",
		"NOW()"}
	q += qbuilder.AddInto(q, item)

	r, e := db.Exec(q)

	if e != nil {
		return werror.New(500, fmt.Sprintf("Cannot create order : %s", e.Error()))
	}

	id, _ := r.LastInsertId()
	newOrder.OrderId = int(id)

	err := AddProduct(db, newOrder.OrderId, newOrder.Products, newOrder.CreatedBy)
	if err != nil {
		return err
	}

	return nil
}

func Get(db wdb.Connection, search string) ([]*Order, *werror.Error) {

	tmp := []OrderGetTmp{}
	q := QueryGet()

	err := db.Select(&tmp, q)

	if err != nil {
		return nil, werror.New(500, fmt.Sprintf("Cannot get orders : %s", err.Error()))
	}
	r, e := parse(tmp)
	if e != nil {
		return nil, e
	}

	return r, nil
}
