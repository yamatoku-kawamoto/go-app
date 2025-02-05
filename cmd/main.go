package main

import (
	"context"
	"goapp/internal/repository"
	"goapp/internal/web"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	engine *web.Engine

	configuration Configuration
)

func main() {
	if err := initialize(); err != nil {
		panic(err)
	}
	if err := run(); err != nil {
		panic(err)
	}
	if err := shutdown(); err != nil {
		panic(err)
	}
}

type Configuration struct {
	Repository *repository.Config
}

func (c Configuration) Validate() error {
	return nil
}

func initialize() error {
	if err := initWeb(); err != nil {
		return err
	}
	if err := configuration.Validate(); err != nil {
		return err
	}
	if err := repository.Initialize(configuration.Repository); err != nil {
		return err
	}
	return nil
}

func run() (err error) {
	start := func() {
		err = engine.Run()
		if err != nil {
			return
		}
	}

	go start()

	return
}

func shutdown() (err error) {
	const timeout = 15 * time.Second

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	complete := make(chan error)
	defer close(complete)

	shutdown := func(ctx context.Context) {
		err = engine.Shutdown(ctx)
		if err != nil {
			complete <- err
			return
		}
		// other shutdown tasks here
		// example:
		// if err := db.Close(); err != nil {
		// 	complete <- err
		// 	return
		// }

		complete <- nil
	}

	go shutdown(ctx)

	select {
	case err := <-complete:
		if err != nil {
			return err
		}
		// graceful shutdown
		return nil
	case <-ctx.Done():
		// timeout
		return ctx.Err()
	}
}
