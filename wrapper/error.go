package wrapper

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
	ActionPending struct {
		ErrorCode string      `json:"errorCode"`
		HTTPCode  string      `json:"httpCode"`
		Message   string      `json:"message"`
		Response  interface{} `json:"response"`
	}

	// IncorrectClientCredentials represents ERROR_CODE 100
	// Incorrect Client Credentials
	IncorrectClientCredentials struct {
		ErrorCode string      `json:"errorCode"`
		HTTPCode  string      `json:"httpCode"`
		Message   string      `json:"message"`
		Response  interface{} `json:"response"`
	}

	// IncorrectUserCredentials represents ERROR_CODE 110
	// Incorrect User Credentials
	IncorrectUserCredentials struct {
		ErrorCode string      `json:"errorCode"`
		HTTPCode  string      `json:"httpCode"`
		Message   string      `json:"message"`
		Response  interface{} `json:"response"`
	}

	// UnauthorizedFingerprint represents ERROR_CODE 120
	// Unauthorized Fingerprint
	UnauthorizedFingerprint struct {
		ErrorCode string      `json:"errorCode"`
		HTTPCode  string      `json:"httpCode"`
		Message   string      `json:"message"`
		Response  interface{} `json:"response"`
	}

	// PayloadError represents ERROR_CODE 200
	// Error in Payload (Error in payload formatting)
	PayloadError struct {
		ErrorCode string      `json:"errorCode"`
		HTTPCode  string      `json:"httpCode"`
		Message   string      `json:"message"`
		Response  interface{} `json:"response"`
	}

	// UnauthorizedAction represents ERROR_CODE 300
	// Unauthorized action (User/Client not allowed to perform this action)
	UnauthorizedAction struct {
		ErrorCode string      `json:"errorCode"`
		HTTPCode  string      `json:"httpCode"`
		Message   string      `json:"message"`
		Response  interface{} `json:"response"`
	}

	// IncorrectValues represents ERROR_CODE 400
	// Incorrect Values Supplied (eg. Insufficient balance, wrong MFA response, incorrect micro deposits)
	IncorrectValues struct {
		ErrorCode string      `json:"errorCode"`
		HTTPCode  string      `json:"httpCode"`
		Message   string      `json:"message"`
		Response  interface{} `json:"response"`
	}

	// ObjectNotFound represents ERROR_CODE 404
	// Object not found
	ObjectNotFound struct {
		ErrorCode string      `json:"errorCode"`
		HTTPCode  string      `json:"httpCode"`
		Message   string      `json:"message"`
		Response  interface{} `json:"response"`
	}

	// ActionNotAllowed represents ERROR_CODE 410
	// Action Not Allowed on the object (either you do not have permissions or the action on this object is not supported)
	ActionNotAllowed struct {
		ErrorCode string      `json:"errorCode"`
		HTTPCode  string      `json:"httpCode"`
		Message   string      `json:"message"`
		Response  interface{} `json:"response"`
	}

	// TooManyRequests represents ERROR_CODE 429
	// Too many requests hit the API too quickly.
	TooManyRequests struct {
		ErrorCode string      `json:"errorCode"`
		HTTPCode  string      `json:"httpCode"`
		Message   string      `json:"message"`
		Response  interface{} `json:"response"`
	}

	// IdempotencyConflict represents ERROR_CODE 450
	// Idempotency key already in use
	IdempotencyConflict struct {
		ErrorCode string      `json:"errorCode"`
		HTTPCode  string      `json:"httpCode"`
		Message   string      `json:"message"`
		Response  interface{} `json:"response"`
	}

	// RequestFailed represents ERROR_CODE 460
	// Request Failed but not due to server error
	RequestFailed struct {
		ErrorCode string      `json:"errorCode"`
		HTTPCode  string      `json:"httpCode"`
		Message   string      `json:"message"`
		Response  interface{} `json:"response"`
	}

	// ServerError represents ERROR_CODE 500
	// Server Error
	ServerError struct {
		ErrorCode string      `json:"errorCode"`
		HTTPCode  string      `json:"httpCode"`
		Message   string      `json:"message"`
		Response  interface{} `json:"response"`
	}

	// ServiceUnavailable represents ERROR_CODE 503
	// Service Unavailable. The server is currently unable to handle the request due to a temporary overload or scheduled maintenance.
	ServiceUnavailable struct {
		ErrorCode string      `json:"errorCode"`
		HTTPCode  string      `json:"httpCode"`
		Message   string      `json:"message"`
		Response  interface{} `json:"response"`
	}
)

