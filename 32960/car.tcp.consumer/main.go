package main

import (
	"log"
	"net/http"
	"os"

	"car.tcp.consumer/app"
	"car.tcp.consumer/conf"
	"car.tcp.consumer/repo"

	"car.tcp.consumer/subtask"
	dapr "github.com/dapr/go-sdk/client"
	daprd "github.com/dapr/go-sdk/service/http"
)

func init() {
	conf.Init()
	repo.Init()
}

func main() {
	if os.Getenv("GOENV") == "local" {
		Test()
	}
	s := daprd.NewService(":8080")

	subtask.VehicleLogin(s)
	subtask.ResendInfo(s)
	subtask.RealtimeInfo(s)
	subtask.VehicleLogout(s)
	subtask.Vehicle(s)
	subtask.Location(s)
	subtask.Fuelcell(s)
	subtask.Extreme(s)
	subtask.EngineTopic(s)
	subtask.DrivemotorTopic(s)
	subtask.ChargeableTempTopic(s)
	subtask.ChargeableVoltageTopic(s)
	subtask.Command(s)

	var err error
	app.DaprClient, err = dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer app.DaprClient.Close()

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}

}

func Test() {
	// subtasktest.VehicleLoginTest()
	// subtasktest.InfoTest()
	// subtasktest.VehicleTest()
	// subtasktest.FuelcellTest()
	// subtasktest.ExtremeTest()
	// subtasktest.EngineTest()
	// subtasktest.DrivemotorTest()
	// subtasktest.ChargeableVoltageTest()
	// subtasktest.ChargeableTempTest()
	// subtasktest.VlogoutTest()

}
