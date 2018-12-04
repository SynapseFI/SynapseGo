package wrapper

/********** GLOBAL VARIABLES **********/

/********** TYPES **********/

type (
	// Statement represents a single statement object
	Statement struct {
		StatementID string      `json:"_id"`
		DateEnd     string      `json:"date_end"`
		DateStart   string      `json:"date_start"`
		URLs        interface{} `json:"urls"`
		Response    interface{}
	}

	// Statements represents a list of transaction objects
	Statements struct {
		Limit          int64       `json:"limit"`
		Page           int64       `json:"page"`
		PageCount      int64       `json:"page_count"`
		StatementCount int64       `json:"statements_count"`
		Statements     []Statement `json:"statements"`
	}
)
