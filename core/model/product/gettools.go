package product

import "github.com/Dlacreme/httpd/back/werror"

func QueryGet() string {
	return `SELECT
    p.ProductId,
    p.Label,
    IFNULL(CONCAT (sf.Label, " ", s.Size, " ", cu.Label), "Undefined") AS Info,
    pr.Price,
    p.TypeId,
	pt.Label AS 'TypeLabel',
	
    IFNULL(ps.ProductStockId, 0) AS 'ProductStockId',

    IFNULL(s.StockId, '0') AS 'StockId',
	IFNULL(CONCAT(s.Label, " ", sf.Label, " - ", s.Size, " " , cu.Label, " (", st.Label, ")"), '') AS 'StockLabel',	
    IFNULL(ps.Quantity, 0) AS 'StockRequired',
    IFNULL(s.Quantity, 0) AS 'StockAvailable',
    IFNULL(s.Size, 0) AS 'StockSize',
    IFNULL(s.TypeId, 0) AS 'StockTypeId',
    IFNULL(st.Label, '') AS 'StockTypeLabel',
    IFNULL(s.FormatId, 0) AS 'StockFormatId',
    IFNULL(sf.Label, '') AS 'StockFormatLabel',

    IFNULL(s.CategoryId, 0) AS 'CategoryId',
    IFNULL(c.Label, '') AS 'CategoryLabel',
    IFNULL(c.UnitId, 0) AS 'CategoryUnitId',
    IFNULL(cu.Label, '') AS 'CategoryUnitLabel',
    c.ParentId AS 'CategoryParentId',
    pc.Label AS 'CategoryParentLabel'

	FROM Product p
	INNER JOIN ProductType pt ON p.TypeId = pt.ProductTypeId AND p.IsDisabled = 0
	INNER JOIN Price pr ON pr.ProductId = p.ProductId AND pr.EndDate IS NULL
	LEFT JOIN ProductStock ps ON ps.ProductId = p.ProductId AND ps.EndDate IS NULL
	LEFT JOIN Stock s ON s.StockId = ps.StockId
	LEFT JOIN Category c ON c.CategoryId = s.CategoryId
	LEFT JOIN StockType st ON st.StockTypeId = s.TypeId
	LEFT JOIN StockFormat sf ON sf.StockFormatId = s.FormatId
	LEFT JOIN CategoryUnit cu ON cu.CategoryUnitId = c.UnitId
	LEFT JOIN Category pc ON pc.CategoryId = c.ParentId `
}

func QueryGetOrphanProduct() string {
	return `SELECT
    p.ProductId AS 'ProductId',
    p.Label,
    p.TypeId,
	pt.Label AS 'Type',
	pr.Price AS 'Price'
	FROM Product p
	INNER JOIN ProductType pt ON pt.ProductTypeId = p.TypeId AND p.IsDisabled = 0
	INNER JOIN Price pr ON pr.ProductId = p.ProductId AND pr.EndDate IS NULL
    WHERE NOT EXISTS (
        SELECT ProductId
        FROM ProductStock ps
        WHERE p.ProductId = ps.ProductId
    )`
}

type ProductTmp struct {
	ProductId int     `db:"ProductId"`
	Label     string  `db:"Label"`
	Info      string  `db:"Info"`
	Price     float64 `db:"Price"`
	TypeId    int     `db:"TypeId"`
	TypeLabel string  `db:"TypeLabel"`

	ProductStockId int `db:"ProductStockId"`

	StockId          int    `db:"StockId"`
	StockLabel       string `db:"StockLabel"`
	StockRequired    int    `db:"StockRequired"`
	StockAvailable   int    `db:"StockAvailable"`
	StockSize        int    `db:"StockSize"`
	StockTypeId      int    `db:"StockTypeId"`
	StockTypeLabel   string `db:"StockTypeLabel"`
	StockFormatId    int    `db:"StockFormatId"`
	StockFormatLabel string `db:"StockFormatLabel"`

	CategoryId          int     `db:"CategoryId"`
	CategoryLabel       string  `db:"CategoryLabel"`
	CategoryUnitId      int     `db:"CategoryUnitId"`
	CategoryUnitLabel   string  `db:"CategoryUnitLabel"`
	CategoryParentId    *int    `db:"CategoryParentId"`
	CategoryParentLabel *string `db:"CategoryParentLabel"`
}

func buildProductFromTmp(tmp []ProductTmp) *Product {
	if len(tmp) == 0 {
		return nil
	}
	compo := make([]ProductStock, len(tmp))
	for i := range tmp {
		compo[i] = ProductStock{tmp[i].StockId, tmp[i].ProductStockId, tmp[i].StockLabel, tmp[i].StockRequired, tmp[i].StockAvailable, tmp[i].StockTypeId}
	}
	return &Product{tmp[0].ProductId, tmp[0].Label, tmp[0].Price, tmp[0].TypeId, tmp[0].TypeLabel, compo}
}

func processBuild(tmp []ProductTmp) ([]*Product, *werror.Error) {
	var res []*Product

	l := len(tmp)

	if l == 0 {
		return res, nil
	}

	var s int = 0
	var id int = tmp[0].ProductId

	for i := range tmp {
		if tmp[i].ProductId != id {
			res = append(res, buildProductFromTmp(tmp[s:i]))
			s = i
			id = tmp[i].ProductId
		}
		if i == l-1 {
			res = append(res, buildProductFromTmp(tmp[s:]))
		}
	}

	return res, nil
}
