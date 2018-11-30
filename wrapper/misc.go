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
func (c *Client) GetInstitutions() *Institutions {
	var institutions Institutions

	req := c.newRequest()

	req.Get(institutionsURL, "", &institutions)

	return &institutions
}
