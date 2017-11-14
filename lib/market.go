package lib

import (
	"github.com/jackmanlabs/errors"
	"net/http"
)

//GET market/clock
func GetClock(client *http.Client) (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetClock() not implemented.")
}

//GET market/ext/quotes
func GetExtQuotes(client *http.Client) (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetExtQuotes() not implemented.")
}

//GET market/news/search
func GetNewsSearch(client *http.Client) (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetNewsSearch() not implemented.")
}

//GET market/news/{id}
func GetNews(client *http.Client, id int) (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetNews() not implemented.")
}

//GET market/options/search
func GetOptionsSearch(client *http.Client) (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetOptionsSearch() not implemented.")
}

//GET market/options/strikes
func GetOptionsStrikes(client *http.Client) (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetOptionsStrikes() not implemented.")
}

//GET market/options/expirations
func GetOptionsExpirations(client *http.Client) (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetOptionsExpirations() not implemented.")
}

//GET market/timesales
func GetTimeSales(client *http.Client) (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetTimeSales() not implemented.")
}

//GET market/toplists
func GetTopLists(client *http.Client) (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetTopLists() not implemented.")
}

//GET market/quotes (streaming)
func GetQuotes(client *http.Client) (*AccountsBalancesResponse, error) {
	return nil, errors.New("GetQuotes() not implemented.")
}
