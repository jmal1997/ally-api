package lib

import (
	"github.com/jackmanlabs/errors"
	"net/http"
)

//GET utility/status
func (client *Client)GetStatus() (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetStatus() not implemented.")
}

//GET utility/version
func (client *Client)GetVersion() (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetVersion() not implemented.")
}
