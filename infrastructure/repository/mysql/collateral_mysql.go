package mysql

import (
	"database/sql"
	domain "pawn-clean/entity"
)

type collateralRepository struct {
	Conn *sql.DB
}

func NewCollateralRepo(conn *sql.DB) domain.CollateralRepository {
	return &collateralRepository{Conn: conn}
}

func (c *collateralRepository) Create(col *domain.Collateral) (int64, error) {
	res, err := c.Conn.Exec("INSERT INTO collateral (id_product, type, ltv, nilai_taksir_bawah, nilai_taksir_atas) VALUES (?, ?, ?, ?, ?)", col.Product.ID, col.Type, col.Ltv, col.NilaiTaksirBawah, col.NilaiTaksirAtas)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (c *collateralRepository) FindByID(id int) (*domain.Collateral, error) {
	col := &domain.Collateral{}
	stmt, err := c.Conn.Prepare("SELECT * FROM collateral WHERE id ?")
	if err != nil {
		return col, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&col)
	if err != nil {
		return col, err
	}
	return col, nil
}
