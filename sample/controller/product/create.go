package product

import (
	"net/http"
	"strconv"
	"yana/core/product/product"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/back/wesult"
)

func LoadCreate(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	v := c.View.New("product/create")
	res := wesult.New(nil, nil)
	res.Partial(w, r, v)
}

func Create(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	// Parse input
	if !c.FormValid("label", "price", "type") {
		c.FlashWarning("Invalid input")
		http.Redirect(w, r, "/product", http.StatusTemporaryRedirect)
		return
	}

	input := []product.CreateProductFromStockCtr{}
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

	input = append(input, product.CreateProductFromStockCtr{r.FormValue("label"), price, t, []product.StockDetailsCtr{}})

	res := product.CreateFromStock(&c, input)
	if res.Error != nil {
		c.FlashWarning("Fail to create new product")
	}

	http.Redirect(w, r, "/product", http.StatusTemporaryRedirect)
	return
}
