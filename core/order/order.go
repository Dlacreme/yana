package order

import (
	"yana/core/model/order"
	"yana/core/model/user"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/back/wesult"
)

func Create(c *flight.Info, input NewOrderCtr) wesult.Result {

	// TODO : check errors
	o := order.New(input.TableId, input.CreatedBy, user.IsStaff(c.DB, input.CreatedBy), input.Products)
	order.Create(c.DB, &o)

	return wesult.New(nil, nil)
}

func Get(c *flight.Info, search string) wesult.Result {
	r, e := order.Get(c.DB, search)
	if e != nil {
		return wesult.New(nil, e)
	}
	return wesult.New(r, nil)
}
