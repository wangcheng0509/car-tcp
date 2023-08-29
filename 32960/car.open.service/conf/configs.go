package conf

import (
	"os"

	"github.com/BurntSushi/toml"
)

var (
	Conf = &local{}
)

// local 配置中心
type local struct {
	Env   Env
	HTTP  HTTP
	Dapr  Dapr
	Redis Redis
	MySQL MySQL
	CK    ClickHouse
}

type ()

// 加载配置
func (c *local) load(content []byte) error {
	err := toml.Unmarshal(content, c)
	if err != nil {
		return err
	}

	if env := os.Getenv("GOENV"); env != "" {
		c.Env.GoEnv = env
	}
	return nil
}
