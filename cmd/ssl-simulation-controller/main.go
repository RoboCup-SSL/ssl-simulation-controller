package main

import (
	"flag"
	"github.com/RoboCup-SSL/ssl-simulation-controller/internal/simctl"
	"os"
	"os/signal"
	"syscall"
)

var visionAddress = flag.String("visionAddress", "224.5.23.2:10006", "The address (ip+port) from which vision packages are received")
var trackerAddress = flag.String("trackerAddress", "224.5.23.2:10010", "The address (ip+port) from which tracker packages are received")
var refereeAddress = flag.String("refereeAddress", "224.5.23.1:10003", "The address (ip+port) from which referee packages are received")
var simControlPort = flag.String("simControlPort", "10300", "The port to which simulation control packets are send")

func main() {
	flag.Parse()
	ctl := simctl.NewSimulationController(*visionAddress, *refereeAddress, *trackerAddress, *simControlPort)
	ctl.Start()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	ctl.Stop()
}
