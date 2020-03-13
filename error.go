package synapse

/********** GLOBAL VARIABLES **********/

/********** TYPES **********/

type (
	// ResponseError represents an error returned by the SynapseFI API
	ResponseError struct {
		ErrorCode string `json:"errorCode"`
		HTTPCode  string `json:"httpCode"`
		Message   string `json:"message"`
	}

	// ActionPending represents ERROR_CODE 10
	// Accepted, but action pending
	ActionPending struct {
		ErrorCode string `json:"errorCode"`
		HTTPCode  string `json:"httpCode"`
		Message   string `json:"message"`
	}

	// IncorrectClientCredentials represents ERROR_CODE 100
	// Incorrect Client Credentials
	IncorrectClientCredentials struct {
		ErrorCode string `json:"errorCode"`
		HTTPCode  string `json:"httpCode"`
		Message   string `json:"message"`
	}

	// IncorrectUserCredentials represents ERROR_CODE 110
	// Incorrect User Credentials
	IncorrectUserCredentials struct {
		ErrorCode string `json:"errorCode"`
		HTTPCode  string `json:"httpCode"`
		Message   string `json:"message"`
	}

	// UnauthorizedFingerprint represents ERROR_CODE 120
	// Unauthorized Fingerprint
	UnauthorizedFingerprint struct {
		ErrorCode string `json:"errorCode"`
		HTTPCode  string `json:"httpCode"`
		Message   string `json:"message"`
	}

	// PayloadError represents ERROR_CODE 200
	// Error in Payload (Error in payload formatting)
	PayloadError struct {
		ErrorCode string `json:"errorCode"`
		HTTPCode  string `json:"httpCode"`
		Message   string `json:"message"`
	}

	// UnauthorizedAction represents ERROR_CODE 300
	// Unauthorized action (User/Client not allowed to perform this action)
	UnauthorizedAction struct {
		ErrorCode string `json:"errorCode"`
		HTTPCode  string `json:"httpCode"`
		Message   string `json:"message"`
	}

	// IncorrectValues represents ERROR_CODE 400
	// Incorrect Values Supplied (eg. Insufficient balance, wrong MFA response, incorrect micro deposits)
	IncorrectValues struct {
		ErrorCode string `json:"errorCode"`
		HTTPCode  string `json:"httpCode"`
		Message   string `json:"message"`
	}

	// ObjectNotFound represents ERROR_CODE 404
	// Object not found
	ObjectNotFound struct {
		ErrorCode string `json:"errorCode"`
		HTTPCode  string `json:"httpCode"`
		Message   string `json:"message"`
	}

	// ActionNotAllowed represents ERROR_CODE 410
	// Action Not Allowed on the object (either you do not have permissions or the action on this object is not supported)
	ActionNotAllowed struct {
		ErrorCode string `json:"errorCode"`
		HTTPCode  string `json:"httpCode"`
		Message   string `json:"message"`
	}

	// TooManyRequests represents ERROR_CODE 429
	// Too many requests hit the API too quickly.
	TooManyRequests struct {
		ErrorCode string `json:"errorCode"`
		HTTPCode  string `json:"httpCode"`
		Message   string `json:"message"`
	}

	// IdempotencyConflict represents ERROR_CODE 450
	// Idempotency key already in use
	IdempotencyConflict struct {
		ErrorCode string `json:"errorCode"`
		HTTPCode  string `json:"httpCode"`
		Message   string `json:"message"`
	}

	// RequestFailed represents ERROR_CODE 460
	// Request Failed but not due to server error
	RequestFailed struct {
		ErrorCode string `json:"errorCode"`
		HTTPCode  string `json:"httpCode"`
		Message   string `json:"message"`
	}

	// ServerError represents ERROR_CODE 500
	// Server Error
	ServerError struct {
		ErrorCode string `json:"errorCode"`
		HTTPCode  string `json:"httpCode"`
		Message   string `json:"message"`
	}

	// ServiceUnavailable represents ERROR_CODE 503
	// Service Unavailable. The server is currently unable to handle the request due to a temporary overload or scheduled maintenance.
	ServiceUnavailable struct {
		ErrorCode string `json:"errorCode"`
		HTTPCode  string `json:"httpCode"`
		Message   string `json:"message"`
	}

	// ServerTimeout represents ERROR_CODE 504
	// Server Timeout
	ServerTimeout struct {
		ErrorCode string `json:"errorCode"`
		HTTPCode  string `json:"httpCode"`
		Message   string `json:"message"`
	}

	// DefaultError represent and unhandled HTTP error
	// Pass this instead of nil
	DefaultError struct {
		ErrorCode string "000"
		HTTPCode string "000"
		Message string "Error code not found"
	}
)

