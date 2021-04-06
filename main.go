package main

import (
	"context"
	"fmt"
	"github.com/eshaanmangal/Go-Project-Structure/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	Port string
}

func (app *App) Start() {

	l := log.New(os.Stdout, " stack-api ", log.LstdFlags)
	sm := mux.NewRouter()

	sm.HandleFunc("/push/{element}", services.HandlePushElement)
	s := &http.Server{
		Addr:         app.Port,
		Handler:      sm,
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			fmt.Printf("Sorry could not connect to server %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 10)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received a Graceful Shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}

func main() {
	server := App{
		Port: ":8080",
	}
	server.Start()
}