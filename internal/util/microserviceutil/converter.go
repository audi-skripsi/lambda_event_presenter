package microserviceutil

import (
	"github.com/audi-skripsi/lambda_event_presenter/internal/model"
	"github.com/audi-skripsi/lambda_event_presenter/pkg/dto"
)

func EventModelToRespDto(m *[]model.EventLog) (resp *[]dto.PublicEventData) {
	if m == nil {
		return nil
	}
	resp = &[]dto.PublicEventData{}

	for _, v := range *m {
		*resp = append(*resp, dto.PublicEventData{
			UID:         v.UID,
			Level:       v.Level,
			Message:     v.Data.(string),
			ServiceName: v.AppName,
			Timestamp:   v.Timestamp.Unix(),
		})
	}
	return
}
