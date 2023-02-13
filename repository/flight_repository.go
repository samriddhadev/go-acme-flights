package repository

import (
	"context"

	"github.com/samriddhadev/go-acme-flights/config"
	"github.com/samriddhadev/go-acme-flights/domain"
)

func NewAcmeFlightRepository() AcmeFlightRepository {
	return AcmeFlightRepository {

	}
}

type AcmeFlightRepository struct {
	BaseRepository
}

func (repository *AcmeFlightRepository) FindAll(cfg *config.Config) (*[]domain.Flight, error) {
	ctx := context.Background()
	db := repository.GetDB(cfg)
	flights := []domain.Flight{}
	err := db.NewSelect().Model(&flights).Scan(ctx)
	return &flights, err
}

func (repository *AcmeFlightRepository) InsertOne(cfg *config.Config, flight *domain.Flight) error {
	ctx := context.Background()
	db := repository.GetDB(cfg)
	_, err := db.NewInsert().Model(flight).Exec(ctx)
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
	_, err := db.NewUpdate().Model(flight).Where("id = ?", id).Exec(ctx); if err != nil {
		return &domain.Flight{}, err
	}
	return repository.FindOne(cfg, id)
}

func (repository *AcmeFlightRepository) DeleteOne(cfg *config.Config, id int) error {
	ctx := context.Background()
	db := repository.GetDB(cfg)
	flight, err := repository.FindOne(cfg, id); if err != nil {
		return err
	}
	_, err = db.NewDelete().Model(&flight).Exec(ctx)
	return err
}