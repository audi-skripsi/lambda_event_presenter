package service

import (
	"time"

	"github.com/audi-skripsi/lambda_event_presenter/pkg/dto"
)

func (s *service) Ping() (resp dto.PublicPingResponse) {
	return dto.PublicPingResponse{
		Message:   "pong",
		Timestamp: time.Now().Unix(),
	}
}
