package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	ct "yana/api/common/controller"
	"yana/core/product/product"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/back/werror"
	"github.com/Dlacreme/httpd/back/wesult"
)

func CreateProductsFromStock(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	var ctr []product.CreateProductFromStockCtr
	err := json.NewDecoder(r.Body).Decode(&ctr)

	if err != nil {
		ct.Final(w, wesult.New(nil, werror.New(400, fmt.Sprintf("Invalid input : %s", err.Error()))))
		return
	}

	ct.Final(w, product.CreateFromStock(&c, ctr))
}
