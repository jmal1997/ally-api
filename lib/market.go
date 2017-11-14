package lib

import (
	"fmt"
	"github.com/jackmanlabs/errors"
	"strings"
)

//GET market/clock
func (client *Client) GetClock() (*ClockResponse, error) {
	var (
		path   string         = "/market/clock.xml"
		target *ClockResponse = new(ClockResponse)
	)

	err := client.get(path, target)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return target, nil
}

//GET market/ext/quotes
func (client *Client) GetExtQuotes(symbols []string) (*ExtQuotesResponse, error) {
	var (
		path   string             = fmt.Sprintf("/market/ext/quotes.xml?symbols=%s", strings.Join(symbols, ","))
		target *ExtQuotesResponse = new(ExtQuotesResponse)
	)

	err := client.get(path, target)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return target, nil
}

//GET market/news/search
func (client *Client) GetNewsSearch() (*NewsSearchResponse, error) {
	var (
		path   string              = "/market/news/search.xml"
		target *NewsSearchResponse = new(NewsSearchResponse)
	)

	err := client.get(path, target)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return target, nil
}

//GET market/news/{id}
func (client *Client) GetNews(id int) (*NewsResponse, error) {
	var (
		path   string        = fmt.Sprintf("/market/news/%d.xml", id)
		target *NewsResponse = new(NewsResponse)
	)

	err := client.get(path, target)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return target, nil
}

//GET market/options/search
func (client *Client) GetOptionsSearch() (*OptionsSearchResponse, error) {
	var (
		path   string                 = "/market/options/search.xml"
		target *OptionsSearchResponse = new(OptionsSearchResponse)
	)

	err := client.get(path, target)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return target, nil
}

//GET market/options/strikes
func (client *Client) GetOptionsStrikes() (*OptionsStrikesResponse, error) {
	var (
		path   string                  = "/market/options/strikes.xml"
		target *OptionsStrikesResponse = new(OptionsStrikesResponse)
	)

	err := client.get(path, target)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return target, nil
}

//GET market/options/expirations
func (client *Client) GetOptionsExpirations() (*OptionsExpirationsResponse, error) {
	var (
		path   string                      = "/market/options/expirations.xml"
		target *OptionsExpirationsResponse = new(OptionsExpirationsResponse)
	)

	err := client.get(path, target)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return target, nil
}

//GET market/timesales
func (client *Client) GetTimeSales() (*TimeSalesResponse, error) {
	var (
		path   string             = "/market/timesales.xml"
		target *TimeSalesResponse = new(TimeSalesResponse)
	)

	err := client.get(path, target)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return target, nil
}

//GET market/toplists
func (client *Client) GetTopLists() (*TopListsResponse, error) {
	var (
		path   string            = "/market/toplists.xml"
		target *TopListsResponse = new(TopListsResponse)
	)

	err := client.get(path, target)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return target, nil
}

//GET market/quotes (streaming)
func (client *Client) GetQuotes() (*QuotesResponse, error) {
	var (
		path   string          = "/market/quotes.xml"
		target *QuotesResponse = new(QuotesResponse)
	)

	err := client.get(path, target)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return target, nil
}

type QuotesResponse struct{ Response }
type TopListsResponse struct{ Response }
type TimeSalesResponse struct{ Response }
type OptionsExpirationsResponse struct{ Response }
type OptionsStrikesResponse struct{ Response }
type OptionsSearchResponse struct{ Response }
type NewsResponse struct{ Response }
type NewsSearchResponse struct{ Response }
type ExtQuotesResponse struct{ Response }
type ClockResponse struct{ Response }
