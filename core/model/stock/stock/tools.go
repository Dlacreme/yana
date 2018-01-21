package stock

import (
	"fmt"
	"strings"

	"github.com/Dlacreme/httpd/back/wdb"
	"github.com/Dlacreme/httpd/back/werror"
)

// isExisting checks if the stock is already existing
func isExisting(db wdb.Connection, label string, typeId int, formatId int) (bool, *werror.Error) {
	var res int
	err := db.Get(&res, fmt.Sprintf(`
			SELECT COUNT(StockId)
			FROM Stock 
			WHERE Label LIKE '%s'
			AND TypeId = %d
			AND FormatId = %d
		`, strings.ToLower(label), typeId, formatId))
	if err != nil {
		return true, werror.New(500, fmt.Sprintf("Cannot check is stock is exsting : %s", err.Error()))
	}
	if res > 0 {
		return true, nil
	}
	return false, nil
}
