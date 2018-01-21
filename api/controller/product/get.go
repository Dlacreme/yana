package product

import (
	"net/http"
	"strconv"

	ct "yana/api/common/controller"
	"yana/core/product/product"

	"github.com/Dlacreme/httpd/back/flight"
)

func Get(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	search := r.URL.Query().Get("search")

	ct.Final(w, product.Get(&c, search))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	id := r.URL.Query().Get("id")
	idInt, _ := strconv.Atoi(id)
	ct.Final(w, product.DisableProducts(&c, idInt))
}
