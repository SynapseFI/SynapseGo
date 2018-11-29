package wrapper

/********** GLOBAL VARIABLES **********/
const institutionsURL = _url + "/institutions"

/********** TYPES **********/
type (
	// Institutions represents a list of Synapse institutions
	Institutions struct {
		Banks interface{} `json:"banks"`
	}
)

/********** METHODS **********/

// GetInstitutions returns all of the nodes associated with a user
func (c *Client) GetInstitutions() (*Institutions, *Error) {
	var institutions Institutions

	h := c.getHeaderInfo("")
	req := c.newRequest(h)

	_, err := req.Get(institutionsURL, "", &institutions)

	if err != nil {
		return nil, err
	}

	return &institutions, nil
}
