package stock

import (
	"encoding/json"
	"fmt"
	"net/http"

	ct "yana/api/common/controller"
	"yana/core/stock/yimport"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/back/werror"
	"github.com/Dlacreme/httpd/back/wesult"
)

// Import will dynamically create of import a new product
func Import(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	var ctr []yimport.ImportCtr

	err := json.NewDecoder(r.Body).Decode(&ctr)

	if err != nil {
		ct.Final(w, wesult.New(nil, werror.New(400, fmt.Sprintf("Invalid input : %s", err.Error()))))
		return
	}

	ct.Final(w, yimport.Import(&c, ctr))
	return
}
