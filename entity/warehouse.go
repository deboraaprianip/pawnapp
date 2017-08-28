package entity

type WarehouseRepository interface {
	Create(warehouse *Warehouse) (int64, error)
	FindByID(id int) (*Warehouse, error)
}

type Warehouse struct {
	ID               int
	Collateral       *Collateral
	Nama             string
	HargaPasarAtas   float64
	HargaPasarBawah  float64
	WarehouseDetails []WarehouseDetail
}

func (w *Warehouse) GetTopAssessment(taksirAtas float64, ltv int) float64 {
	hargaPasarAtas := w.HargaPasarAtas
	hargaTaksirAtas := (hargaPasarAtas * float64(ltv) * taksirAtas) / 100
	return hargaTaksirAtas
}

func (w *Warehouse) GetBottomAssessment(taksirBawah float64, ltv int) float64 {
	hargaPasarBawah := w.HargaPasarBawah
	hargaTaksirBawah := (hargaPasarBawah * float64(ltv) * taksirBawah) / 100
	return hargaTaksirBawah
}
