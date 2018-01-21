package product

import (
	"fmt"
	"time"

	"github.com/Dlacreme/httpd/back/wdb"
	"github.com/Dlacreme/httpd/back/werror"
)

type Price struct {
	PriceId   int
	ProductId int
	Price     float64
	StartDate time.Time
	EndDate   *time.Time
}

func UpdateProductPrice(db wdb.Connection, productId int, price float64) *werror.Error {
	// Disable previous
	disQ := fmt.Sprintf(`UPDATE Price SET EndDate = NOW() WHERE ProductId = %d AND EndDate IS NULL`, productId)
	_, e := db.Exec(disQ)
	if e != nil {
		return werror.New(500, fmt.Sprintf("Cannot disable previous price : %s", e.Error()))
	}

	// Enable new one
	enaQ := fmt.Sprintf(`INSERT INTO Price (ProductId, Price, StartDate) VALUES (%d, %g, NOW())`, productId, price)
	_, e = db.Exec(enaQ)
	if e != nil {
		return werror.New(500, fmt.Sprintf("Cannot create new price : %s", e.Error()))
	}

	return nil
}
