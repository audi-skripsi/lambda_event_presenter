package jsonutil

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/audi-skripsi/lambda_event_presenter/pkg/errors"
)

func ConvertToObject(r *http.Request, i interface{}) (err error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = errors.ErrBadRequest
		return
	}

	err = json.Unmarshal(b, i)
	if err != nil {
		err = errors.ErrBadRequest
		return
	}

	return
}
