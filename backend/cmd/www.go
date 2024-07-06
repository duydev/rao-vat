package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"duydev.io.vn/rao-vat/infra/http"
)

func main() {
	http.StartServer()

	go gracefulShutdown()

	// keep main goroutine run forever
	forever := make(chan int)
	<-forever
}

func gracefulShutdown() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-s
		fmt.Println("Shutdown gracefully...")

		//

		os.Exit(0)
	}()
}
