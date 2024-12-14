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
	"errors"
	"context"
)


var con *grpc.ClientConn

func init() {
	port := "grpc-server:3000"
    err := errors.New("err")
	con, err = grpc.NewClient(port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial:", err)
		return
	}
}

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
// 		port := "0.0.0.0:8080"
//     	err := errors.New("err")
// 		con, err := grpc.NewClient(port, grpc.WithInsecure())
// //		con, err := grpc.Dial("0.0.0.0:8080", grpc.WithInsecure())
// 		if err != nil {
// 			fmt.Println("Dial:", err)
// 			return
// 	    }
		log.Printf("Endpoint: %s", req.URL.Path)
		
		if req.URL.Path != "/api/v1/" {
			http.NotFound(w, req)
			return
		}
		client := protoapi.NewRandomClient(con)
		log.Printf("%#v", client)
		r, err := AskingDateTime(context.Background(), client)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Server Date and Time:", r.Value)

		fmt.Fprintf(w, r.Value)
}

func AskingDateTime(ctx context.Context, m protoapi.RandomClient) (*protoapi.DateTime, error) {
	log.Println("Here")
	request := &protoapi.RequestDateTime{
		Value: "Please send me the date and time",
	}

	return m.GetDate(ctx, request)
}