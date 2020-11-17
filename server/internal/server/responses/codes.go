package responses

import "net/http"

type ClientErrorCode int
type SuccessCode int

type AvailableStatusCodes struct {
	BadRequest          ClientErrorCode
	UnauthorizedRequest ClientErrorCode
	ForbiddenRequest    ClientErrorCode
	NotFoundResponse    ClientErrorCode
	Conflict            ClientErrorCode
	Success             SuccessCode
	Created             SuccessCode
	Accepted            SuccessCode
}

var StatusCodes = AvailableStatusCodes{
	BadRequest:          ClientErrorCode(http.StatusBadRequest),
	UnauthorizedRequest: ClientErrorCode(http.StatusUnauthorized),
	ForbiddenRequest:    ClientErrorCode(http.StatusForbidden),
	NotFoundResponse:    ClientErrorCode(http.StatusNotFound),
	Conflict:            ClientErrorCode(http.StatusConflict),
	Success:             SuccessCode(http.StatusOK),
	Created:             SuccessCode(http.StatusCreated),
	Accepted:            SuccessCode(http.StatusAccepted),
}
