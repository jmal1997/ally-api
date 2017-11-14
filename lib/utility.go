package lib

import (
	"github.com/jackmanlabs/errors"
)

//GET utility/status
func (client *Client) GetStatus() (*StatusResponse, error) {
	var (
		path   string          = "/utility/status.xml"
		target *StatusResponse = new(StatusResponse)
	)

	err := client.get(path, target)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return target, nil
}

//GET utility/version
func (client *Client) GetVersion() (*VersionResponse, error) {
	var (
		path   string           = "/utility/version.xml"
		target *VersionResponse = new(VersionResponse)
	)

	err := client.get(path, target)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return target, nil
}

type StatusResponse struct{ Response }
type VersionResponse struct{ Response }
