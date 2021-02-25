package main

import (
	"flag"
	"fmt"
	"github.com/oxodao/app-template/config"
	"github.com/oxodao/app-template/fixtures"
	"github.com/oxodao/app-template/log"
	"github.com/oxodao/app-template/services"
	"os"
	"os/signal"
	"syscall"
)

const (
	appname = "AppTemplate"
	version = "0.1"
)

func main() {
	log.CreateLogger(log.LevelDebug)
	log.InfoAlways(fmt.Sprintf("%v [v.%v]", appname, version))

	if err := run(); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	prv, err := services.NewProvider(cfg)
	if err != nil {
		return err
	}

	shouldInitializeDb := flag.Bool("init-db", false, "Should initialize the database")
	forceDb := flag.Bool("force", false, "Force the DB initialization")
	flag.Parse()

	if *shouldInitializeDb {
		fixtures.InitializeDB(prv, *forceDb)
	}

	go RunServer(prv)

	log.InfoAlways(appname + " is running.\nCTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	return nil
}
