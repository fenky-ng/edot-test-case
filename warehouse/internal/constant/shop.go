package constant

type WarehouseStatus string

const (
	WarehouseStatus_Active  = "ACTIVE"
	WarehouseStatus_Inctive = "INACTIVE"
)

func (t WarehouseStatus) String() string {
	return string(t)
}
