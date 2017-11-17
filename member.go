package ally_api

import (
	"github.com/jackmanlabs/errors"
)

//GET member/profile
func (client *Client) GetProfile() (*ProfileResponse, error) {
	var (
		path   string           = "/market/clock.xml"
		target *ProfileResponse = new(ProfileResponse)
	)

	err := client.get(path, target)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return target, nil
}

type ProfileResponse struct {
	Response
	Date     string  `xml:"date"`
	Status   Status  `xml:"status"`
	Message  string  `xml:"message"`
	UnixTime float64 `xml:"unixtime"`
}

type Status struct {
	Current  string `xml:"current"`
	Next     string `xml:"next"`
	ChangeAt string `xml:"change_at"`
}
