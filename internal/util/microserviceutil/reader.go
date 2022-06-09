package microserviceutil

import (
	"net/http"
	"strconv"
	"time"

	"github.com/audi-skripsi/lambda_event_presenter/internal/dto"
)

func ReadCriteriaFromRequest(r *http.Request) (resp dto.SearchEventCriteria) {
	params := r.URL.Query()

	level := params.Get("level")
	if level != "" {
		resp.Level = &level
	}

	message := params.Get("message")
	if message != "" {
		resp.Message = &message
	}

	timeStart := params.Get("timeStart")
	if timeStart != "" {
		timestampStart, err := strconv.ParseInt(timeStart, 10, 64)
		if err == nil {
			dateTimeStart := time.Unix(timestampStart, 0)
			resp.TimeStart = &dateTimeStart
		}
	}

	timeEnd := params.Get("timeEnd")
	if timeEnd != "" {
		timestampStart, err := strconv.ParseInt(timeEnd, 10, 64)
		if err == nil {
			dateTimeEnd := time.Unix(timestampStart, 0)
			resp.TimeEnd = &dateTimeEnd
		}
	}

	return
}
