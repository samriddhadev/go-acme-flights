package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/samriddhadev/go-acme-flights/config"
	"github.com/samriddhadev/go-acme-flights/domain"
	"github.com/uptrace/bun"
)

func NewAcmeFlightRepository() AcmeFlightRepository {
	return AcmeFlightRepository{}
}

type AcmeFlightRepository struct {
	BaseRepository
}

func (repository *AcmeFlightRepository) FindAll(cfg *config.Config) (*[]domain.Flight, error) {
	ctx := context.Background()
	db := repository.GetDB(cfg)
	flights := []domain.Flight{}
	err := db.NewSelect().Model(&flights).Relation("Segment").Scan(ctx)
	return &flights, err
}

func (repository *AcmeFlightRepository) InsertOne(cfg *config.Config, flight *domain.Flight) error {
	ctx := context.Background()
	db := repository.GetDB(cfg)
	err := db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		var segment domain.FlightSegment = *flight.Segment
		segmentResult, err := tx.NewInsert().Model(&segment).Exec(ctx)
		if err != nil {
			log.Println(err)
			return err
		}
		if _, err = segmentResult.RowsAffected(); err != nil {
			log.Println(err)
			return err
		}
		flight.SegmentID = segment.Id
		flight.Segment.Id = segment.Id
		_, err = tx.NewInsert().Model(flight).Exec(ctx)
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	})
	return err
}

func (repository *AcmeFlightRepository) FindOne(cfg *config.Config, id int) (*domain.Flight, error) {
	ctx := context.Background()
	db := repository.GetDB(cfg)
	flight := domain.Flight{}
	err := db.NewSelect().Model(&flight).Where("id = ?", id).Scan(ctx)
	return &flight, err
}

func (repository *AcmeFlightRepository) UpdateOne(cfg *config.Config, id int, flight *domain.Flight) (*domain.Flight, error) {
	ctx := context.Background()
	db := repository.GetDB(cfg)
	_, err := db.NewUpdate().Model(flight).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return &domain.Flight{}, err
	}
	return repository.FindOne(cfg, id)
}

func (repository *AcmeFlightRepository) DeleteOne(cfg *config.Config, id int) error {
	ctx := context.Background()
	db := repository.GetDB(cfg)
	flight, err := repository.FindOne(cfg, id)
	if err != nil {
		return err
	}
	_, err = db.NewDelete().Model(&flight).Exec(ctx)
	return err
}
