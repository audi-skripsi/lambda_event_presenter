package handler

import (
	"context"
	"net/http"

	"github.com/audi-skripsi/lambda_event_presenter/pkg/dto"
	"github.com/audi-skripsi/lambda_event_presenter/pkg/util/httputil"
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
