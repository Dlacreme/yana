package order

import "github.com/Dlacreme/httpd/back/router"

func Load() {
	router.Get("/order", Get)
	router.Put("/order", Create)
	router.Put("/order/product", AddProduct)
	router.Delete("/order/product", RemoveProduct)
	router.Post("/order/status", UpdateStatus)
	router.Post("/order/payment", UpdatePayment)
}
