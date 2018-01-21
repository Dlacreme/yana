
/*
	ProductId int
	Label     string
	Info      string
	Price     float64
	TypeId    int
	TypeLabel string


	StockId          int
	StockLabel       string
	StockRequired   int
	StockAvailable   int
	StockSize        int
	StockTypeId      int
	StockTypeLabel   string
	StockFormatId    int
	StockFormatLabel string

	CategoryId          int
	CategoryLabel       string
	CategoryUnitId      int
	CategoryUnitLabel   string
	CategoryParentId    int
	CategoryParentLabel string
    */

SELECT
    p.ProductId,
    p.Label,
    CONCAT (sf.Label, " ", s.Size, " ", cu.Label) AS Info,
    pr.Price,
    p.TypeId,
    pt.Label AS 'TypeLabel',

    s.StockId,
    s.Label AS 'StockLabel',
    ps.Quantity AS 'StockRequired',
    s.Quantity AS 'StockAvailable',
    s.Size AS 'StockSize',
    s.TypeId AS 'StockTypeId',
    st.Label AS 'StockTypeLabel',
    s.FormatId AS 'StockFormatId',
    sf.Label AS 'StockFormatLabel',

    s.CategoryId,
    c.Label AS 'CategoryLabel',
    c.UnitId AS 'CategoryUnitId',
    cu.Label AS 'CategoryUnitLabel',
    c.ParentId AS 'CategoryParentId',
    pc.Label AS 'CategoryParentLabel'

FROM Product p
INNER JOIN ProductType pt ON p.TypeId = pt.ProductTypeId
INNER JOIN Price pr ON pr.ProductId = p.ProductId AND pr.EndDate IS NULL
INNER JOIN ProductStock ps ON ps.ProductId = p.ProductId AND ps.EndDate IS NULL
INNER JOIN Stock s ON s.StockId = ps.StockId
INNER JOIN Category c ON c.CategoryId = s.CategoryId
INNER JOIN StockType st ON st.StockTypeId = s.TypeId
INNER JOIN StockFormat sf ON sf.StockFormatId = s.FormatId
INNER JOIN CategoryUnit cu ON cu.CategoryUnitId = c.UnitId
LEFT JOIN Category pc ON pc.CategoryId = c.ParentId
