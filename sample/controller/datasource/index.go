package datasource

import (
	"net/http"
	"yana/sample/viewmodel/gen"

	"github.com/Dlacreme/httpd/back/wdb"
	"github.com/Dlacreme/httpd/back/werror"
	"github.com/Dlacreme/httpd/back/wesult"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/back/router"
)

// LoadRoutes the routes.
func LoadRoutes() {
	router.Get("/list/producttype", ProductType)
	router.Get("/list/stock", Stock)
}

// Index is the landing page. Will redirect to Home.Index if user is already logged
func ProductType(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	q := "SELECT ProductTypeId AS 'id', Label AS 'text' FROM ProductType"
	res := select2FromQuery(c.DB, q)
	res.ToJson(w)
}

func Stock(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	q := "SELECT StockId AS 'id', CONCAT(s.Label, ' (', st.Label, ') ', sf.Label, ' ', s.Size, ' ', cu.Label) as 'text' FROM Stock s INNER JOIN StockType st ON st.StockTypeId = s.TypeId INNER JOIN StockFormat sf ON sf.StockFormatId = s.FormatId INNER JOIN Category c ON c.CategoryId = s.CategoryId INNER JOIN CategoryUnit cu ON cu.CategoryUnitId = c.UnitId;"
	res := select2FromQuery(c.DB, q)
	res.ToJson(w)
}

func select2FromQuery(db wdb.Connection, q string) wesult.Result {
	item := []gen.Select2Item{}

	err := db.Select(&item, q)

	if err != nil {
		return wesult.New(nil, werror.New(500, "Cannot get data"))
	}
	return wesult.New(gen.Select2{true, item}, nil)
}
