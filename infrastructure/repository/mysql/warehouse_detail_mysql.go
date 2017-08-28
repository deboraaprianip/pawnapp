package mysql

import (
	"database/sql"
	domain "pawnapp/entity"
)

type warehouseDetailRepository struct {
	Conn *sql.DB
}

func NewWarehouseDetailRepo(conn *sql.DB) domain.WarehouseDetailRepository {
	return &warehouseDetailRepository{Conn: conn}
}

func (w *warehouseDetailRepository) Create(ware *domain.WarehouseDetail) (int64, error) {
	res, err := w.Conn.Exec("INSERT INTO ware_collateral_detail(id_simulasi, id_ware_kolateral) VALUES (?, ?, ?)", ware.Simulation.ID, ware.Warehouses.ID)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (w *warehouseDetailRepository) FindByID(id int) (*domain.WarehouseDetail, error) {
	ware := &domain.WarehouseDetail{}
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
