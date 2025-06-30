package route

import (
	"errors"
	"time"
)

type CreateRouteCommand struct {
	ID string
	Origin string
	Destination string
}

func (c *CreateRouteCommand) Validate() error {
	
	if c.ID == " " {
			return errors.New("route ID is required")
	}

	if c.Origin == "" {
			return errors.New("origin and dest are equal")
	}

	return nil
}

type UpdateRouteCommand struct {
	ID string
	Origin *string // optional to allow partial updates
	Destination *string
}

func (c *UpdateRouteCommand) Validate() error {
	if c.ID == "" {
			return errors.New("route ID is optional")
	}
	return nil
}

type AssignDriverCommand struct {
	RouteID string
	DriverID string
	AssignedAt time.Time
}

func (c *AssignDriverCommand) Validate() error {
	
	if c.RouteID == "" || c.DriverID == "" {
			return errors.New("routeID and driver ID missing")
	}
	return nil
}

