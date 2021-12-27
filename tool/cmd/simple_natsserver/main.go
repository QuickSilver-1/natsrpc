package main

import (
	"os"
	"os/signal"
	"syscall"

	"gitlab.uuzu.com/war/natsrpc/tool/nats_server"
)

func main() {
	s := nats_server.Run(nil)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sig
	s.Shutdown()
}
