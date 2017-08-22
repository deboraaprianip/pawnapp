package entity

import "time"

type SimulationRepository interface {
	Create(simulation *Simulation) (int64, error)
	FindById(id int) (*Simulation, error)
}

type Simulation struct {
	ID               int
	CreatedAt        time.Time
	NilaiPinjaman    float64
	JenisPinjaman    string
	Durasi           int
	NilaiTaksirAtas  float64
	NilaiTaksirBawah float64
	WarehouseDetails []WarehouseDetail
}
