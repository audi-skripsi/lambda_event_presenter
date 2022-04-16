package service

import (
	"github.com/audi-skripsi/lambda_event_presenter/internal/util/collectionutil"
	"github.com/audi-skripsi/lambda_event_presenter/internal/util/converterutil"
	"github.com/audi-skripsi/lambda_event_presenter/pkg/dto"
)

func (s *service) StoreEvent(event dto.EventLog) (err error) {
	collName := collectionutil.ExtractEventLogCollName(dto.EventLog(event))

	err = s.repository.SegragateCollection(collName)
	if err != nil {
		s.logger.Errorf("error segregating event: %+v", err)
		return
	}

	res, err := s.repository.InsertEvent(converterutil.EventLogDtoToModel(event), collName)
	if err != nil {
		s.logger.Errorf("error on inserting event with uid of %+v: %+v", event.UID, err)
		return
	}
	s.logger.Infof("success inserting event with details: %+v", res)
	return
}
