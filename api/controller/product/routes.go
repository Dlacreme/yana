package product

import (
	"github.com/Dlacreme/httpd/back/router"
)

func Load() {
	router.Get("/product", Get)
	router.Put("/product/create", CreateProductsFromStock)
	router.Delete("/product", Delete)
}
