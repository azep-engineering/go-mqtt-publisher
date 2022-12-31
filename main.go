package main

import (
	"azep/mqtt-publisher/interfaces"
	"azep/mqtt-publisher/interfaces/mqtt"
	"github.com/shkshariq/go-util/config"
)

func main() {

	config.ParseAppConfig()

	mqtt.Init()

	interfaces.Run()
}
