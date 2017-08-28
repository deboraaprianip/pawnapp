package usecases

import (
	"errors"
	domain "pawn-clean/entity"
)

var ErrInvalidArgument = errors.New("Invalid Argumnent")

type wareCollateralInteractor struct {
	wareDetailRepo domain.WarehouseDetailRepository
	wareCollRepo   domain.WarehouseRepository
	simulationRepo domain.SimulationRepository
}

func (w *wareCollateralInteractor) Create(ware *domain.WarehouseDetail) (int64, error) {
	_, err := w.wareCollRepo.FindByID(ware.Warehouses.ID)
	if err != nil {
		return 0, err
	}

	_, err = w.simulationRepo.FindByID(ware.Simulation.ID)
	if err != nil {
		return 0, err
	}

	idNew, err := w.wareDetailRepo.Create(ware)
	if err != nil {
		return 0, err
	}
	return idNew, nil
}
