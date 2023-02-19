package model

import "time"

type FlightFilter struct {
	Origin      string    `form:"origin"`
	Destination string    `form:"destination"`
	FromDate    time.Time `form:"from_date"`
	ToDate      time.Time `form:"to_date"`
}
