package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

    srv := &http.Server {
        Addr: "0.0.0.0:8080",
        WriteTimeout: time.Second*15,
        ReadTimeout: time.Second*15,
        IdleTimeout: time.Second*15,
        Handler: router,
    }

    go func() {
        if err := srv.ListenAndServe(); err != nil {
            log.Println(err)
        }
    }()

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)

    <-c

    ctx,cancel := context.WithTimeout(context.Background(),time.Second*15 )
    defer cancel()

    srv.Shutdown(ctx)
    log.Println("shutting down")
    os.Exit(0)
}
