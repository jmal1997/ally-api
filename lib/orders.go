package lib

import (
	"github.com/jackmanlabs/errors"
)

//GET accounts/{id}/orders
func (client *Client) GetAccountOrders(id int) (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetAccountOrders() not implemented.")
}

//POST accounts/{id}/orders
func (client *Client) PostAccountOrders(id int) (*AccountsBalancesResponse, error) {
	return nil, errors.New("PostAccountOrders() not implemented.")
}

//POST accounts/{id}/orders/preview
func (client *Client) PostAccountOrderPreview(id int) (*AccountsBalancesResponse, error) {
	return nil, errors.New("PostAccountOrderPreview() not implemented.")
}
