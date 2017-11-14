package lib

import (
	"github.com/jackmanlabs/errors"
)

//GET market/clock
func (client *Client) GetClock() (*ClockResponse, error) {
	return nil, errors.New("GetClock() not implemented.")
}

//GET market/ext/quotes
func (client *Client) GetExtQuotes() (*ExtQuotesResponse, error) {
	return nil, errors.New("GetExtQuotes() not implemented.")
}

//GET market/news/search
func (client *Client) GetNewsSearch() (*NewsSearchResponse, error) {
	return nil, errors.New("GetNewsSearch() not implemented.")
}

//GET market/news/{id}
func (client *Client) GetNews(id int) (*NewsResponse, error) {
	return nil, errors.New("GetNews() not implemented.")
}

//GET market/options/search
func (client *Client) GetOptionsSearch() (*OptionsSearchResponse, error) {
	return nil, errors.New("GetOptionsSearch() not implemented.")
}

//GET market/options/strikes
func (client *Client) GetOptionsStrikes() (*OptionsStrikesResponse, error) {
	return nil, errors.New("GetOptionsStrikes() not implemented.")
}

//GET market/options/expirations
func (client *Client) GetOptionsExpirations() (*OptionsExpirationsResponse, error) {
	return nil, errors.New("GetOptionsExpirations() not implemented.")
}

//GET market/timesales
func (client *Client) GetTimeSales() (*TimeSalesResponse, error) {
	return nil, errors.New("GetTimeSales() not implemented.")
}

//GET market/toplists
func (client *Client) GetTopLists() (*TopListsResponse, error) {
	return nil, errors.New("GetTopLists() not implemented.")
}

//GET market/quotes (streaming)
func (client *Client) GetQuotes() (*QuotesResponse, error) {
	return nil, errors.New("GetQuotes() not implemented.")
}

type QuotesResponse struct{}
type TopListsResponse struct{}
type TimeSalesResponse struct{}
type OptionsExpirationsResponse struct{}
type OptionsStrikesResponse struct{}
type OptionsSearchResponse struct{}
type NewsResponse struct{}
type NewsSearchResponse struct{}
type ExtQuotesResponse struct{}
type ClockResponse struct{}
