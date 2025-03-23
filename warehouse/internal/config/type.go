package config

type Configuration struct {
	DatabaseDSN               string `env:"DATABASE_URL"`
	UserRestServiceAddress    string `env:"USER_REST_SERVICE_ADDRESS"`
	ShopRestServiceAddress    string `env:"SHOP_REST_SERVICE_ADDRESS"`
	ProductRestServiceAddress string `env:"PRODUCT_REST_SERVICE_ADDRESS"`
	ApiKeyExternalOrder       string `env:"API_KEY_EXT_ORDER"`
}
