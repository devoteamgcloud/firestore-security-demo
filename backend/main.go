package main

import (
	"os"
	"os/signal"

	"github.com/kaan-devoteam/firestore-security-demo/api/v1/server"
)

func main() {
	s := server.New()
	stopServerIfInterruptIsCaught(s)
	s.Run()
}

func stopServerIfInterruptIsCaught(server *server.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			server.Shutdown()
		}
	}()
}
