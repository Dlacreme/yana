package product

import (
	"yana/core/model/enum"
	"yana/core/model/product"
)

func buildProduct(item CreateProductFromStockCtr) *product.Product {

	typeId, typeLabel := enum.GetKey(enum.ProductType, item.TypeId)

	compos := make([]product.ProductStock, len(item.Stocks))

	for i := range item.Stocks {
		compos[i] = product.ProductStock{item.Stocks[i].StockId, -1, "Undefined", item.Stocks[i].Quantity, -1, -1}
	}

	pdt := product.Product{-1, item.Label, item.Price, typeId, typeLabel, compos}

	return &pdt
}

func buildCompo(input []StockDetailsCtr) []product.ProductStock {
	compos := make([]product.ProductStock, len(input))

	for i := range input {
		compos[i] = product.ProductStock{input[i].StockId, -1, "Undefined", input[i].Quantity, -1, -1}
	}

	return compos
}
