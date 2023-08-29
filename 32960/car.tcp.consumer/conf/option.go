package conf

type (
	// Env ..
	Env struct {
		GoEnv string
	}
	// Redis redis配置
	Redis struct {
		Addr     string
		Password string
		DB       int
	}
	// MySQL 配置
	MySQL struct {
		Type     string
		User     string
		Password string
		Host     string
		DBName   string
		Debug    string
	}
	// Dapr dapr
	Dapr struct {
		PubsubName string
	}
	// ClickHouse db
	ClickHouse struct {
		User     string
		Password string
		Host     string
		PortTCP  string
		DBName   string
		DeBug    string
	}
)
