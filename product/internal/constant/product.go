package constant

type ProductStatus string

const (
	ProductStatus_Active   ProductStatus = "ACTIVE"
	ProductStatus_Inactive ProductStatus = "INACTIVE"
)

func (t ProductStatus) String() string {
	return string(t)
}
