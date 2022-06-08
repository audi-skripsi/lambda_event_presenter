package handler

import (
	"net/http"

	"github.com/audi-skripsi/lambda_event_presenter/pkg/dto"
	"github.com/audi-skripsi/lambda_event_presenter/pkg/util/httputil"
)

type PingHandler func() (resp dto.PublicPingResponse)

func HandlePing(handler PingHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := handler()
		httputil.WriteSuccessResponse(w, resp)
	}
}
