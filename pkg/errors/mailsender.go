package errors

import (
	"fmt"
)

type RequestError struct {
	StatusCode int
	ErrCode    uint16
	Err        error
}

/*-----------------------External calls errors-----------------------*/
const (
	CredsMarshalingError uint16 = 1 << iota
	ClientRequestError
	BadRequestError
	UnauthorisedError
	ForbiddenError
	AuthServiceBadGatewayError
	AuthServiceUnavailableError
	AuthServiceUserNotFound
	TemplateNotFound
	ExternalServiceBadGatewayError
	ExternalServiceUnavailableError
)

/*-------------------------------------------------------------------*/

type AuthServiceMappingItem struct {
	Description string
	Status      int
}

var AuthServiceErrorsDescriptionMap = map[uint16]*AuthServiceMappingItem{
	CredsMarshalingError:        {"Credentials marshaling error.", 400},
	ClientRequestError:          {"Client request error.", 400},
	BadRequestError:             {"Bad request error.", 400},
	UnauthorisedError:           {"Unauthorized.", 401},
	ForbiddenError:              {"Forbidden.", 403},
	AuthServiceBadGatewayError:  {"Auth service bad gateway.", 502},
	AuthServiceUnavailableError: {"Auth service unavailable.", 503},
	AuthServiceUserNotFound:     {"Auth service user not found.", 400},
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

func NewRequestError(errCode uint16, err error) *RequestError {
	descItem := AuthServiceErrorsDescriptionMap[errCode]

	var statusCode int
	if descItem != nil {
		statusCode = descItem.Status
	}

	return &RequestError{
		StatusCode: statusCode,
		ErrCode:    errCode,
		Err:        err,
	}
}
