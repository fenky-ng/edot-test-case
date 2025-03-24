package constant

type ShopStatus string

const (
	ShopStatus_Active   ShopStatus = "ACTIVE"
	ShopStatus_Inactive ShopStatus = "INACTIVE"
)

func (t ShopStatus) String() string {
	return string(t)
}
