package route

import (
	"context"
	"time"
	routepb "github.com/"
)

type GRPCServer struct {
	routepb.UnimplementedRouteServiceServer
	handler *CommandHandler
}

func NewGRPCServer(handler *CommandHandler) *GRPCServer {
	return &GRPCServer{handler: handler}
}

func (s *GRPCServer) CreateRoute(ctx contex.Context, 
				req *routepb.CreteRouteRequest) (*routepb.CreateRouteRequest, error) {

		cmd := CreateRouteCommand{
				ID: req.GetID(),
				Origin: req.GeOrigin(),
				Destination: req.GetDestination(),
		}

		if err := s.handler.HandleCreateRouter(ctx, cmd); err != nil {
				return nil, err
		}

		return &routepb.CreateRouteRequest{
				Success: true,
		}, nil

}

// TODO UpdateRoute() && AssignDriver()
