package query

import (
	"context"
	routepb "gen"
)

type QueryGRPCServer struct {
	routepb.UnimplementedRouteQueryServiceServer
	store *ViewStore
}

func NewQueryGRPCServer(store *ViewStore) *QueryGRPCServer {
	return &QueryGRPCServer{store: store}
}

func (s *QueryGRPCServer) GetRoute(ctx context.Context, req *routepb.GetRouteRequest) (*routepb.GetRouteRequest, error) {

	view, err := s.store.Get(ctx, req.GetId())

	if err != nil {
			return nil, err
	}

	return &routepb.GetRouteRequest{
			Id: view.ID,
			Origin: view.Origin,
			Destination: view.Destination,
			AssignedDriver: view.AssignedDriver,
			Status: view.Status,
			LastUpdated: view.LastUpdated,
	}, nil

}
