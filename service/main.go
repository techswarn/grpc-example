package main

import (
	"log"
	"net/http"
	"github.com/rs/cors"
	"os"
	"time"
	"fmt"
	"github.com/Educative-Content/protoapi"
	"google.golang.org/grpc"
)

func main() {
	PORT := ":8002"
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}
	log.Println("Using port number: ", PORT)
	mux := http.NewServeMux()

	mux.Handle("/api/v1/", http.HandlerFunc(getValue))
	handler := cors.Default().Handler(mux)
	s := &http.Server{
		Addr:         PORT,
		Handler:      handler,
		IdleTimeout:  100 * time.Second,
		ReadTimeout:  100 * time.Second,
		WriteTimeout: 100 * time.Second,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Println(err)
		return
	}
}

func getValue(w http.ResponseWriter,req *http.Request) {
		log.Printf("Endpoint: %s", req.URL.Path)
		
		if req.URL.Path != "/api/v1/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "Welcome to the home page!")
}