/********** METHODS **********/

func (e *ActionPending) Error() string {
	return formatErrorMessage(e.ErrorCode, e.Message)
}

func (e *IncorrectClientCredentials) Error() string {
	return formatErrorMessage(e.ErrorCode, e.Message)
}

func (e *IncorrectUserCredentials) Error() string {
	return formatErrorMessage(e.ErrorCode, e.Message)
}

func (e *UnauthorizedFingerprint) Error() string {
	return formatErrorMessage(e.ErrorCode, e.Message)
}

func (e *PayloadError) Error() string {
	return formatErrorMessage(e.ErrorCode, e.Message)
}

func (e *UnauthorizedAction) Error() string {
	return formatErrorMessage(e.ErrorCode, e.Message)
}

func (e *IncorrectValues) Error() string {
	return formatErrorMessage(e.ErrorCode, e.Message)
}

func (e *ObjectNotFound) Error() string {
	return formatErrorMessage(e.ErrorCode, e.Message)
}

func (e *ActionNotAllowed) Error() string {
	return formatErrorMessage(e.ErrorCode, e.Message)
}

func (e *TooManyRequests) Error() string {
	return formatErrorMessage(e.ErrorCode, e.Message)
}

func (e *IdempotencyConflict) Error() string {
	return formatErrorMessage(e.ErrorCode, e.Message)
}

func (e *RequestFailed) Error() string {
	return formatErrorMessage(e.ErrorCode, e.Message)
}

func (e *ServerError) Error() string {
	return formatErrorMessage(e.ErrorCode, e.Message)
}

func (e *ServiceUnavailable) Error() string {
	return formatErrorMessage(e.ErrorCode, e.Message)
}

func handleAPIError(errorCode, httpCode, message string, data map[string]interface{}) error {

	apiErrors := map[string]error{
		"":    nil,
		"10":  &ActionPending{errorCode, httpCode, message, data},
		"100": &IncorrectClientCredentials{errorCode, httpCode, message, data},
		"110": &IncorrectUserCredentials{errorCode, httpCode, message, data},
		"120": &UnauthorizedFingerprint{errorCode, httpCode, message, data},
		"200": &PayloadError{errorCode, httpCode, message, data},
		"300": &UnauthorizedAction{errorCode, httpCode, message, data},
		"400": &IncorrectValues{errorCode, httpCode, message, data},
		"404": &ObjectNotFound{errorCode, httpCode, message, data},
		"410": &ActionNotAllowed{errorCode, httpCode, message, data},
		"429": &TooManyRequests{errorCode, httpCode, message, data},
		"450": &IdempotencyConflict{errorCode, httpCode, message, data},
		"460": &RequestFailed{errorCode, httpCode, message, data},
		"500": &ServerError{errorCode, httpCode, message, data},
		"503": &ServiceUnavailable{errorCode, httpCode, message, data},
	}

	return apiErrors[errorCode]
}

func handleHTTPError(d []byte) error {
	data := read(d)

	errCode := data["error_code"].(string)
	httpCode := data["http_code"].(string)
	msg := data["error"].(map[string]interface{})["en"].(string)

	apiErr := handleAPIError(errCode, httpCode, msg, data)

	switch apiErr.(type) {
	case *IncorrectUserCredentials:

	}

	return apiErr
}

// HELPER METHODS

func formatErrorMessage(code, msg string) string {
	return "ERROR_CODE " + code + "\n" + msg
}
