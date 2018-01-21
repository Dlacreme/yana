package stock

import (
	"github.com/Dlacreme/httpd/back/router"
)

// Load set routing
func Load() {
	router.Post("/stock/import", Import)
	router.Get("/stock", Report)
	router.Get("/report/unusedstock", UnusedStock)
}
