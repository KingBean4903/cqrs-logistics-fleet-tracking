package main

import "time"

type Route struct {
	ID string
	Origin string
	Destination string
	AssignedDriverID string
	Status RouteStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RouteStatus string

const (
	RouteStatusCreated    RouteStatus = "Created"
	RouteStatusCompleted  RouteStatus = "Completed"
	RouteStatusInProgress RouteStatus = "InProgress"
	RouteStatusCancelled  RouteStatus = "Cancelled"
)
