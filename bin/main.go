package main

import (
	"os/signal"
	"os"
	"syscall"
	"dora.org/bara/gateway/protocol"
	"github.com/istio/istio/pkg/log"
)

func main() {
	var sign = make(chan os.Signal, 1)
	protocol.InitConfig()
	signal.Notify(sign, os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	<-sign
	log.Info("Shutdown")
}
