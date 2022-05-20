package service

import (
	"sync"
	"time"

	indto "github.com/audi-skripsi/lambda_event_presenter/internal/dto"
	"github.com/audi-skripsi/lambda_event_presenter/internal/model"
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

	err = s.batchInsertEvent(converterutil.EventLogDtoToModel(event), collName)
	if err != nil {
		s.logger.Errorf("error on inserting event with uid of %+v: %+v", event.UID, err)
		return
	}

	return
}

func (s *service) batchInsertEvent(event model.EventLog, collName string) (err error) {
	var batch *indto.EventBatch
	batch, ok := s.BatchMap[collName]
	if !ok {
		batch = &indto.EventBatch{
			Mu:        &sync.Mutex{},
			CollName:  collName,
			EventData: []model.EventLog{},
		}
		s.BatchMap[collName] = batch
	}

	batch.Mu.Lock()
	batch.EventData = append(batch.EventData, event)
	if len(batch.EventData) > 50 {
		err = s.repository.BatchInsertEvent(batch)
		if err != nil {
			s.logger.Errorf("error inserting event: %+v", err)
		}
		s.logger.Infof("success inserting event: %+v", len(batch.EventData))
		batch.EventData = nil
	}
	batch.Mu.Unlock()
	return
}

func (s *service) initBatchCron() {
	go func() {
		for {
			for _, v := range s.BatchMap {
				go func(batch *indto.EventBatch) {
					if len(batch.EventData) > 0 {
						batch.Mu.Lock()
						err := s.repository.BatchInsertEvent(batch)
						if err != nil {
							s.logger.Errorf("error batch insert: %+v", err)
						}
						batch.EventData = nil
						batch.Mu.Unlock()
					}
				}(v)
			}
			time.Sleep(10 * time.Second)
		}
	}()
}
