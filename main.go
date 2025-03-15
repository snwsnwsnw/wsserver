package main

import (
	"context"
	"github.com/snwsnwsnw/wsserver/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server.LoadWSAddrs()
	server.Start(context.Background())
	server.RedisSubscribe()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	server.Server.Stop()
}
