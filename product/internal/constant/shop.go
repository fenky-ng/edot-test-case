package constant

type ShopStatus string

const (
	ShopStatus_Active   = "ACTIVE"
	ShopStatus_Inactive = "INACTIVE"
)

func (t ShopStatus) String() string {
	return string(t)
}
