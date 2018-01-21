package yimport

import (
	"fmt"
	"strings"

	"yana/core/model/enum"
	"yana/core/model/stock/category"
	"yana/core/model/stock/stock"

	"github.com/Dlacreme/httpd/back/wdb"
	"github.com/Dlacreme/httpd/back/werror"
)

func getCategory(d *[]category.Category, label string) (int, string) {
	data := *d
	for i := range data {
		if strings.ToLower(data[i].Label) == strings.ToLower(label) {
			return data[i].CategoryId, data[i].Label
		}
	}
	return 0, "Undefined"
}

func dataToStock(db wdb.Connection, _input *[]ImportCtr, _output *[]*stock.Stock) *werror.Error {
	existingCats, err := category.Get(db, nil)
	if err != nil {
		return err
	}

	input := *_input
	output := *_output

	for i := range input {
		catId, catLabel := getCategory(existingCats, input[i].Category)

		if catId == -1 {
			return werror.New(400, fmt.Sprintf("Unknown category : %s", input[i].Category))
		}

		formatId, formatLabel := enum.GetValue(enum.StockFormat, input[i].Format)
		typeId, typeLabel := enum.GetValue(enum.StockType, input[i].Type)

		output[i] = &stock.Stock{0, input[i].Label, input[i].Quantity, input[i].Size, typeId, typeLabel, catId, catLabel, formatId, formatLabel}
	}

	return nil
}

func fillExistingIds(db wdb.Connection, _output *[]*stock.Stock) *werror.Error {

	output := *_output

	for i := range output {
		item, err := stock.GetByUniq(db, output[i].Label, output[i].TypeId, output[i].FormatId)
		if err != nil {
			return err
		}
		if item != nil {
			output[i].Quantity += item.Quantity
			output[i].StockId = item.StockId
		}
	}

	return nil
}
