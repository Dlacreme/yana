package controller

import (
	"yana/sample/controller/dashboard"
	"yana/sample/controller/datasource"
	"yana/sample/controller/order"
	"yana/sample/controller/product"
	"yana/sample/controller/public"
	"yana/sample/controller/static"
	"yana/sample/controller/status"
	"yana/sample/controller/stock"
)

func LoadRoutes() {

	status.Load()
	static.Load()

	public.LoadRoutes()
	dashboard.LoadRoutes()
	order.LoadRoutes()
	product.LoadRoutes()
	stock.LoadRoutes()

	datasource.LoadRoutes()
}
