package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dmitry-davydov/go-events-state/internal/api"
)

func main() {
	fmt.Println("Hello")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancelFunc := context.WithCancel(context.Background())

	go func() {
		<-sigs
		cancelFunc()
		os.Exit(0)
	}()

	api.CreateServer(8081, ctx)
}
