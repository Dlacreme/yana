package yimport

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"yana/core/model/stock/stock"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/back/werror"
	"github.com/Dlacreme/httpd/back/wesult"
)

// Import :
func Import(c *flight.Info, data []ImportCtr) wesult.Result {

	sts := make([]*stock.Stock, len(data))
	err := dataToStock(c.DB, &data, &sts)
	if err != nil {
		return wesult.New(nil, err)
	}

	err = fillExistingIds(c.DB, &sts)
	if err != nil {
		return wesult.New(nil, err)
	}

	var new []*stock.Stock
	var up []*stock.Stock

	for i := range sts {
		if sts[i].StockId > 0 {
			up = append(up, sts[i])
			continue
		}
		new = append(new, sts[i])
	}

	e := stock.Update(c.DB, up)
	if e != nil {
		return wesult.New(nil, e)
	}

	e = stock.Insert(c.DB, new)
	if e != nil {
		return wesult.New(nil, e)
	}

	return wesult.New(sts, nil)
}

func ImportFile(c *flight.Info, fileUrl string) wesult.Result {

	var data []ImportCtr

	// Read line by line
	inFile, _ := os.Open(fileUrl)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	// Remove header
	scanner.Scan()

	for scanner.Scan() {
		values := strings.Split(strings.TrimSpace(scanner.Text()), ",")
		price, priceerror := strconv.ParseFloat(values[2], 64)
		if priceerror != nil {
			return wesult.New(nil, werror.New(401, fmt.Sprintf("Invalid price for : %s\n", values[0])))
		}
		quantity, quantityerror := strconv.Atoi(values[3])
		if quantityerror != nil {
			return wesult.New(nil, werror.New(401, fmt.Sprintf("Invalid quantity for : %s\n", values[0])))
		}
		size, sizeerror := strconv.Atoi(values[5])
		if sizeerror != nil {
			return wesult.New(nil, werror.New(401, fmt.Sprintf("Invalid size for : %s\n", values[0])))
		}
		data = append(data, ImportCtr{values[0], values[1], price, quantity, values[4], size, values[6]})
	}

	res := Import(c, data)

	if res.Error != nil {
		return res
	}

	// Save file

	return res
}
