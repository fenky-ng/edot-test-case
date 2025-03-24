package constant

const (
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json"
	HeaderAuth      = "Authorization"
	HeaderApiKey    = "X-API-KEY"
	AuthBearer      = "Bearer"

	JwtKey    = "jwt"
	UserIdKey = "userId"
)

const (
	// user
	UserGetProfileUri = "/api/v1/users/me"

	// shop
	ShopGetMyShopUri = "/api/v1/shops/me"

	// product
	ProductGetProductsUri = "/api/v1/products"

	// warehouse
	WarehouseDeductStocksUri = "/api/v1/warehouses/stocks/deduct"
)
