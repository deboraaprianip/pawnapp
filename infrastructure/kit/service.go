package kit

import domain "pawnapp/entity"

type Sercive interface {
	Create(ware *domain.Warehouse) (int64, error)
	Appraise(id int) (*warehouseResponse, error)
}
