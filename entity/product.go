package entity

import "time"

type ProductRepository interface {
	Create(product *Product) (int64, error)
	FindByID(id int) (*Product, error)
	GetMaxDurationById(id int) (int64, error)
}

type Product struct {
	ID                int
	Nama              string
	TglMulai          time.Time
	TglAkhir          time.Time
	MaksHarian        int
	MaksMingguan      int
	MaksBulanan       int
	MaksTerlambat     int
	MaksPerpanjangan  int
	BiayaAdmin        float64
	BiayaMaterai      float64
	BiayaPenjemputan  float64
	Asuransi1         float64
	Asuransi2         float64
	Asuransi3         float64
	Asuransi4         float64
	TarifBiayaModal   float64
	PoinTerlambat     float64
	BiayaPerpanjangan float64
	Status            int
	PaidType          string
	Collaterals       []Collateral
}
