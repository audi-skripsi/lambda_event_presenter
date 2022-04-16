package collectionutil

import (
	"fmt"
	"time"

	"github.com/audi-skripsi/lambda_event_presenter/pkg/dto"
)

const CollNameFormat = "%s_%d_%d_%s"

func ExtractEventLogCollName(event dto.EventLog) string {
	appName := event.AppName

	nowTime := time.Now().UTC()
	currYear := nowTime.Year()
	currMonth := int(nowTime.Month())

	return fmt.Sprintf(CollNameFormat, appName, currYear, currMonth, event.Level)
}
