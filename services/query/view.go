package query

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)


type RouteView struct {
	ID string `json:"id"`
	Origin string `json:"origin"`
	Destination string `json:"destination"`
	AssignedDriver string `json:"assigned_driver,omitempty"`
  Status string `json:"status"`
	LastUpdated int64 `json:"last_updated"`
}

type ViewStore struct {
	redis *redis.Client
}

func (v *ViewStore) NewViewStore(client *redis.Client) *ViewStore {
		return &ViewStore{client: client]
}

func (v *ViewStore) Set(ctx context.Context, view *RouteView) error {
		key := v.routeKey(view.ID)

		data, err := json.Marshal(view)
		if err != nil {
				return err
		}

		return v.client.Set(ctx, key, data, 0).Err()
}

func (v *ViewStore) Get(ctx context.Context, routeID string) (*RouteView, error) {

	key := v.routeKey(routeID)

	raw, err := v.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var view RouteView
	if err := json.Unmarshal([]byte(raw), &view); err != nil {
			return nil, err
	}

	return &view, nil
}

func (v *ViewStore) routekey(routeID string) string {
			return fmt.Sprintf("route:view:%s", routeID)
}
