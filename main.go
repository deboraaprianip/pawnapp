package main

import (
	"context"
	"net/http"
	"os"
	"pawn-clean/config"
	"pawn-clean/infrastructure"
	handler "pawn-clean/infrastructure/kit"
	repo "pawn-clean/infrastructure/repository/mysql"
	"pawn-clean/usecases"

	"github.com/go-kit/kit/log"
)

func main() {
	var ctx = context.Background()
	// log setting, nantinya akan diganti jadi gunain logrus
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = &serializedLogger{Logger: logger}
	logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)

	settings := config.LoadConfig()
	conn := infrastructure.NewSqlHandler(settings.Db)
	defer conn.Close()

	warehouseDetailRepo := repo.NewWarehouseDetailRepo(conn)
	warehouseRepo := repo.NewWarehouseRepo(conn)
	collateralRepo := repo.NewCollateralRepo(conn)
	productRepo := repo.NewProductRepo(conn)

	warehouseInteractor := new(usecases.WarehouseInteractor)
	warehouseInteractor.WarehouseRepo = warehouseRepo
	warehouseInteractor.CollateralRepo = collateralRepo
	warehouseInteractor.ProductRepo = productRepo

	httpLogger := log.NewContext(logger).With("component", "http")
	mux := http.NewServeMux()

	mux.Handle("/pawn/", handler.MakeHandler())

}
