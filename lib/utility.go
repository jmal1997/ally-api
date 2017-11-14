package lib

import (
	"github.com/jackmanlabs/errors"
	"net/http"
)

//GET utility/status
func GetStatus(client *http.Client) (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetStatus() not implemented.")
}

//GET utility/version
func GetVersion(client *http.Client) (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetVersion() not implemented.")
}
