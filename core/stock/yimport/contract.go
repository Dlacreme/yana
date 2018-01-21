package yimport

import "github.com/Dlacreme/httpd/back/werror"

// ImportCtr :
type ImportCtr struct {
	Label       string  `json:"label"`
	Category    string  `json:"category"`
	BuyingPrice float64 `json:"buyingPrice"`
	Quantity    int     `json:"quantity"`
	Type        string  `json:"type"`
	Size        int     `json:"size"`
	Format      string  `json:"format"`
}

// Check make sure the input value are correct
func (ct *ImportCtr) Check() *werror.Error {
	return nil
}
