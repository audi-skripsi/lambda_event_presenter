package service

import (
	"context"

	"github.com/audi-skripsi/lambda_event_presenter/internal/util/microserviceutil"
	"github.com/audi-skripsi/lambda_event_presenter/pkg/dto"
)

func (s *service) GetAllMicroservicesData(ctx context.Context) (microservicesData dto.PublicMicroservicesNameResponse, err error) {
	rawMicroservicesName, err := s.repository.GetAllMicroservicesName(ctx)
	if err != nil {
		s.logger.Errorf("error getting raw microservices name, error: %+v", err)
		return
	}

	processedMicroservicesData := microserviceutil.ParseDataFromCollections(rawMicroservicesName)
	err = s.repository.StoreMicroservicesData(ctx, processedMicroservicesData)
	if err != nil {
		s.logger.Errorf("error storing microservices data of: %+v, error: %+v", processedMicroservicesData, err)
	}

	microservicesData = dto.PublicMicroservicesNameResponse{Microservices: processedMicroservicesData}
	return
}

func (s *service) GetMicroserviceDataAnalytics(ctx context.Context, id string) (resp dto.PublicMicroserviceAnalyticsResponse, err error) {
	count, err := s.repository.GetMicroserviceAllEventDataCount(ctx, id)
	if err != nil {
		s.logger.Errorf("error occured for finding data count for: %+v, error: %+v", id, err)
		return
	}

	resp.EventDataCount = count
	return
}
