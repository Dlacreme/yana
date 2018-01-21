package product

import (
	"net/http"
	"strconv"
	"yana/core/product/product"

	"github.com/Dlacreme/httpd/back/flight"
)

func LoadAddCompo(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	c.Param("id")
	id, parseError := strconv.Atoi(c.Param("id"))
	if parseError != nil {
		c.FlashWarning("Invalid id")
	}

	v := c.View.New("product/compoadd")
	res := product.GetId(&c, id)

	res.Partial(w, r, v)
}

func AddCompo(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	c.Param("id")
	id, parseError := strconv.Atoi(c.Param("id"))
	if parseError != nil {
		c.FlashWarning("Invalid id")
	}

	if !c.FormValid("stock", "quantity") {
		c.FlashWarning("Invalid input")
		res, v := loadData(&c)
		res.RenderView(w, r, v)
		return
	}

	var stockId int
	var quantity int

	stockId, parseError = strconv.Atoi(r.FormValue("stock"))
	if parseError != nil {
		c.FlashWarning("Invalid stockid")
		http.Redirect(w, r, "/product", http.StatusTemporaryRedirect)
		return
	}
	quantity, parseError = strconv.Atoi(r.FormValue("quantity"))
	if parseError != nil {
		c.FlashWarning("Invalid quantity")
		http.Redirect(w, r, "/product", http.StatusTemporaryRedirect)
		return
	}

	s := product.StockDetailsCtr{stockId, quantity}
	data := product.AddCompo(&c, product.AddCompoCtr{id, []product.StockDetailsCtr{s}})

	if data.Error != nil {
		c.FlashWarning("Unable to insert new stock")
	}
	http.Redirect(w, r, "/product", http.StatusTemporaryRedirect)
	return
}

func RemoveCompo(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	psId, perr := strconv.Atoi(c.Param("productstockid"))
	if perr != nil {
		c.FlashWarning("Invalid stock id")
		http.Redirect(w, r, "/product", http.StatusTemporaryRedirect)
		return
	}
	data := product.RemoveCompo(&c, psId)

	if data.Error != nil {
		c.FlashWarning("Unable to delete compo")
	}

	http.Redirect(w, r, "/product", http.StatusTemporaryRedirect)
}
