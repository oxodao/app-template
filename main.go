package main

import (
	"fmt"
	"github.com/oxodao/app-template/log"
	"github.com/oxodao/app-template/services"
)

const (
	appname = "AppTemplate"
	version = "0.1"
)

func main() {
	log.CreateLogger(log.LevelDebug)
	log.InfoAlways(fmt.Sprintf("%v [v.%v]", appname, version))

	prv := services.NewProvider()

}
