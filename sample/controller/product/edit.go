package product

import (
	"net/http"
	"strconv"
	mproduct "yana/core/model/product"
	"yana/core/product/product"

	"github.com/Dlacreme/httpd/back/flight"
)

func LoadEdit(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	c.Param("id")
	id, parseError := strconv.Atoi(c.Param("id"))
	if parseError != nil {
		c.FlashWarning("Invalid id")
	}
	v := c.View.New("product/edit")
	res := product.GetId(&c, id)

	res.Partial(w, r, v)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	id, parseError := strconv.Atoi(c.Param("id"))
	if parseError != nil {
		c.FlashWarning("Invalid id")
		http.Redirect(w, r, "/product", http.StatusTemporaryRedirect)
		return
	}

	price, parseError := strconv.ParseFloat(r.FormValue("price"), 64)
	if parseError != nil {
		c.FlashWarning("Invalid price")
		http.Redirect(w, r, "/product", http.StatusTemporaryRedirect)
		return
	}
	t, err := strconv.Atoi(r.FormValue("type"))
	if err != nil {
		c.FlashWarning("Invalid type")
		http.Redirect(w, r, "/product", http.StatusTemporaryRedirect)
		return
	}

	pdt := mproduct.Product{id, r.FormValue("label"), price, t, "", []mproduct.ProductStock{}}

	ee := pdt.Save(c.DB)
	if ee != nil {
		c.FlashWarning("Cannot update product")
	}

	http.Redirect(w, r, "/product", http.StatusTemporaryRedirect)
}
