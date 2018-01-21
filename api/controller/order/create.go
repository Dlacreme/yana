package order

import (
	"encoding/json"
	"fmt"
	"net/http"

	ct "yana/api/common/controller"
	"yana/core/order"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/back/werror"
	"github.com/Dlacreme/httpd/back/wesult"
)

func Create(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	var ctr order.NewOrderCtr

	err := json.NewDecoder(r.Body).Decode(&ctr)

	if err != nil {
		ct.Final(w, wesult.New(nil, werror.New(400, fmt.Sprintf("Invalid input : %s", err.Error()))))
		return
	}

	ct.Final(w, order.Create(&c, ctr))
}

func AddProduct(w http.ResponseWriter, r *http.Request) {

}

func RemoveProduct(w http.ResponseWriter, r *http.Request) {

}
