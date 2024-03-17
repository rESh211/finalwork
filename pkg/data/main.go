package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rESh211/finalwork/blob/main/finalwork/pkg/config"

	"github.com/astrviktor/skillbox_diploma/pkg/server"
)

func main() {
	config.GlobalConfig = config.NewConfig("config.yaml")
	config.GlobalConfig = config.ForHerokuConfig(config.GlobalConfig)

	go server.StartServer()

	exit := make(chan os.Signal, 0)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	<-exit
}
