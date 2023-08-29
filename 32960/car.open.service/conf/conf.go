package conf

import "os"

// Init 初始化
func Init() {
	b, err := os.ReadFile("conf/config.toml")
	if err != nil {
		panic(err)
	}
	Conf.load(b)
}
