package product

import (
	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/back/wesult"

	"yana/core/model/product"
)

func Get(c *flight.Info, search string) wesult.Result {
	res, err := product.Get(c.DB, search)
	if err != nil {
		return wesult.New(nil, err)
	}
	return wesult.New(res, nil)
}

func GetOrphan(c *flight.Info) wesult.Result {
	res, err := product.GetOrphan(c.DB)
	if err != nil {
		return wesult.New(nil, err)
	}
	return wesult.New(res, nil)
}

func CreateFromStock(c *flight.Info, input []CreateProductFromStockCtr) wesult.Result {

	var newPdt = make([]*product.Product, len(input))

	for i := range input {
		newPdt[i] = buildProduct(input[i])
	}

	product.InsertMultiple(c.DB, newPdt)

	return wesult.New(newPdt, nil)
}

func AddCompo(c *flight.Info, input AddCompoCtr) wesult.Result {
	err := product.AddCompo(c.DB, input.ProductId, buildCompo(input.Stocks))
	if err != nil {
		return wesult.New(nil, err)
	}
	return wesult.New(nil, nil)
}

func RemoveCompo(c *flight.Info, pdtStockId int) wesult.Result {
	err := product.RemoveCompo(c.DB, pdtStockId)
	if err != nil {
		return wesult.New(nil, err)
	}
	return wesult.New(nil, nil)
}

func DisableProduct(c *flight.Info, id int) wesult.Result {
	err := product.DisableMultiple(c.DB, []int{id})
	if err != nil {
		return wesult.New(nil, err)
	}
	return wesult.New(nil, nil)
}

func GetId(c *flight.Info, id int) wesult.Result {
	pdt, err := product.GetId(c.DB, id)
	return wesult.New(pdt, err)
}
