package router

import (
	"net/http"

	"github.com/audi-skripsi/lambda_event_presenter/cmd/webservice/handler"
	"github.com/audi-skripsi/lambda_event_presenter/internal/config"
	"github.com/audi-skripsi/lambda_event_presenter/internal/service"
	"github.com/gorilla/mux"
)

type RouterParams struct {
	Conf    *config.Config
	Service service.Service
	Router  *mux.Router
}

func Init(params *RouterParams) {
	params.Router.HandleFunc(PingPath, handler.HandlePing(params.Service.Ping)).Methods(http.MethodGet)
}
