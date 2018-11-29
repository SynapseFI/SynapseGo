package wrapper

/********** GLOBAL VARIABLES **********/

/********** TYPES **********/

type (
	// Error represents an error returned by the SynapseFI API
	Error struct {
		ErrorCode string      `json:"errorCode"`
		HTTPCode  string      `json:"httpCode"`
		Message   interface{} `json:"message"`
		Response  interface{} `json:"response"`
	}
)

/********** METHODS **********/

var httpStatusResponses = map[string]interface{}{
	"200": "OK",
	"202": "Accepted, but not final response",
	"400": "Bad request to API. Missing a field or invalid field",
	"401": "Authentication Error",
}

func handleHTTPError(data []byte) *Error {
	d := read(data)

	return &Error{
		ErrorCode: d["error_code"].(string),
		HTTPCode:  d["http_code"].(string),
		Message:   d["error"].(map[string]interface{})["en"].(string),
		Response:  d,
	}
}

func handleAPIError(code int) {

}
