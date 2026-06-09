package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//1. Create a server instance
	server := &http.Server{
		Addr:    ":8085",
		Handler: http.HandlerFunc(sampleHandler),
	}

	//2. Run the server in a background goroutine so it doesn't block main
	go func() {
		log.Println("Starting http server on :8085")
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server forced to shutdown unexpectedly: %v", err)
		}
	}()

	//3. Create a context that listens for incoming SIGINT OR SIGTERM signals
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	//block main execution hererr until a signal is received
	<-ctx.Done()
	log.Println("Shutdown signal received. Sutthing down...")

	//4. Set a maximum cleanup timeout window (e.g. 10 seconds)
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//5. Trigger the server shutdown process
	//This stops acception a new requests and waits for ongoing requests to finish
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown with errors: %v", err)
	}

	log.Println("Server exited cleanly.")
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	//simulated inflight work
	time.Sleep(4 * time.Second)
	w.Write([]byte("Task finished successfully!"))
}
