package stock

import (
	"net/http"

	ct "yana/api/common/controller"
	"yana/core/stock/report"
	"yana/core/stock/yimport"

	"github.com/Dlacreme/httpd/back/flight"
)

func Report(w http.ResponseWriter, r *http.Request) {

	c := flight.Context(w, r)

	search := r.URL.Query().Get("search")

	ct.Final(w, report.Report(&c, search))
}

func UnusedStock(w http.ResponseWriter, r *http.Request) {

	c := flight.Context(w, r)

	ct.Final(w, yimport.StockFinishedProductNotUsed(&c))
}
