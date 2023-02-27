package repository

import (
	"context"
	"database/sql"

	"github.com/samriddhadev/go-acme-flights/core/config"
	"github.com/samriddhadev/go-acme-flights/core/logger"
	"github.com/samriddhadev/go-acme-flights/domain"
	"github.com/samriddhadev/go-acme-flights/model"
	"github.com/uptrace/bun"
)

func NewAcmeFlightRepository(logger *logger.AcmeLogger) AcmeFlightRepository {
	return AcmeFlightRepository{
		logger: logger,
	}
}

type AcmeFlightRepository struct {
	BaseRepository
	logger *logger.AcmeLogger
}

func (repository *AcmeFlightRepository) FindAll(cfg *config.Config, filter *model.FlightFilter) (*[]domain.Flight, error) {
	ctx := context.Background()
	db := repository.GetDB(cfg)
	defer db.Close()
	flights := []domain.Flight{}
	query := db.NewSelect().Model(&flights).Relation("Segment")
	if filter != (&model.FlightFilter{}) {
		if len(filter.Origin) > 0 {
			query = query.Where("segment.origin = ?", filter.Origin)
		}
		if len(filter.Destination) > 0 {
			query = query.Where("segment.destination = ?", filter.Destination)
		}
		if !filter.FromDate.IsZero() {
			query = query.Where("scheduled_departure_time >= ?", filter.FromDate)
		}
		if !filter.ToDate.IsZero() {
			query = query.Where("scheduled_arrival_time >= ?", filter.ToDate)
		}
	}
	err := query.Scan(ctx)
	return &flights, err
}

func (repository *AcmeFlightRepository) InsertOne(cfg *config.Config, flight *domain.Flight) error {
	ctx := context.Background()
	db := repository.GetDB(cfg)
	defer db.Close()
	err := db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		var segment domain.FlightSegment = *flight.Segment
		segmentResult, err := tx.NewInsert().Model(&segment).Exec(ctx)
		if err != nil {
			repository.logger.Error(err.Error())
			return err
		}
		if _, err = segmentResult.RowsAffected(); err != nil {
			repository.logger.Error(err.Error())
			return err
		}
		flight.SegmentID = segment.Id
		flight.Segment.Id = segment.Id
		_, err = tx.NewInsert().Model(flight).Exec(ctx)
		if err != nil {
			repository.logger.Error(err.Error())
			return err
		}
		return nil
	})
	return err
}

func (repository *AcmeFlightRepository) FindOne(cfg *config.Config, id int) (*domain.Flight, error) {
	ctx := context.Background()
	db := repository.GetDB(cfg)
	defer db.Close()
	flight := domain.Flight{}
	err := db.NewSelect().Model(&flight).Relation("Segment").Where("fl.id = ?", id).Scan(ctx)
	return &flight, err
}

func (repository *AcmeFlightRepository) UpdateOne(cfg *config.Config, id int, flight *domain.Flight) (*domain.Flight, error) {
	ctx := context.Background()
	db := repository.GetDB(cfg)
	defer db.Close()
	err := db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		var segment domain.FlightSegment = *flight.Segment
		segmentResult, err := tx.NewUpdate().Model(&segment).Where("flsg.id = ?", segment.Id).Exec(ctx)
		if err != nil {
			repository.logger.Error(err.Error())
			return err
		}
		if _, err = segmentResult.RowsAffected(); err != nil {
			repository.logger.Error(err.Error())
			return err
		}
		flight.SegmentID = segment.Id
		flight.Segment.Id = segment.Id
		flight.Id = int64(id)
		_, err = tx.NewUpdate().Model(flight).Where("fl.id = ?", id).Exec(ctx)
		if err != nil {
			repository.logger.Error(err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		return &domain.Flight{}, err
	}
	return flight, nil
}

func (repository *AcmeFlightRepository) DeleteOne(cfg *config.Config, id int) error {
	ctx := context.Background()
	db := repository.GetDB(cfg)
	defer db.Close()
	err := db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		flight, err := repository.FindOne(cfg, id)
		if err != nil {
			return err
		}
		_, err = tx.NewDelete().Model(flight).Where("fl.id = ?", id).Exec(ctx)
		if err != nil {
			return err
		}
		_, err = tx.NewDelete().Model(flight.Segment).Where("flsg.id = ?", flight.Segment.Id).Exec(ctx)
		return err
	})
	return err
}

func (repository *AcmeFlightRepository) FindOneFlightSegment(cfg *config.Config, id int) (*domain.FlightSegment, error) {
	flight, err := repository.FindOne(cfg, id)
	if err != nil {
		return &domain.FlightSegment{}, err
	}
	return flight.Segment, nil
}
