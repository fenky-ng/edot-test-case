package constant

type ShopStatus string

const (
	ShopStatus_Active  = "ACTIVE"
	ShopStatus_Inctive = "INACTIVE"
)

func (t ShopStatus) String() string {
	return string(t)
}
