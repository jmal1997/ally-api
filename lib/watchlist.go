package lib

import (
	"fmt"
	"github.com/jackmanlabs/errors"
)

//GET watchlists
func (client *Client) GetWatchlists() (*WatchListsResponse, error) {
	var (
		path   string              = "/watchlists.xml"
		target *WatchListsResponse = new(WatchListsResponse)
	)

	err := client.get(path, target)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return target, nil
}

//POST watchlists
func (client *Client) PostWatchlists() (*AccountsBalancesResponse, error) {
	return nil, errors.New("PostWatchlists() not implemented.")
}

//GET watchlists/{id}
func (client *Client) GetWatchlist(id int) (*WatchListResponse, error) {
	var (
		path   string             = fmt.Sprintf("/watchlists/%d.xml", id)
		target *WatchListResponse = new(WatchListResponse)
	)

	err := client.get(path, target)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return target, nil
}

//DELETE watchlists/{id}
func (client *Client) DeleteWatchlist(id int) (*AccountsBalancesResponse, error) {
	return nil, errors.New("DeleteWatchlist() not implemented.")
}

//POST watchlists/{id}/symbols
func (client *Client) PostWatchlistSymbols(id int) (*AccountsBalancesResponse, error) {
	return nil, errors.New("PostWatchlistSymbols() not implemented.")
}

//DELETE watchlists/{id}/symbols
func (client *Client) DeleteWatchlistSymbols(id int) (*AccountsBalancesResponse, error) {
	return nil, errors.New("DeleteWatchlistSymbols() not implemented.")
}

type WatchListResponse struct{ Response }
type WatchListsResponse struct{ Response }
