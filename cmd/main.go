package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"otus/cmd/app"
)

func main() {
	webApp := app.Initialize()
	go func() {
		if err := webApp.Run(); err != nil {
			if err == http.ErrServerClosed {
			} else {
				log.Fatalf("error occured while running http server: %s", err.Error())
			}
		}
	}()
	closeFn := func(closed func()) {
		if err := webApp.Shutdown(context.Background()); err != nil {
			os.Exit(-1)
		}
		closed()
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	closeFn(func() {
		log.Printf("close by signal %s", sig)
	})
}
