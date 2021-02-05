package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	contents, err := ioutil.ReadFile("problems/merge_meetings.go")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	log.Printf(string(contents))

	fs, err := ioutil.ReadDir("problems")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	for _, file := range fs {
		log.Printf(file.Name())
	}

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh,
		os.Interrupt,
		syscall.SIGABRT,
		syscall.SIGQUIT,
		syscall.SIGTERM)

	go func() {
		log.Printf("Listening on http://%s\n", addr)

		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	<-stopCh

	log.Println("Shutting down the server...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("Server failed to shutdown gracefully: %v", err)
	}
	log.Println("Server gracefully shutdown")
}
