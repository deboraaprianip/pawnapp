package usecases

import domain "pawn-clean/entity"

type warehouseResponse struct {
	IDWarehouse      int
	HargaTaksirAtas  float64
	HargaTaksirBawah float64
	MaxDuration      int64
}

type WarehouseInteractor struct {
	WarehouseRepo  domain.WarehouseRepository
	CollateralRepo domain.CollateralRepository
	ProductRepo    domain.ProductRepository
}

func (w *WarehouseInteractor) Create(ware *domain.Warehouse) (int64, error) {
	res, err := w.WarehouseRepo.FindByID(ware.ID)
	if err != nil {
		return 0, err
	}
	idNew, err := w.WarehouseRepo.Create(res)
	if err != nil {
		return 0, err
	}
	return idNew, nil
}

func (w *WarehouseInteractor) Appraise(id int) (*warehouseResponse, error) {
	warehouse, err := w.WarehouseRepo.FindByID(id)
	if err != nil {
		return &warehouseResponse{}, nil
	}
	collateral, err := w.CollateralRepo.FindByID(warehouse.Collateral.ID)
	if err != nil {
		return &warehouseResponse{}, nil
	}
	maxDuration, err := w.ProductRepo.GetMaxDurationById(collateral.Product.ID)
	if err != nil {
		return &warehouseResponse{}, nil
	}

	nilaiTaksirAtas := warehouse.GetTopAssessment(collateral.NilaiTaksirAtas, collateral.Ltv)
	nilaiTaksirBawah := warehouse.GetBottomAssessment(collateral.NilaiTaksirBawah, collateral.Ltv)

	result := &warehouseResponse{
		IDWarehouse:      id,
		HargaTaksirAtas:  nilaiTaksirAtas,
		HargaTaksirBawah: nilaiTaksirBawah,
		MaxDuration:      maxDuration,
	}
	return result, nil
}
