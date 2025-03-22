package constant

type ShopStatus string

const (
	ShopStatus_Active  ShopStatus = "ACTIVE"
	ShopStatus_Inctive ShopStatus = "INACTIVE"
)

func (t ShopStatus) String() string {
	return string(t)
}
