package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/dmitry-davydov/go-events-state/internal/api"
	"github.com/dmitry-davydov/go-events-state/internal/models"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancelFunc := context.WithCancel(context.Background())

	port := 0
	portString := os.Getenv("APP_SERVER_PORT")
	if len(portString) > 0 {
		if parsedPort, err := strconv.Atoi(portString); err == nil {
			port = parsedPort
		}
	}

	if len(portString) == 0 {
		if parsedPort, err := pickUnusedPort(); err == nil {
			port = parsedPort
		}
	}

	if port == 0 {
		panic("Server Port not provided")
	}

	interruptedLifeTime := 0
	interruptedLifeTimeString := os.Getenv("APP_INTERRUPTED_LIFETIME")
	if len(interruptedLifeTimeString) > 0 {
		if parsed, err := strconv.Atoi(interruptedLifeTimeString); err == nil {
			interruptedLifeTime = parsed
		}
	}

	if interruptedLifeTime == 0 {
		panic("Interrupted State lifetime not provided")
	}


	cfg := models.Configuration{
		HttpServerPort: port,
		InterruptedValidPeriod: interruptedLifeTime,
		DieAppContext: ctx,
	}

	_, cfg.LogHttp = os.LookupEnv("APP_DEBUG")

	go func() {
		<-sigs
		cancelFunc()
	}()

	api.CreateServer(cfg)
}

func pickUnusedPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	port := l.Addr().(*net.TCPAddr).Port
	if err := l.Close(); err != nil {
		return 0, err
	}
	return port, nil
}
