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
