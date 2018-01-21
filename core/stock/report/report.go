package report

import (
	"yana/core/model/stock/stock"

	"github.com/Dlacreme/httpd/back/flight"

	"github.com/Dlacreme/httpd/back/wesult"
)

// Report :
func Report(c *flight.Info, search string) wesult.Result {

	stocks, err := stock.Get(c.DB, search)

	if err != nil {
		return wesult.New(nil, err)
	}

	return wesult.New(stocks, nil)
}
