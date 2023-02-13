package service

import (
	"github.com/samriddhadev/go-acme-flights/config"
	"github.com/samriddhadev/go-acme-flights/domain"
	"github.com/samriddhadev/go-acme-flights/model"
	"github.com/samriddhadev/go-acme-flights/repository"
)


func NewAcmeFlightService(flightRepository repository.AcmeFlightRepository) AcmeFlightService {
	return AcmeFlightService{
		flightRepository: flightRepository,
	}
}

type AcmeFlightService struct {
	flightRepository repository.AcmeFlightRepository
}

 func (service *AcmeFlightService) GetFlights(cfg *config.Config) (*[]model.Flight, error) {
	flights := []model.Flight{}
	entities, err := service.flightRepository.FindAll(cfg)
	if err != nil {
		return &[]model.Flight{}, err
	}
	for _, entity := range *entities {
		flights = append(flights, model.Flight{
			Id: entity.Id,
			Name: entity.Segment.Name,
			Origin: entity.Segment.Origin,
			Destination: entity.Segment.Destination,
			Miles: entity.Segment.Miles,
			ScheduledDepartureTime: entity.ScheduledDepartureTime,
			ScheduledArrivalTime: entity.ScheduledArrivalTime,
			FirstClassBaseCost: entity.FirstClassBaseCost,
			EconomyClassBaseCost: entity.EconomyClassBaseCost,
			NumFirstClassSeats: entity.NumFirstClassSeats,
			NumEconomyClassSeats:  entity.NumEconomyClassSeats,
			AirplaneTypeId: "12S-BOEING",
		})
	}
	return &flights, nil
}

func (service *AcmeFlightService) CreateFlight(cfg *config.Config, flight *model.Flight) error {
	segmentEntity := domain.FlightSegment{
		Name: flight.Name,
		Origin: flight.Origin,
		Destination: flight.Destination,
		Miles: flight.Miles,
	}
	entity := domain.Flight{
		ScheduledDepartureTime: flight.ScheduledDepartureTime,
		ScheduledArrivalTime: flight.ScheduledArrivalTime,
		FirstClassBaseCost: flight.FirstClassBaseCost,
		EconomyClassBaseCost: flight.EconomyClassBaseCost,
		NumFirstClassSeats: flight.NumFirstClassSeats,
		NumEconomyClassSeats: flight.NumEconomyClassSeats,
		AirplaneTypeId: flight.AirplaneTypeId,
		Segment: &segmentEntity,
	}
	err := service.flightRepository.InsertOne(cfg, &entity)
	if err != nil {
		return err
	}
	return nil
}