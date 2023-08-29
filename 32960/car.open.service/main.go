package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"car.open.service/api"
	"car.open.service/conf"
	"car.open.service/core"
	"car.open.service/repo"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/gin-gonic/gin"
)

var httpPort string

func init() {
	conf.Init()
	repo.Init()

	flag.StringVar(&httpPort, "httpPort", "80", "-httpPort=")
}

func main() {
	flag.Parse()

	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go HTTPServer()

EXIT:
	for {
		sig := <-sc
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	// cleanFunc()
	time.Sleep(time.Second)
	os.Exit(state)
}

// HTTPServer ..
func HTTPServer() {
	e := gin.New()
	g := e.Group("car")
	{
		g.POST("/cmd", new(api.Cmd).Send)
	}

	http.ListenAndServe(":"+httpPort, e)
}

// DaprServer ..
func DaprServer() {
	if os.Getenv("GOENV") != "local" {
		var err error
		core.DaprClient, err = dapr.NewClient()
		if err != nil {
			panic(err)
		}
		defer core.DaprClient.Close()
	}
}
