package entity

type CollateralRepository interface {
	Create(coll *Collateral) (int64, error)
	FindById(id int) (*Collateral, error)
}

type Collateral struct {
	ID               int
	Product          *Product
	Nama             string
	Type             string
	Ltv              int
	NilaiTaksirBawah float64
	NilaiTaksirAtas  float64
	Warehouses       []Warehouse
}
