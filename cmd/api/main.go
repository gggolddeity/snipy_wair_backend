package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gggolddeity/snipy_wair_backend/internal/config"
	"github.com/gggolddeity/snipy_wair_backend/internal/server"
)

func main() {
	cfg := config.Load()

	srv := server.New(cfg)

	httpSrv := &http.Server{
		Addr:    cfg.Addr,
		Handler: srv.Router,
	}

	go func() {
		log.Printf("API listening on %s", cfg.Addr)
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	log.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpSrv.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown: %v", err)
	}
	if err := srv.Close(ctx); err != nil {
		log.Printf("Resources Close: %v", err)
	}
	log.Println("Bye.")
}
