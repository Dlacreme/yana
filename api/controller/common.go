package controller

import (
	"yana/api/controller/home"
	"yana/api/controller/order"
	"yana/api/controller/product"
	"yana/api/controller/session"
	"yana/api/controller/static"
	"yana/api/controller/status"
	"yana/api/controller/stock"
)

// LoadRoutes sets the app routing
func LoadRoutes() {
	static.Load()
	home.Load()
	status.Load()

	session.Load()
	stock.Load()
	product.Load()
	order.Load()
}
