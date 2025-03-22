package constant

const (
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json"
	HeaderAuth      = "Authorization"
	AuthBearer      = "Bearer"

	JwtKey    = "jwt"
	UserIdKey = "userId"
)

const (
	// user
	UserGetProfileUri = "/api/v1/users/me"

	// shop
	ShopGetMyShopUri = "/api/v1/shops/me"
	ShopGetShopsUri  = "/api/v1/shops"
)
