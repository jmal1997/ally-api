package lib

import (
	"github.com/jackmanlabs/errors"
	"net/http"
)

//GET accounts/{id}/orders
func GetAccountOrders(client *http.Client, id int) (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetAccountOrders() not implemented.")
}

//POST accounts/{id}/orders
func PostAccountOrders(client *http.Client, id int) (*AccountsBalancesResponse, error) {
	return nil, errors.New("PostAccountOrders() not implemented.")
}

//POST accounts/{id}/orders/preview
func PostAccountOrderPreview(client *http.Client, id int) (*AccountsBalancesResponse, error) {
	return nil, errors.New("PostAccountOrderPreview() not implemented.")
}
