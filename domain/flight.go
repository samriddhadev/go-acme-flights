package domain

import (
	"time"

	"github.com/uptrace/bun"
)

type Flight struct {
	bun.BaseModel          `bun:"table:flights,alias:fl"`
	Id                     int `bun:",pk,autoincrement"`
	SegmentId              int
	Segment                *FlightSegment `bun:"rel:belongs-to,join:segment_id=id"`
	ScheduledDepartureTime time.Time
	ScheduledArrivalTime   time.Time
	FirstClassBaseCost     float64
	EconomyClassBaseCost   float64
	NumFirstClassSeats     int
	NumEconomyClassSeats   int
	AirplaneTypeId         string
}
