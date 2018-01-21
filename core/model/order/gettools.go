package order

import (
	"time"

	"yana/core/model/product"

	"github.com/Dlacreme/httpd/back/wdb"
	"github.com/Dlacreme/httpd/back/werror"
)

type OrderGetTmp struct {
	OrderId            int     `db:"OrderId"`
	TableId            int     `db:"TableId"`
	TableLabel         string  `db:"TableLabel"`
	StaffId            int     `db:"StaffId"`
	StaffName          string  `db:"StaffName"`
	MemberId           int     `db:"MemberId"`
	MemberName         string  `db:"MemberName"`
	FromMember         bool    `db:"IsFromApp"`
	StatusId           int     `db:"StatusId"`
	StatusLabel        string  `db:"StatusLabel"`
	IsPaid             bool    `db:"IsPaid"`
	TotalPrice         float64 `db:"TotalPrice"`
	AmountPaid         float64 `db:"AmountPaid"`
	PaymentMethodId    int     `db:"PaymentMethodId"`
	PaymentMethodLabel string  `db:"PaymentMethodLabel"`
	InvoiceId          string  `db:"InvoiceId"`
	CreatedBy          int     `db:"CreatedBy"`
	UpdatedBy          int     `db:"UpdatedBy"`
	CreatedAt          string  `db:"CreatedAt"`
	UpdatedAt          string  `db:"UpdatedAt"`

	ProductStockId int `db:"ProductStockId"`

	OrderProductId int     `db:"OrderProductId"`
	ProductId      int     `db:"ProductId"`
	Label          string  `db:"Label"`
	Info           string  `db:"Info"`
	Price          float64 `db:"Price"`
	TypeId         int     `db:"TypeId"`
	TypeLabel      string  `db:"TypeLabel"`

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

func BuildOrderFromTmp(input []OrderGetTmp) *Order {
	cAt, _ := time.Parse(wdb.ParseDateLayout(), input[0].CreatedAt)
	uAt, _ := time.Parse(wdb.ParseDateLayout(), input[0].UpdatedAt)
	order := Order{input[0].OrderId, input[0].TableId, input[0].TableLabel, input[0].StaffId, input[0].StaffName, input[0].MemberId, input[0].MemberName, input[0].IsPaid, input[0].TotalPrice, input[0].AmountPaid, input[0].PaymentMethodId, input[0].PaymentMethodLabel, input[0].InvoiceId, input[0].CreatedBy, input[0].UpdatedBy, cAt, uAt, []OrderProduct{}}

	var s int = 0
	var id int = input[0].OrderProductId
	l := len(input)

	for i := range input {
		if input[i].OrderProductId != id {
			order.Products = append(order.Products, BuildProductFromTmp(input[s:i]))
			s = i
			id = input[i].OrderProductId
		}
		if i == l-1 {
			order.Products = append(order.Products, BuildProductFromTmp(input[s:]))
		}
		// exRow := getOrderProduct(order.Products, input[i].OrderProductId)
		// if exRow == nil {
		// 	exRow = createProductDetails(order, input[i])
		// 	order.Products = append(order.Products, *exRow)
		// }
		// // exRow.Product.Composition = append(exRow.Product.Composition, )
		// updateProductDetails(exRow, input[i])
		// fmt.Printf("LOOP : %v\n", exRow.Product.Composition)
	}

	return &order
}

func BuildProductFromTmp(input []OrderGetTmp) OrderProduct {

	var res = OrderProduct{input[0].OrderProductId, input[0].ProductId, nil}

	var compos []product.ProductStock
	for i := range input {
		compos = append(compos, product.ProductStock{input[i].StockId, input[i].ProductStockId, input[i], input[i].StockLabel, input[i].StockRequired, input[i].StockAvailable, input[i].StockTypeId})
	}
	res.Product = &product.Product{input[0].ProductId, input[0].Label, input[0].Price, input[0].TypeId, input[0].TypeLabel, compos}

	return res
}

func parse(input []OrderGetTmp) ([]*Order, *werror.Error) {
	var res []*Order

	if len(input) == 0 {
		return res, nil
	}

	var s int = 0
	var id int = input[0].OrderId
	l := len(input)

	for i := range input {
		if input[i].OrderId != id {
			res = append(res, BuildOrderFromTmp(input[s:i]))
			s = i
			id = input[i].OrderId
		}
		if i == l-1 {
			res = append(res, BuildOrderFromTmp(input[s:]))
		}
	}

	return res, nil
}

func QueryGet() string {

	return "SELECT o.OrderId		, o.TableId, t.Label		,  IFNULL(o.StaffId, 0) AS 'StaffId'		, IFNULL(us.Name, 'No staff') AS 'StaffName'		, IFNULL(o.MemberId, 0) AS 'MemberId'		, IFNULL(ms.Name, 'No member') AS 'MemberName'		, IF(o.CreatedBy = o.MemberId, 1, 0) AS 'IsFromApp'		, o.StatusId		, os.Label AS 'StatusLabel'		, o.IsPaid		, o.TotalPrice		, o.AmountPaid		, IFNULL(o.PaymentMethodId, 0) AS 'PaymentMethodId'		, IFNULL(opm.Label, 'Not paid yet') AS 'PaymentMethodLabel'		, o.InvoiceId		, o.CreatedBy		, o.UpdatedBy		, o.CreatedAt		, o.UpdatedAt, op.OrderProductId         ,p.ProductId,        p.Label,        CONCAT (sf.Label, ' ', s.Size, ' ', cu.Label) AS Info,        pr.Price,        p.TypeId,        pt.Label AS 'TypeLabel',   ps.ProductStockId AS 'ProductStockId'   ,   s.StockId,        s.Label AS 'StockLabel',        ps.Quantity AS 'StockRequired',        s.Quantity AS 'StockAvailable',        s.Size AS 'StockSize',        s.TypeId AS 'StockTypeId',        st.Label AS 'StockTypeLabel',        s.FormatId AS 'StockFormatId',        sf.Label AS 'StockFormatLabel',        s.CategoryId,        c.Label AS 'CategoryLabel',        c.UnitId AS 'CategoryUnitId',        cu.Label AS 'CategoryUnitLabel',        c.ParentId AS 'CategoryParentId',        pc.Label AS 'CategoryParentLabel'		FROM `Order` o		INNER JOIN `Table` t ON t.TableId = o.TableId		LEFT JOIN User us ON us.UserId = o.StaffId 		LEFT JOIN User ms ON ms.UserId = o.MemberId 		INNER JOIN `OrderStatus` os ON os.OrderStatusId = o.StatusId		LEFT JOIN PaymentMethod opm ON opm.PaymentMethodId = o.PaymentMethodId        LEFT JOIN OrderProduct op ON op.OrderId = o.OrderId        LEFT JOIN Product p ON p.ProductId = op.ProductId        INNER JOIN ProductType pt ON p.TypeId = pt.ProductTypeId AND p.IsDisabled = 0        INNER JOIN Price pr ON pr.ProductId = p.ProductId AND pr.EndDate IS NULL        INNER JOIN ProductStock ps ON ps.ProductId = p.ProductId AND ps.EndDate IS NULL        INNER JOIN Stock s ON s.StockId = ps.StockId        INNER JOIN Category c ON c.CategoryId = s.CategoryId        INNER JOIN StockType st ON st.StockTypeId = s.TypeId        INNER JOIN StockFormat sf ON sf.StockFormatId = s.FormatId        INNER JOIN CategoryUnit cu ON cu.CategoryUnitId = c.UnitId        LEFT JOIN Category pc ON pc.CategoryId = c.ParentId    ORDER BY  o.CreatedAt, o.OrderId, OrderProductId, ProductId    ;"
	/*

			SELECT o.OrderId
				, o.TableId
				, t.Label
				, IFNULL(o.StaffId, 0) AS 'StaffId'
				, IFNULL(us.Name, 'No staff') AS 'StaffName'
				, IFNULL(o.MemberId, 0) AS 'MemberId'
				, IFNULL(ms.Name, 'No member') AS 'MemberName'
				, IF(o.CreatedBy = o.MemberId, 1, 0) AS 'IsFromApp'
				, o.StatusId
				, os.Label AS 'StatusLabel'
				, o.IsPaid
				, o.TotalPrice
				, o.AmountPaid
				, IFNULL(o.PaymentMethodId, 0) AS 'PaymentMethodId'
				, IFNULL(opm.Label, 'Not paid yet') AS 'PaymentMethodLabel'
				, o.InvoiceId
				, o.CreatedBy
				, o.UpdatedBy
				, o.CreatedAt
				, o.UpdatedAt

				,op.OrderProductId
		        ,p.ProductId,
		        p.Label,
		        CONCAT (sf.Label, " ", s.Size, " ", cu.Label) AS Info,
		        pr.Price,
		        p.TypeId,
		        pt.Label AS 'TypeLabel',

				ps.ProductStockId AS 'ProductStockId'

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

				FROM `Order` o
				INNER JOIN `Table` t ON t.TableId = o.TableId
				LEFT JOIN User us ON us.UserId = o.StaffId
				LEFT JOIN User ms ON ms.UserId = o.MemberId
				INNER JOIN `OrderStatus` os ON os.OrderStatusId = o.StatusId
				LEFT JOIN PaymentMethod opm ON opm.PaymentMethodId = o.PaymentMethodId
		        LEFT JOIN OrderProduct op ON op.OrderId = o.OrderId
		        LEFT JOIN Product p ON p.ProductId = op.ProductId
		        INNER JOIN ProductType pt ON p.TypeId = pt.ProductTypeId AND p.IsDisabled = 0
		        INNER JOIN Price pr ON pr.ProductId = p.ProductId AND pr.EndDate IS NULL
		        INNER JOIN ProductStock ps ON ps.ProductId = p.ProductId AND ps.EndDate IS NULL
		        INNER JOIN Stock s ON s.StockId = ps.StockId
		        INNER JOIN Category c ON c.CategoryId = s.CategoryId
		        INNER JOIN StockType st ON st.StockTypeId = s.TypeId
		        INNER JOIN StockFormat sf ON sf.StockFormatId = s.FormatId
		        INNER JOIN CategoryUnit cu ON cu.CategoryUnitId = c.UnitId
				LEFT JOIN Category pc ON pc.CategoryId = c.ParentId
				ORDER BY  o.CreatedAt, o.OrderId, OrderProductId, ProductId
				;

	*/
}

func getOrderProduct(input []OrderProduct, orderProductId int) *OrderProduct {
	for i := range input {
		if input[i].OrderProductId == orderProductId {
			return &input[i]
		}
	}
	return nil
}

func createProductDetails(order Order, data OrderGetTmp) *OrderProduct {
	return &OrderProduct{data.OrderProductId, data.ProductId, nil}
}

func updateProductDetails(orderProduct *OrderProduct, data OrderGetTmp) {
	if orderProduct.Product == nil {
		var compos []product.ProductStock
		orderProduct.Product = &product.Product{data.ProductId, data.Label, data.Price, data.TypeId, data.TypeLabel, compos}
	}
	orderProduct.Product.Composition = append(orderProduct.Product.Composition, product.ProductStock{data.StockId, data.ProductStockId, data.StockLabel, data.StockRequired, data.StockAvailable, data.StockTypeId})
	// return product.ProductStock{data.StockId, data.StockRequired, data.StockAvailable, data.StockTypeId}
}
