package route

import (
	"context"
	"encoding/json"
	"github.com/segementio/kafka-go"
)

type KafkaPublisher struct {
	routeCreatedWriter *kafka.Writer
	routeUpdatedWriter *kafka.Writer
	driverAssignedWriter *kafka.Writer
}

// Constructor
func NewKafkaPublisher(brokers []string) *KafkaPublisher {
		
	return &KafkaPublisher {
			routeCreatedWriter: &kafka.Writer{
					Addr: kafka.TCP(brokers...),
					Topic: "fleet.route.created",
					Balancer: &kafka.LeastBytes{},
			},
			routeUpdatedWriter: &kafka.Writer{
					Addr: kafka.TCP(brokers...),
					Topic: "fleet.route.updated",
					Balancer: &kafka.LeastBytes{},
			},
			driverAssignedWriter: &kafka.Writer{
					Addr: kafka.TCP(brokers...),
					Topic: "fleet.driver.assigned",
					Balancer: &kafka.LeastBytes{},
			}
	}
}

func (p *KafkaPublisher) PublishRoute(ctx context.Context, event RouteCreatedEvent) error {
	return p.publish(ctx, p.routeCreatedWriter, event.ID, event)
}


func (p *KafkaPublisher) publish(ctx context.Context, wrier *kafka.Writer, key string,event any) error {
	value, err := json.Marshal(event)

	if err != nil {
		return err
	}

	msg := kafka.Message{
			Key: []byte(key),
			Value: value
	}
	return writer.WriteMessage(ctx, msg)
}
