package config

type Configuration struct {
	DatabaseDSN                 string `env:"DATABASE_URL"`
	UserRestServiceAddress      string `env:"USER_REST_SERVICE_ADDRESS"`
	ShopRestServiceAddress      string `env:"SHOP_REST_SERVICE_ADDRESS"`
	ProductRestServiceAddress   string `env:"PRODUCT_REST_SERVICE_ADDRESS"`
	WarehouseRestServiceAddress string `env:"WAREHOUSE_REST_SERVICE_ADDRESS"`
	WarehouseRestServiceApiKey  string `env:"WAREHOUSE_REST_SERVICE_API_KEY"`
}
