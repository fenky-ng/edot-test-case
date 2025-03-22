package constant

type ShopWarehouseStatus string

const (
	ShopWarehouseStatus_Active   = "ACTIVE"
	ShopWarehouseStatus_Inactive = "INACTIVE"
)

func (t ShopWarehouseStatus) String() string {
	return string(t)
}
