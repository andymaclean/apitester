package test_server

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Run() {
	address := flag.String("address", "localhost:6677", "Service address")

	flag.Parse()

	fmt.Printf("TEST SERVER listening on %s\n", *address)

	server := &http.Server{
		Addr: *address,
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from %s\n", *address)
	})

	sig := make(chan bool)

	http.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Goodbye from %s\n", *address)
		sig <- true
	})

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	<-sig

	time.Sleep(1 * time.Second)
}
