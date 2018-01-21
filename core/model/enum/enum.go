package enum

import (
	"strings"

	"github.com/Dlacreme/httpd/back/wdb"
	"github.com/Dlacreme/httpd/back/wenum"
)

var (
	UserType   map[int]string
	UserStatus map[int]string

	StockType    map[int]string
	CategoryUnit map[int]string
	Category     map[int]string
	StockFormat  map[int]string

	ProductType map[int]string

	OrderStatus   map[int]string
	PaymentMethod map[int]string

	Table map[int]string
)

func Init(db wdb.Connection) {
	UserType, _ = wenum.BuildEnum(db, "UserType", "UserTypeId", "Label")
	UserStatus, _ = wenum.BuildEnum(db, "UserStatus", "UserStatusId", "Label")

	StockType, _ = wenum.BuildEnum(db, "StockType", "StockTypeId", "Label")
	CategoryUnit, _ = wenum.BuildEnum(db, "CategoryUnit", "CategoryUnitId", "Label")
	Category, _ = wenum.BuildEnum(db, "Category", "CategoryId", "Label")
	StockFormat, _ = wenum.BuildEnum(db, "StockFormat", "StockFormatId", "Label")

	ProductType, _ = wenum.BuildEnum(db, "ProductType", "ProductTypeId", "Label")

	OrderStatus, _ = wenum.BuildEnum(db, "OrderStatus", "OrderStatusId", "Label")
	PaymentMethod, _ = wenum.BuildEnum(db, "PaymentMethod", "PaymentMethodId", "Label")

	Table, _ = wenum.BuildEnum(db, "`Table`", "TableId", "Label")
}

func GetKey(enum map[int]string, key int) (int, string) {
	return key, enum[key]
}

func GetValue(enum map[int]string, value string) (int, string) {
	for i := range enum {
		if strings.ToLower(enum[i]) == strings.ToLower(value) {
			return i, enum[i]
		}
	}
	return 0, "undefined"
}
