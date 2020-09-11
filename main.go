package main

import (
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/infiniteprimes/grpc-gateway-template/controller"
	"github.com/infiniteprimes/grpc-gateway-template/example"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	PORT = ":9192"
)

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("faled to listen: %v", err)
	}

	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()

	grpc_zap.ReplaceGrpcLoggerV2(zapLogger)
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(zapLogger),
		)))

	example.RegisterYourServiceServer(s, controller.NewServer())
	log.Println("grpc server started, listen on 9192")
	s.Serve(lis)
}