/********** METHODS **********/

func (e *ActionPending) Error() string {
	return formatErrorMessage(e.HTTPCode, e.ErrorCode, e.Message)
}

func (e *IncorrectClientCredentials) Error() string {
	return formatErrorMessage(e.HTTPCode, e.ErrorCode, e.Message)
}

func (e *IncorrectUserCredentials) Error() string {
	return formatErrorMessage(e.HTTPCode, e.ErrorCode, e.Message)
}

func (e *UnauthorizedFingerprint) Error() string {
	return formatErrorMessage(e.HTTPCode, e.ErrorCode, e.Message)
}

func (e *PayloadError) Error() string {
	return formatErrorMessage(e.HTTPCode, e.ErrorCode, e.Message)
}

func (e *UnauthorizedAction) Error() string {
	return formatErrorMessage(e.HTTPCode, e.ErrorCode, e.Message)
}

func (e *IncorrectValues) Error() string {
	return formatErrorMessage(e.HTTPCode, e.ErrorCode, e.Message)
}

func (e *ObjectNotFound) Error() string {
	return formatErrorMessage(e.HTTPCode, e.ErrorCode, e.Message)
}

func (e *ActionNotAllowed) Error() string {
	return formatErrorMessage(e.HTTPCode, e.ErrorCode, e.Message)
}

func (e *TooManyRequests) Error() string {
	return formatErrorMessage(e.HTTPCode, e.ErrorCode, e.Message)
}

func (e *IdempotencyConflict) Error() string {
	return formatErrorMessage(e.HTTPCode, e.ErrorCode, e.Message)
}

func (e *RequestFailed) Error() string {
	return formatErrorMessage(e.HTTPCode, e.ErrorCode, e.Message)
}

func (e *ServerError) Error() string {
	return formatErrorMessage(e.HTTPCode, e.ErrorCode, e.Message)
}

func (e *ServiceUnavailable) Error() string {
	return formatErrorMessage(e.HTTPCode, e.ErrorCode, e.Message)
}

func (e *ServerTimeout) Error() string {
	return formatErrorMessage(e.HTTPCode, e.ErrorCode, e.Message)
}

func (e *DefaultError) Error() string {
	return formatErrorMessage(e.HTTPCode, e.ErrorCode, e.Message)
}

func handleAPIError(errorCode, httpCode, message string) error {
	apiErrors := map[string]error{
		"":    &DefaultError{},
		"10":  &ActionPending{errorCode, httpCode, message},
		"100": &IncorrectClientCredentials{errorCode, httpCode, message},
		"110": &IncorrectUserCredentials{errorCode, httpCode, message},
		"120": &UnauthorizedFingerprint{errorCode, httpCode, message},
		"200": &PayloadError{errorCode, httpCode, message},
		"300": &UnauthorizedAction{errorCode, httpCode, message},
		"400": &IncorrectValues{errorCode, httpCode, message},
		"404": &ObjectNotFound{errorCode, httpCode, message},
		"410": &ActionNotAllowed{errorCode, httpCode, message},
		"429": &TooManyRequests{errorCode, httpCode, message},
		"450": &IdempotencyConflict{errorCode, httpCode, message},
		"460": &RequestFailed{errorCode, httpCode, message},
		"500": &ServerError{errorCode, httpCode, message},
		"503": &ServiceUnavailable{errorCode, httpCode, message},
		"504": &ServerTimeout{errorCode, httpCode, message},
	}

	return apiErrors[error_code]
	
}

func handleHTTPError(d []byte) error {
	data := readStream(d)

	errCode := data["error_code"].(string)
	httpCode := data["http_code"].(string)
	msg := data["error"].(map[string]interface{})["en"].(string)

	return handleAPIError(errCode, httpCode, msg)
}

// HELPER METHODS

func formatErrorMessage(httpCode, errorCode, msg string) string {
	return "http_code " + httpCode + " error_code " + errorCode + " " + msg
}

func formatErrorObject(httpCode, errorCode, msg string) map[string]interface{} {
	return map[string]interface{}{
		"http_code":  httpCode,
		"error_code": errorCode,
		"error":      msg,
	}
}
