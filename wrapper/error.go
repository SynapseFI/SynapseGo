package wrapper

import (
	"encoding/json"
	"fmt"
)

/********** GLOBAL VARIABLES **********/

/********** TYPES **********/

type (
	// ResponseError represents an error returned by the SynapseFI API
	ResponseError struct {
		ErrorCode string      `json:"errorCode"`
		HTTPCode  string      `json:"httpCode"`
		Message   string      `json:"message"`
		Response  interface{} `json:"response"`
	}

	// ActionPending represents ERROR_CODE 10
	// Accepted, but action pending
	ActionPending struct{ Message string }

	// IncorrectClientCredentials represents ERROR_CODE 100
	// Incorrect Client Credentials
	IncorrectClientCredentials struct{ Message string }

	// IncorrectUserCredentials represents ERROR_CODE 110
	// Incorrect User Credentials
	IncorrectUserCredentials struct{ Message string }

	// UnauthorizedFingerprint represents ERROR_CODE 120
	// Unauthorized Fingerprint
	UnauthorizedFingerprint struct{ Message string }

	// PayloadError represents ERROR_CODE 200
	// Error in Payload (Error in payload formatting)
	PayloadError struct{ Message string }

	// UnauthorizedAction represents ERROR_CODE 300
	// Unauthorized action (User/Client not allowed to perform this action)
	UnauthorizedAction struct{ Message string }

	// IncorrectValues represents ERROR_CODE 400
	// Incorrect Values Supplied (eg. Insufficient balance, wrong MFA response, incorrect micro deposits)
	IncorrectValues struct{ Message string }

	// ObjectNotFound represents ERROR_CODE 404
	// Object not found
	ObjectNotFound struct{ Message string }

	// ActionNotAllowed represents ERROR_CODE 410
	// Action Not Allowed on the object (either you do not have permissions or the action on this object is not supported)
	ActionNotAllowed struct{ Message string }

	// TooManyRequests represents ERROR_CODE 429
	// Too many requests hit the API too quickly.
	TooManyRequests struct{ Message string }

	// IdempotencyConflict represents ERROR_CODE 450
	// Idempotency key already in use
	IdempotencyConflict struct{ Message string }

	// RequestFailed represents ERROR_CODE 460
	// Request Failed but not due to server error
	RequestFailed struct{ Message string }

	// ServerError represents ERROR_CODE 500
	// Server Error
	ServerError struct{ Message string }

	// ServiceUnavailable represents ERROR_CODE 503
	// Service Unavailable. The server is currently unable to handle the request due to a temporary overload or scheduled maintenance.
	ServiceUnavailable struct{ Message string }
)

/********** METHODS **********/

func (e *ActionPending) Error() string {
	return e.Message
}

func (e *IncorrectClientCredentials) Error() string {
	return e.Message
}

func (e *IncorrectUserCredentials) Error() string {
	return e.Message
}

func (e *UnauthorizedFingerprint) Error() string {
	return e.Message
}

func (e *PayloadError) Error() string {
	return e.Message
}

func (e *UnauthorizedAction) Error() string {
	return e.Message
}

func (e *IncorrectValues) Error() string {
	return e.Message
}

func (e *ObjectNotFound) Error() string {
	return e.Message
}

func (e *ActionNotAllowed) Error() string {
	return e.Message
}

func (e *TooManyRequests) Error() string {
	return e.Message
}

func (e *IdempotencyConflict) Error() string {
	return e.Message
}

func (e *RequestFailed) Error() string {
	return e.Message
}

func (e *ServerError) Error() string {
	return e.Message
}

func (e *ServiceUnavailable) Error() string {
	return e.Message
}

func handleAPIError(errorCode, message string) error {
	apiErrors := map[string]error{
		"":    nil,
		"10":  &ActionPending{message},
		"100": &IncorrectClientCredentials{message},
		"110": &IncorrectUserCredentials{message},
		"120": &UnauthorizedFingerprint{message},
		"200": &PayloadError{message},
		"300": &UnauthorizedAction{message},
		"400": &IncorrectValues{message},
		"404": &ObjectNotFound{message},
		"410": &ActionNotAllowed{message},
		"429": &TooManyRequests{message},
		"450": &IdempotencyConflict{message},
		"460": &RequestFailed{message},
		"500": &ServerError{message},
		"503": &ServiceUnavailable{message},
	}

	return apiErrors[errorCode]
}

func handleHTTPError(d []byte) error {
	data := read(d)
	errCode := data["error_code"].(string)
	httpCode := data["http_code"].(string)
	msg := data["error"].(map[string]interface{})["en"].(string)

	res := &ResponseError{
		ErrorCode: errCode,
		HTTPCode:  httpCode,
		Message:   msg,
		Response:  data,
	}

	md, err := json.MarshalIndent(&res, "", "  ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(md))

	return handleAPIError(errCode, msg)
}
