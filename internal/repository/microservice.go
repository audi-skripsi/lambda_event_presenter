package repository

import (
	"context"
	"time"

	"github.com/audi-skripsi/lambda_event_presenter/pkg/dto"
	"github.com/audi-skripsi/lambda_event_presenter/pkg/util/jsonutil"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	cacheKeyMicroservicesData = "audi_lambda_microservices"
)

func (r *repository) GetAllMicroservicesName(ctx context.Context) (collections []string, err error) {
	collections, err = r.mongo.ListCollectionNames(ctx, bson.D{})
	if err != nil {
		r.logger.Errorf("error get all collections: %+v", err)
	}

	return
}

func (r *repository) StoreMicroservicesData(ctx context.Context, microservicesData []dto.PublicMicroserviceData) (err error) {
	microservicesStr, err := jsonutil.ConvertToJSONString(microservicesData)
	if err != nil {
		r.logger.Errorf("error converting to json string: %+v", err)
		return
	}

	_, err = r.redis.SetEX(ctx, cacheKeyMicroservicesData, microservicesStr, time.Duration(6)*time.Hour).Result()
	if err != nil {
		r.logger.Errorf("error setting to cache for microservice data: %+v, error: %+v", microservicesData, err)
	}
	return
}
