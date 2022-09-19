package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"parking-service/internal/router"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type config struct {
	port   int
	env    string
	dbHost string
	dbUser string
	dbPass string
	dbName string
	dbPort int
}

func main() {
	cfg := config{
		port:   8080,
		env:    "dev",
		dbHost: "localhost",
		dbPort: 5432,
		dbUser: "postgres",
		dbPass: "example",
		dbName: "parkingservice",
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.dbHost, cfg.dbPort, cfg.dbUser, cfg.dbPass, cfg.dbName,
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, "postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	router := router.New(db)

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", cfg.port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 15,
		Handler:      router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
