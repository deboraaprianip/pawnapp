package mysql

import (
	"database/sql"
	domain "pawnapp/entity"
)

type warehouseRepository struct {
	Conn *sql.DB
}

func NewWarehouseRepo(conn *sql.DB) domain.WarehouseRepository {
	return &warehouseRepository{Conn: conn}
}

func (w *warehouseRepository) Create(ware *domain.Warehouse) (int64, error) {
	res, err := w.Conn.Exec("INSERT INTO ware_collateral(id_kolateral, nama, harga_pasar_atas, harga_pasar_bawah) VALUES (?, ?, ?, ?)", ware.Collateral.ID, ware.Nama, ware.HargaPasarAtas, ware.HargaPasarBawah)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (w *warehouseRepository) FindByID(id int) (*domain.Warehouse, error) {
	ware := &domain.Warehouse{}
	stmt, err := w.Conn.Prepare("SELECT * FROM ware_collateral_detail WHERE id ?")
	if err != nil {
		return ware, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&ware)
	if err != nil {
		return ware, err
	}
	return ware, nil
}
