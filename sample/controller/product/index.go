package product

import (
	"net/http"
	"strconv"
	"yana/core/product/product"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/back/router"
	"github.com/Dlacreme/httpd/back/wesult"
	"github.com/Dlacreme/httpd/view"
)

// LoadRoutes the routes.
func LoadRoutes() {
	router.Get("/product", Index)
	router.Post("/product", Index)

	router.Get("/product/create", LoadCreate)
	router.Post("/product/create", Create)

	router.Get("/product/edit/:id", LoadEdit)
	router.Post("/product/edit/:id", Edit)

	router.Get("/product/remove/:id", RemoveProduct)

	router.Get("/product/addcompo/:id", LoadAddCompo)
	router.Post("/product/addcompo/:id", AddCompo)
	router.Get("/product/removecompo/:productstockid", RemoveCompo)
}

// Index is the landing page. Will redirect to Home.Index if user is already logged
func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	res := product.Get(&c, "")
	v := c.View.New("product/index")

	res.RenderView(w, r, v)
}

func loadData(c *flight.Info) (wesult.Result, *view.Info) {
	v := c.View.New("product/index")
	res := product.Get(c, "")

	// orphan := product.GetOrphan(c)
	// v.Vars["orphans"] = orphan.Output

	return res, v
}

func RemoveProduct(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	c.Param("id")
	id, parseError := strconv.Atoi(c.Param("id"))
	if parseError != nil {
		c.FlashWarning("Invalid id")
	}

	data := product.DisableProduct(&c, id)
	if data.Error != nil {
		c.FlashWarning("Unable to delete product")
	}
	http.Redirect(w, r, "/product", http.StatusTemporaryRedirect)
	return
}
