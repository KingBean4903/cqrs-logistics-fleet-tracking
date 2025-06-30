package route

import "time"

type RouteCreatedEvent struct {
		ID string
		Origin string
		Destination string
		CreatedAt time.Time
}

type RouteUpdatedEvent struct {
		ID string
		Origin *string
		Destination *string
		UpdatedAt time.Time
}

type DriverAssignedEvent struct {
		RouteID string
		DriverID string
		AssignedAt time.Time
}
