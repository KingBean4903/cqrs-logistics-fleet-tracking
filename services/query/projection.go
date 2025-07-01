package query

import (
	"context"
	"log"
	"encoding/json"
	"github.com/segmentio/kafka.go"
)

type Projection struct {
	Brokers []string
	GroupID string
	ViewStore *ViewStore
}

// Constructor
func NewProjection(brokers []string, groupID string,
				strore *ViewStore) *Projection {

					return &Projection{
							Brokers: brokers,
							GroupID: groupID,
							ViewStore: store
					}

}

func (p *Projection) consume(ctx context.Context, topic string, handler func(context.Context, []byte) error) { 
	r := kafka.NewReader(kafka.ReaderConfig{
				Brokers: p.Brokers,
				GroupID: p.GroupID,
				Topic: topic,
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
				log.Printf("Kafka read error from topic %s: %v", topic, err)
				continue
		}

		if err := handler(ctx, msg.Value); err != nil {
						log.Printf("Error handling event on topic %s : %v", topic, err)
		}
	}
}

func (p *Projection) Start(ctx context.Context) error {
	
	go p.consume(ctx, "fleet.route.created", p.applyRouteCreated)
	go p.consume(ctx, "fleet.route.updated", p.applyRouteUpdated)
	go p.consume(ctx, "fleet.driver.assigned", p.applyDriverAssigned)
	return nil
}

func (p *Projection) applyRouteCreated(ctx context.Context, data []byte) error {
	
	var e RouteCreatedEvent
	if err := json.Marshal(data, &e); err != nil { 
			return err
	}
	
	view := &RouteView{
			ID: e.ID,
			Origin: e.Origin,
			Destination: e.Destination,
			Status: "created",
			LastUpdated: e.CreatedAt.Unix(),
	}

	return p.ViewStore.Set(ctx, view)

}

