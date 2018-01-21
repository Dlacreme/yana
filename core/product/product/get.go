package product

import (
	"yana/core/model/product/productType"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/back/wesult"
)

func GetProductType(c *flight.Info, search string) wesult.Result {
	types, err := productType.Get(c.DB, search)
	if err != nil {
		return wesult.New(nil, err)
	}
	return wesult.New(types, nil)
}
