package product

type StockDetailsCtr struct {
	StockId  int `json:"stockId"`
	Quantity int `json:"quantity"`
}

type CreateProductFromStockCtr struct {
	Label  string            `json:"label"`
	Price  float64           `json:"price"`
	TypeId int               `json:"typeId"`
	Stocks []StockDetailsCtr `json:"stocks"`
}

type AddCompoCtr struct {
	ProductId int               `json:"productId"`
	Stocks    []StockDetailsCtr `json:"stocks"`
}
