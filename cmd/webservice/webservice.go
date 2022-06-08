package webservice

import (
	"net/http"

	"github.com/audi-skripsi/lambda_event_presenter/cmd/webservice/router"
	"github.com/audi-skripsi/lambda_event_presenter/internal/config"
	"github.com/audi-skripsi/lambda_event_presenter/internal/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type InitWebserviceParams struct {
	Logger  *logrus.Entry
	Conf    *config.Config
	Service service.Service
}

func Init(params *InitWebserviceParams) {
	var err error
	r := mux.NewRouter()

	router.Init(&router.RouterParams{
		Conf:    params.Conf,
		Service: params.Service,
		Router:  r,
	})

	err = http.ListenAndServe(params.Conf.AppAddress, r)
	if err != nil {
		params.Logger.Fatalf("[Init] error listening to: %s, error: %+v", params.Conf.AppAddress, err)
	}
}
