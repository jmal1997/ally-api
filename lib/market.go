package lib

import (
	"github.com/jackmanlabs/errors"
)

//GET market/clock
func (client *Client)GetClock() (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetClock() not implemented.")
}

//GET market/ext/quotes
func (client *Client)GetExtQuotes() (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetExtQuotes() not implemented.")
}

//GET market/news/search
func (client *Client)GetNewsSearch() (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetNewsSearch() not implemented.")
}

//GET market/news/{id}
func (client *Client)GetNews(id int) (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetNews() not implemented.")
}

//GET market/options/search
func (client *Client)GetOptionsSearch() (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetOptionsSearch() not implemented.")
}

//GET market/options/strikes
func (client *Client)GetOptionsStrikes() (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetOptionsStrikes() not implemented.")
}

//GET market/options/expirations
func (client *Client)GetOptionsExpirations() (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetOptionsExpirations() not implemented.")
}

//GET market/timesales
func (client *Client)GetTimeSales() (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetTimeSales() not implemented.")
}

//GET market/toplists
func (client *Client)GetTopLists() (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetTopLists() not implemented.")
}

//GET market/quotes (streaming)
func (client *Client)GetQuotes() (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetQuotes() not implemented.")
}
