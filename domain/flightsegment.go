package domain

import "github.com/uptrace/bun"

type FlightSegment struct {
	bun.BaseModel `bun:"table:flights_segments,alias:flsg"`
	Id            int64 `bun:",pk,autoincrement,alias:segment_id"`
	Name          string
	Origin        string
	Destination   string
	Miles         int
}
