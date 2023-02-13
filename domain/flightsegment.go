package domain

import "github.com/uptrace/bun"

type FlightSegment struct {
	bun.BaseModel `bun:"table:flights_segments,alias:flsg"`
	Id            int `bun:",pk,autoincrement"`
	Name          string
	Origin        string
	Destination   string
	Miles         int
}
