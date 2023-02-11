package model

import "time"

type Flight struct {
	Id                     int
	Name                   string
	Origin                 string
	Destination            string
	Miles                  int
	ScheduledDepartureTime time.Time
	ScheduledArrivalTime   time.Time
	FirstClassBaseCost     float64
	EconomyClassBaseCost   float64
	NumFirstClassSeats     int
	NumEconomyClassSeats   int
	AirplaneTypeId         string
}
