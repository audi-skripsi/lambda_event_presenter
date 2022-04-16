package converterutil

import (
	"github.com/audi-skripsi/lambda_event_presenter/internal/model"
	"github.com/audi-skripsi/lambda_event_presenter/pkg/dto"
)

func EventLogDtoToModel(logDto dto.EventLog) (logModel model.EventLog) {
	return model.EventLog{
		UID:       logDto.UID,
		Level:     logDto.Level,
		AppName:   logDto.AppName,
		Timestamp: logDto.Timestamp,
		Data:      logDto.Data,
	}
}
