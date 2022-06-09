package handler

import (
	"context"
	"net/http"

	indto "github.com/audi-skripsi/lambda_event_presenter/internal/dto"
	"github.com/audi-skripsi/lambda_event_presenter/internal/util/microserviceutil"
	"github.com/audi-skripsi/lambda_event_presenter/pkg/dto"
	"github.com/audi-skripsi/lambda_event_presenter/pkg/errors"
	"github.com/audi-skripsi/lambda_event_presenter/pkg/util/httputil"
	"github.com/gorilla/mux"
)

type GetAllMicroservicesDataHandler func(ctx context.Context) (microservicesData dto.PublicMicroservicesNameResponse, err error)

func HandleGetAllMicroservicesData(handler GetAllMicroservicesDataHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := handler(r.Context())
		if err != nil {
			httputil.WriteErrorResponse(w, err)
			return
		}
		httputil.WriteSuccessResponse(w, resp)
	}
}

type GetMicroservicesDataAnalyticsHandler func(ctx context.Context, id string) (resp dto.PublicMicroserviceAnalyticsResponse, err error)

func HandleGetAllMicroservicesDataAnalytics(handler GetMicroservicesDataAnalyticsHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paths := mux.Vars(r)
		id, ok := paths["id"]
		if !ok {
			httputil.WriteErrorResponse(w, errors.ErrBadRequest)
			return
		}

		resp, err := handler(r.Context(), id)
		if err != nil {
			httputil.WriteErrorResponse(w, errors.ErrBadRequest)
			return
		}
		httputil.WriteSuccessResponse(w, resp)
	}
}

type GetMicroserviceEventsHandler func(ctx context.Context, id string, criteria indto.SearchEventCriteria) (resp dto.PublicMicroserviceEventSearchResponse, err error)

func HandleGetMicroserviceEvents(handler GetMicroserviceEventsHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paths := mux.Vars(r)
		id, ok := paths["id"]
		if !ok {
			httputil.WriteErrorResponse(w, errors.ErrBadRequest)
			return
		}

		criteria := microserviceutil.ReadCriteriaFromRequest(r)

		resp, err := handler(r.Context(), id, criteria)
		if err != nil {
			httputil.WriteErrorResponse(w, err)
			return
		}

		httputil.WriteSuccessResponse(w, resp)
	}
}
