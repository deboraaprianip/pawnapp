package mysql

import (
	"database/sql"
	domain "pawnapp/entity"
)

type productRepository struct {
	Conn *sql.DB
}

func NewProductRepo(conn *sql.DB) domain.ProductRepository {
	return &productRepository{
		Conn: conn,
	}
}

func (p *productRepository) Create(product *domain.Product) (int64, error) {
	res, err := p.Conn.Exec("INSERT INTO product (nama, tgl_mulai, tgl_akhir, maks_harian, maks_mingguan, maks_bulanan, maks_terlambat, maks_perpanjangan, biaya_admin, biaya_materai, biaya_penjemputan, asuransi1, asuransi2, asuransi3, asuransi4, tarif_biaya_modal, poin_terlambat, biaya_perpanjangan, status, paid_type) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		product.Nama, product.TglMulai, product.TglAkhir, product.MaksHarian, product.MaksMingguan, product.MaksBulanan, product.MaksTerlambat, product.MaksPerpanjangan, product.BiayaAdmin, product.BiayaMaterai, product.BiayaPenjemputan,
		product.Asuransi1, product.Asuransi2, product.Asuransi3, product.Asuransi4, product.TarifBiayaModal, product.PoinTerlambat, product.BiayaPerpanjangan, product.Status, product.PaidType)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (p *productRepository) FindByID(id int) (*domain.Product, error) {
	product := &domain.Product{}
	stmt, err := p.Conn.Prepare("SELECT * FROM product WHERE id ?")
	if err != nil {
		return product, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&product)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *productRepository) GetMaxDurationById(id int) (int64, error) {
	var maxDuration int64
	stmt, err := p.Conn.Prepare("SELECT maks_mingguan FROM product WHERE id ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&maxDuration)
	if err != nil {
		return 0, err
	}
	return maxDuration, nil
}
