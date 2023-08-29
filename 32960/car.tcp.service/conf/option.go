package conf

var (
	Conf = &local{}
)

// local 配置中心
type local struct {
	EnvStr Env
	Dapr   Dapr
	Redis  Redis
	MySQL  MySQL
	Local  Local
}

type (
	// Env ..
	Env struct {
		GoEnv string
	}
	// Dapr dapr
	Dapr struct {
		PubsubName      string
		DaprCmdSubTopic string
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
	Local struct {
		LocalHost string
		LocalPort string
	}
)
