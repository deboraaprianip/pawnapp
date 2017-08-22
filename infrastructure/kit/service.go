package kit

import domain "pawn-clean/entity"

type Sercive interface {
	Create(ware *domain.Warehouse) (int64, error)
	Appraise(id int) (*warehouseResponse, error)
}
