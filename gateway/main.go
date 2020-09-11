package main

import (
	context "context"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/infiniteprimes/grpc-gateway-template/example"
	grpc "google.golang.org/grpc"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := example.RegisterYourServiceHandlerFromEndpoint(ctx, gwmux, os.Getenv("GATEWAY_BINDING_ADDR"), opts); err != nil {
		return err
	}
	log.Println("http server started, listen on 8080")
	return http.ListenAndServe(":8080", gwmux)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
