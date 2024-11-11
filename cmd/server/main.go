package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/liambardong/grpc-ndbc/api/proto/v1"
	"github.com/liambardong/grpc-ndbc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Initialize listner
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// Register our service
	stationService := service.NewStationService()
	pb.RegisterStationServiceServer(s, stationService)

	// Register reflectionService on gRPC server
	// This will allow tools like grpcurl to introspect the service
	reflection.Register(s)

	//Start gRPC server
	log.Printf("Starting gRPC Server on %s", lis.Addr().String())
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()

	// Close when interupt signal is sent
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down gRPC server...")
	s.GracefulStop()
	log.Println("Server stopped")
}
