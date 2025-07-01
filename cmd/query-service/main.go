package main

import (
	"context"
	"log"
	"net"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"github.com/KingBean4903/cqrs-logistics-fleet-tracking/services/query"
	routepb "github.com/KingBean4903/cqrs-logistics-fleet-tracking/api/gen"
)

func main() {
		
	ctx := context.Background()

	// Connect to Redis
	redisClient := redis.NewClient(&redis.Options{
			Addr: "redis:6379",
			DB: 0,
	})

	if err := redis.Client.Ping(ctx).Err(); err != nil {
			log.Fatalf("Failed to connect to redis: %v", err)
	}

	log.Println("Connectd to Redis")

	// Create ViewStore
	store := quer.NewViewStore(redisClient)
	
	// Start pagination
	projection := query.NewProjection(
		[]string{"kafka:9092"},
		"route-query-service",
		store,
	)
	if err := projection.Start(ctx); err != nil {
			log.Fatalf("Failed to start projection: %v", err)
	}

	log.Println("Started event projection form kafka")

	// Start GRPC Server
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen on port :50052: %v", err)
	}

	grpcServer := grpc.NewServer()

	routepb.RegisterRouteQueryService(grpcServer, queryNewQueryGRPCServer(route))

	log.Println("Route Query Service running on port :50052")

	if err := grpcServer.Server(lis); err != nil {
		log.Fatalf("gRPC Server failed: %v", err)
	}

}
