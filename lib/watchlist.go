package lib

import (
	"github.com/jackmanlabs/errors"
	"net/http"
)

//GET watchlists
func (client *Client)GetWatchlists() (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetWatchlists() not implemented.")
}

//POST watchlists
func (client *Client)PostWatchlists() (*AccountsBalancesResponse, error) {
	return nil, errors.New("PostWatchlists() not implemented.")
}

//GET watchlists/{id}
func (client *Client)GetWatchlist(id int) (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetWatchlist() not implemented.")
}

//DELETE watchlists/{id}
func (client *Client)DeleteWatchlist(id int) (*AccountsBalancesResponse, error) {
	return nil, errors.New("DeleteWatchlist() not implemented.")
}

//POST watchlists/{id}/symbols
func (client *Client)PostWatchlistSymbols(id int) (*AccountsBalancesResponse, error) {
	return nil, errors.New("PostWatchlistSymbols() not implemented.")
}

//DELETE watchlists/{id}/symbols
func (client *Client)DeleteWatchlistSymbols(id int) (*AccountsBalancesResponse, error) {
	return nil, errors.New("DeleteWatchlistSymbols() not implemented.")
}
