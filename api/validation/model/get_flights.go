package model

import (
	"time"
)

type GetFlightsQuery struct {
	FromAirport string `query:"from_airport" binding:"required"`
	ToAirport   string
	FromDate    time.Time
	ToDate      time.Time
}

