package mysql

import (
	"database/sql"
	domain "pawn-clean/entity"
)

type simulationRepository struct {
	Conn *sql.DB
}

func NewSimulationRepo(conn *sql.DB) domain.SimulationRepository {
	return simulationRepository{Conn: conn}
}

func (s *simulationRepository) Create(sim *domain.Simulation) (int64, error) {
	res, err := s.Conn.Exec("INSERT INTO simulation (created_at, nilai_pinjaman, jenis_pinjaman, durasi, nilai_taksir_atas, nilai_taksir_bawah) VALUES (?, ?, ?, ?, ?, ?)", sim.CreatedAt, sim.NilaiPinjaman, sim.JenisPinjaman, sim.Durasi, sim.NilaiTaksirAtas, sim.NilaiTaksirBawah)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (s *simulationRepository) FindByID(id int) (*domain.Simulation, error) {
	sim := &domain.Simulation{}
	stmt, err := s.Conn.Prepare("SELECT * FROM simulation WHERE id ?")
	if err != nil {
		return sim, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&sim)
	if err != nil {
		return sim, err
	}
	return sim, nil
}
