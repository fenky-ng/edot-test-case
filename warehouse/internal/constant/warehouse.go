package constant

type WarehouseStatus string

const (
	WarehouseStatus_Active  WarehouseStatus = "ACTIVE"
	WarehouseStatus_Inctive WarehouseStatus = "INACTIVE"
)

func (t WarehouseStatus) String() string {
	return string(t)
}
