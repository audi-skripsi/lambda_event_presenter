package errors

import (
	"errors"

	"github.com/audi-skripsi/lambda_event_presenter/pkg/dto"
)

var (
	ErrBadRequest           = errors.New("bad request")
	ErrMicroserviceNotFound = errors.New("microservice not found")

	ErrInternalServer = errors.New("internal server error")
)

var errorMapping = map[error]dto.ErrorResponse{
	ErrBadRequest:           {Code: 400, Message: ErrBadRequest.Error()},
	ErrMicroserviceNotFound: {Code: 404, Message: ErrMicroserviceNotFound.Error()},
	ErrInternalServer:       {Code: 500, Message: ErrInternalServer.Error()},
}

func GetErrorResponse(er error) (errRes dto.ErrorResponse) {
	errRes, found := errorMapping[er]
	if !found {
		errRes = errorMapping[ErrInternalServer]
	}
	return
}
