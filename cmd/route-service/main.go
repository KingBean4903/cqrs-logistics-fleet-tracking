package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	routepb "gen"
	"route"
)

func main() {
	
	ctx := context.Background()

	kafkaBrokers := []string{"kafka:9092"}
	eventPublisher := route.NewKafkaPublisher(kafkaBrokers)

	// Create Command handler
	commandHandler := route.NewCommandHandler(routeRepo, eventPublisher)

	// Start gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil { 
			log.Fatalf("Failed to listen: %v", err) 
	}

	grpcServer := grpcNewServer()
	routepb.RegisterRouteServiceServer(grpcServer, route.NewGRPCServer(commandHandler))

	log.Println("RouteService in running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("gRPC server failed: %v", err)
	}
}
