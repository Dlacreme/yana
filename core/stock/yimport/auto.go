package yimport

import (
	"fmt"

	"yana/core/model/stock/stock"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/back/werror"
	"github.com/Dlacreme/httpd/back/wesult"
)

func StockFinishedProductNotUsed(c *flight.Info) wesult.Result {

	res := []stock.Stock{}

	q := stock.QueryGet() + `
	LEFT JOIN ProductStock ps ON ps.StockId = s.StockId
    WHERE s.TypeId = 1`

	err := c.DB.Select(&res, q)

	if err != nil {
		return wesult.New(nil, werror.New(500, fmt.Sprintf("Cannot get unused stocks : %s", err.Error())))
	}

	return wesult.New(res, nil)
}
