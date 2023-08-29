package conf

import (
	"os"

	"github.com/BurntSushi/toml"
)

// Init 初始化
func Init() {
	b, err := os.ReadFile("conf/config.toml")
	if err != nil {
		panic(err)
	}
	Conf.load(b)
}

// 加载配置
func (c *local) load(content []byte) error {
	err := toml.Unmarshal(content, c)
	if err != nil {
		return err
	}

	if env := os.Getenv("GOENV"); env != "" {
		c.EnvStr.GoEnv = env
	}

	return nil
}
