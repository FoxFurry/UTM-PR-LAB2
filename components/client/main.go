package main

import (
	"context"
	"github.com/foxfurry/university/PRLab2/components/client/application"
	"github.com/foxfurry/university/PRLab2/components/client/infrastructure/config"
	"os"
	"os/signal"
	"syscall"
)

func init(){
	config.LoadConfig("cfg.yaml")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGSEGV)

	mainApp := application.New()
	go mainApp.Start(ctx)

	<-sigs
	cancel()
}
