package dto

import (
	"sync"

	"github.com/audi-skripsi/lambda_event_presenter/internal/model"
)

type EventBatch struct {
	*sync.Mutex

	CollName  string
	EventData []model.EventLog
}
