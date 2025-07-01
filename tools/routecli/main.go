package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"
	"google.golang.org/grpc"
	routepb ""
)

func main() {
	routeID := flag.String("id", "", "Route ID to query")
	addr    := flag.String("addr", "localhost:50052")
	flag.Parse()

	if *routeID == "" {
			fmt.Println("Usage: routecli --id=<ROUTE_ID>")
			os.Exit(1)
	}

	// Connect to gRPC
	conn, err := gprc.Dial(*addr, grpc.WithBlock(), grpc.WithTimeout(5 * time.Second))
	if err != nil {
			log.Fatalf("Failed to connect to query-service: %v", err)
	}

  defer conn.Close()
	client := routepb.NewRouteQueryServiceClient(conn)

	// Make request
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	res, err := client.GetRoute(ctx, &routepb.GetRouteRequest{Id: *routeID})
	if err != nil {
			log.Fatalf("Failed to fetch route: %v", err)
	}

	// Output result
	fmt.Println("Route view")
	fmt.Println("-------------------------------")
	fmt.Println("ID:             %s\n", res.GetId())
	fmt.Println("Origin          %s\n", res.GetOrigin())
	fmt.Println("Destination          %s\n", res.GetDestination())
	fmt.Println("AssignedDriver          %s\n", res.GetAssignedDriver())
	fmt.Println("Status          %s\n", res.GetStatus())
	fmt.Println("Last Updated:          %s\n", time.Unix(res.GetLastUpdated(), 0).Format(time.RFC1123))

}
