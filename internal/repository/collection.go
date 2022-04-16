package repository

import (
	"context"

	"github.com/audi-skripsi/lambda_event_presenter/internal/constant"
)

func (r *repository) SegragateCollection(name string) (err error) {
	opt1 := r.redis.HGet(
		context.Background(),
		constant.RedisKeyCollectionCheckup,
		name,
	)

	res, err := opt1.Result()
	if err != nil {
		r.logger.Errorf("error interacting with redis for %s: %+v", name, err)
	} else {
		if res == "1" {
			r.logger.Infof("collection found in cache")
			return
		}
	}

	cmd := r.redis.HMSet(context.Background(), constant.RedisKeyCollectionCheckup, name, "1")
	if cmd.Err() != nil {
		r.logger.Errorf("error hset redis for %s: %+v", name, cmd.Err())
	}

	err = nil
	return
}
