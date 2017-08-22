package entity

type WarehouseDetailRepository interface {
	Create(warehouseDetail *WarehouseDetail) (int64, error)
	FindById(id int) (*WarehouseDetail, error)
}

type WarehouseDetail struct {
	ID         int
	Simulation *Simulation
	Warehouses *Warehouse
}
