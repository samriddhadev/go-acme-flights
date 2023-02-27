package service

import (
	"errors"

	"github.com/samriddhadev/go-acme-flights/core/config"
	"github.com/samriddhadev/go-acme-flights/core/logger"
	"github.com/samriddhadev/go-acme-flights/domain"
	"github.com/samriddhadev/go-acme-flights/model"
	"github.com/samriddhadev/go-acme-flights/repository"
)

func NewAcmeFlightService(logger *logger.AcmeLogger, flightRepository repository.AcmeFlightRepository) AcmeFlightService {
	return AcmeFlightService{
		logger : logger,
		flightRepository: flightRepository,
	}
}

type AcmeFlightService struct {
	logger *logger.AcmeLogger
	flightRepository repository.AcmeFlightRepository
}

func (service *AcmeFlightService) GetFlights(cfg *config.Config, flightFilter *model.FlightFilter) (*[]model.Flight, error) {
	flights := []model.Flight{}
	entities, err := service.flightRepository.FindAll(cfg, flightFilter)
	if err != nil {
		return &[]model.Flight{}, err
	}
	for _, entity := range *entities {
		flights = append(flights, model.Flight{
			Id:                     entity.Id,
			Name:                   entity.Segment.Name,
			Origin:                 entity.Segment.Origin,
			Destination:            entity.Segment.Destination,
			Miles:                  entity.Segment.Miles,
			ScheduledDepartureTime: entity.ScheduledDepartureTime,
			ScheduledArrivalTime:   entity.ScheduledArrivalTime,
			FirstClassBaseCost:     entity.FirstClassBaseCost,
			EconomyClassBaseCost:   entity.EconomyClassBaseCost,
			NumFirstClassSeats:     entity.NumFirstClassSeats,
			NumEconomyClassSeats:   entity.NumEconomyClassSeats,
			AirplaneTypeId:         "12S-BOEING",
		})
	}
	return &flights, nil
}

func (service *AcmeFlightService) CreateFlight(cfg *config.Config, flight *model.Flight) error {
	segmentEntity := domain.FlightSegment{
		Name:        flight.Name,
		Origin:      flight.Origin,
		Destination: flight.Destination,
		Miles:       flight.Miles,
	}
	entity := domain.Flight{
		ScheduledDepartureTime: flight.ScheduledDepartureTime,
		ScheduledArrivalTime:   flight.ScheduledArrivalTime,
		FirstClassBaseCost:     flight.FirstClassBaseCost,
		EconomyClassBaseCost:   flight.EconomyClassBaseCost,
		NumFirstClassSeats:     flight.NumFirstClassSeats,
		NumEconomyClassSeats:   flight.NumEconomyClassSeats,
		AirplaneTypeId:         flight.AirplaneTypeId,
		Segment:                &segmentEntity,
	}
	err := service.flightRepository.InsertOne(cfg, &entity)
	if err != nil {
		return err
	}
	return nil
}

func (service *AcmeFlightService) GetFlight(cfg *config.Config, id int) (*model.Flight, error) {
	flight := model.Flight{}
	entity, err := service.flightRepository.FindOne(cfg, id)
	if err != nil {
		return &model.Flight{}, err
	}
	flight = model.Flight{
		Id:                     entity.Id,
		Name:                   entity.Segment.Name,
		Origin:                 entity.Segment.Origin,
		Destination:            entity.Segment.Destination,
		Miles:                  entity.Segment.Miles,
		ScheduledDepartureTime: entity.ScheduledDepartureTime,
		ScheduledArrivalTime:   entity.ScheduledArrivalTime,
		FirstClassBaseCost:     entity.FirstClassBaseCost,
		EconomyClassBaseCost:   entity.EconomyClassBaseCost,
		NumFirstClassSeats:     entity.NumFirstClassSeats,
		NumEconomyClassSeats:   entity.NumEconomyClassSeats,
		AirplaneTypeId:         "12S-BOEING",
	}
	return &flight, nil
}

func (service *AcmeFlightService) UpdateFlight(cfg *config.Config, id int, flight *model.Flight) (*model.Flight, error) {
	segment, err := service.flightRepository.FindOneFlightSegment(cfg, id)
	if err != nil {
		return &model.Flight{}, errors.New("segment not available")
	}
	segmentEntity := domain.FlightSegment{
		Id:          segment.Id,
		Name:        flight.Name,
		Origin:      flight.Origin,
		Destination: flight.Destination,
		Miles:       flight.Miles,
	}
	entity := domain.Flight{
		Id:                     flight.Id,
		ScheduledDepartureTime: flight.ScheduledDepartureTime,
		ScheduledArrivalTime:   flight.ScheduledArrivalTime,
		FirstClassBaseCost:     flight.FirstClassBaseCost,
		EconomyClassBaseCost:   flight.EconomyClassBaseCost,
		NumFirstClassSeats:     flight.NumFirstClassSeats,
		NumEconomyClassSeats:   flight.NumEconomyClassSeats,
		AirplaneTypeId:         flight.AirplaneTypeId,
		Segment:                &segmentEntity,
	}
	var flightEntity *domain.Flight
	flightEntity, err = service.flightRepository.UpdateOne(cfg, id, &entity)
	if err != nil {
		return &model.Flight{}, err
	}
	updatedFlight := model.Flight{
		Id:                     flightEntity.Id,
		Name:                   flightEntity.Segment.Name,
		Origin:                 flightEntity.Segment.Origin,
		Destination:            flightEntity.Segment.Destination,
		Miles:                  flightEntity.Segment.Miles,
		ScheduledDepartureTime: flightEntity.ScheduledDepartureTime,
		ScheduledArrivalTime:   flightEntity.ScheduledArrivalTime,
		FirstClassBaseCost:     flightEntity.FirstClassBaseCost,
		EconomyClassBaseCost:   flightEntity.EconomyClassBaseCost,
		NumFirstClassSeats:     flightEntity.NumFirstClassSeats,
		NumEconomyClassSeats:   flightEntity.NumEconomyClassSeats,
		AirplaneTypeId:         flightEntity.AirplaneTypeId,
	}
	return &updatedFlight, nil
}

func (service *AcmeFlightService) DeleteFlight(cfg *config.Config, id int) error {
	err := service.flightRepository.DeleteOne(cfg, id)
	if err != nil {
		return err
	}
	return nil
}