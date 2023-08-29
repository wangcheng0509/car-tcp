package main

import (
	"fmt"
	"net"

	"car.tcp.service/conf"
	daprsub "car.tcp.service/daprSub"
	"car.tcp.service/entity/enums"
	"car.tcp.service/handle"
	"car.tcp.service/repository"

	"github.com/wangcheng0509/gpkg/gredis"
)

func init() {
	conf.Init()
	enums.Init()
	repository.RegistryMySQL()
	gredis.SetupClient(conf.Conf.Redis.Addr, conf.Conf.Redis.Password)
	getLocalHost()
}

func main() {
	go daprsub.Init()
	handle.ServeStart()
}
func getLocalHost() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
				conf.Conf.Local.LocalHost = ipnet.IP.String()
			}
		}
	}
}
