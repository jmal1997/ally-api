package ally_api

import (
	"github.com/jackmanlabs/errors"
	"encoding/xml"
)

//GET accounts/{id}/orders
func (client *Client) GetAccountOrders(id int) (*AccountOrdersResponse, error) {
	var (
		path   string                 = "/accounts.xml"
		target *AccountOrdersResponse = new(AccountOrdersResponse)
	)

	err := client.get(path, target)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return target, nil
}

//POST accounts/{id}/orders
func (client *Client) PostAccountOrders(id int) (*AccountsBalancesResponse, error) {
	return nil, errors.New("PostAccountOrders() not implemented.")
}

//POST accounts/{id}/orders/preview
func (client *Client) PostAccountOrderPreview(id int) (*AccountsBalancesResponse, error) {
	return nil, errors.New("PostAccountOrderPreview() not implemented.")
}

type AccountOrdersResponse struct {
	Response
	Orders []Order `xml:"orderstatus>order"`
}

type Order struct {
	XMLName xml.Name `xml:"order"`
	Fixml
	FixmlMessage FixmlMessage `xml:"fixmlmessage"`
}

type FixmlMessage struct{
	//XMLName xml.Name `xml:"fixmlmessage"`
	Text string `xml:",cdata"`
}