package wrapper

import (
	"fmt"

	"github.com/pkg/errors"
)

/********** GLOBAL VARIABLES **********/

/********** TYPES **********/

type (
	// Error represents an error returned by the SynapseFI API
	Error struct {
		ErrorCode string      `json:"errorCode"`
		HTTPCode  string      `json:"httpCode"`
		Message   string      `json:"message"`
		Response  interface{} `json:"response"`
	}
)

/********** METHODS **********/

func handleAPIError(code int) {

}

func handleHTTPError(data []byte) *Error {
	d := read(data)
	msg := d["error"].(map[string]interface{})["en"].(string)

	handleStackTrace(msg)

	return &Error{
		ErrorCode: d["error_code"].(string),
		HTTPCode:  d["http_code"].(string),
		Message:   msg,
		Response:  d,
	}
}

func handleStackTrace(message string) (int, error) {
	if developerMode == true {
		cause := errors.New(message)
		err := errors.WithStack(cause)
		return fmt.Printf("%+v", err)
	}

	return 0, nil
}
