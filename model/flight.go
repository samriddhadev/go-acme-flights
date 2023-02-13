package model

import "time"

type Flight struct {
	Id                     int64     `json:"id"`
	Name                   string    `json:"name"`
	Origin                 string    `json:"origin"`
	Destination            string    `json:"destination"`
	Miles                  int       `json:"miles"`
	ScheduledDepartureTime time.Time `json:"scheduled_departure_time"` //rfc3339
	ScheduledArrivalTime   time.Time `json:"scheduled_arrival_time"`
	FirstClassBaseCost     float64   `json:"first_class_base_price"`
	EconomyClassBaseCost   float64   `json:"economy_class_base_price"`
	NumFirstClassSeats     int       `json:"num_first_class_seats"`
	NumEconomyClassSeats   int       `json:"num_economy_class_seats"`
	AirplaneTypeId         string    `json:"airplane_type_id"`
}
