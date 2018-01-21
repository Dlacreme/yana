package order

type NewOrderCtr struct {
	TableId   int   `json:"tableId"`
	CreatedBy int   `json:"createdBy"`
	Products  []int `json:"products"`
}
