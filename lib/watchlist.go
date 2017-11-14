package lib

import (
	"github.com/jackmanlabs/errors"
	"net/http"
)

//GET watchlists
func GetWatchlists(client *http.Client) (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetWatchlists() not implemented.")
}

//POST watchlists
func PostWatchlists(client *http.Client) (*AccountsBalancesResponse, error) {
	return nil, errors.New("PostWatchlists() not implemented.")
}

//GET watchlists/{id}
func GetWatchlist(client *http.Client, id int) (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetWatchlist() not implemented.")
}

//DELETE watchlists/{id}
func DeleteWatchlist(client *http.Client, id int) (*AccountsBalancesResponse, error) {
	return nil, errors.New("DeleteWatchlist() not implemented.")
}

//POST watchlists/{id}/symbols
func PostWatchlistSymbols(client *http.Client, id int) (*AccountsBalancesResponse, error) {
	return nil, errors.New("PostWatchlistSymbols() not implemented.")
}

//DELETE watchlists/{id}/symbols
func DeleteWatchlistSymbols(client *http.Client, id int) (*AccountsBalancesResponse, error) {
	return nil, errors.New("DeleteWatchlistSymbols() not implemented.")
}
