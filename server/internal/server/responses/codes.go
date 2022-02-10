package responses

import "net/http"

// ClientErrorCode contains HTTP codes for client error responses.
type ClientErrorCode int

// SuccessCode contains HTTP codes for success responses.
type SuccessCode int

// AvailableStatusCodes contains all available variable StatusCodes.
type AvailableStatusCodes struct {
	BadRequest          ClientErrorCode
	UnauthorizedRequest ClientErrorCode
	ForbiddenRequest    ClientErrorCode
	NotFoundResponse    ClientErrorCode
	Conflict            ClientErrorCode
	MustUpgrade         ClientErrorCode
	Success             SuccessCode
	Created             SuccessCode
	Accepted            SuccessCode
}

// StatusCodes is the filled version of AvailableStatusCodes with the correct HTTP codes.
var StatusCodes = AvailableStatusCodes{
	BadRequest:          ClientErrorCode(http.StatusBadRequest),
	UnauthorizedRequest: ClientErrorCode(http.StatusUnauthorized),
	ForbiddenRequest:    ClientErrorCode(http.StatusForbidden),
	NotFoundResponse:    ClientErrorCode(http.StatusNotFound),
	Conflict:            ClientErrorCode(http.StatusConflict),
	MustUpgrade:         ClientErrorCode(http.StatusUpgradeRequired),
	Success:             SuccessCode(http.StatusOK),
	Created:             SuccessCode(http.StatusCreated),
	Accepted:            SuccessCode(http.StatusAccepted),
}
