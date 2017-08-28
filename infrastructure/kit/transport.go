package kit

import (
	"context"
	"errors"
	"net/http"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(ctx context.Context, ws Sercive, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(),
	}

	warehouseHandler := kithttp.NewServer(
		makeWarehouseEndpoint(ws),
		decodeWarehouseRequest,
		encodeResponse,
	)
}

var errBadRoute = errors.New("bad route")

func decodeWarehouseRequest(r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	return id, nil
}

func encodeResponse(w http.ResponseWriter, response interface{}) error {

}

func encodeError(ctx context.Context, err error, w *http.ResponseWriter){
	statusCode := http.StatusOK
	if err != nil{
		switch u. {
			case err.

		}
	}
}