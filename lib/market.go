package lib

import (
	"encoding/xml"
	"fmt"
	"github.com/jackmanlabs/errors"
	"strconv"
	"strings"
	"time"
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

type ExtQuotesResponse struct {
	Response
	Quotes []Quote `xml:"quotes>quote"`
}

type ClockResponse struct{ Response }

type Quote struct {
	Adp100           float64   `xml:"adp_100"`            // Stock,	Average Daily Price - 100 day
	Adp200           float64   `xml:"adp_200"`            // Stock,	Average Daily Price - 200 day
	Adp50            float64   `xml:"adp_50"`             // Stock,	Average Daily Price - 50 day
	Adv21            int       `xml:"adv_21"`             // Stock,	Average Daily Volume - 21 day
	Adv30            int       `xml:"adv_30"`             // Stock,	Average Daily Volume - 30 day
	Adv90            int       `xml:"adv_90"`             // Stock,	Average Daily Volume - 90 day
	Ask              float64   `xml:"ask"`                // Stock, Option	Ask price
	AskTime          string    `xml:"ask_time"`           // Stock, Option	Time of latest ask
	Asksz            float64   `xml:"asksz"`              // Stock, Option	Size of latest ask (in 100's)
	Basis            string    `xml:"basis"`              // Stock, Option	Reported precision (quotation decimal places)
	Beta             float64   `xml:"beta"`               // Stock,	Beta volatility measure
	Bid              float64   `xml:"bid"`                // Stock, Option	Bid price
	BidTime          string    `xml:"bid_time"`           // Stock, Option	Time of latest bid
	Bidsz            float64   `xml:"bidsz"`              // Stock, Option	Size of latest bid (in 100's)
	Bidtick          string    `xml:"bidtick"`            // Stock,	Tick direction since last bid
	Chg              float64   `xml:"chg"`                // Stock, Option	Change since prior day close (cl)
	ChgSign          string    `xml:"chg_sign"`           // Stock, Option	Change sign (e, u, d) as even, up, down
	ChgT             string    `xml:"chg_t"`              // Stock, Option	change in text format
	Cl               float64   `xml:"cl"`                 // Stock, Option	previous close
	ContractSize     string    `xml:"contract_size"`      // Option,	contract size for option
	Cusip            string    `xml:"cusip"`              // Stock,	Cusip
	Date             string    `xml:"date"`               // Stock, Option	Trade date of last trade
	Datetime         time.Time `xml:"datetime"`           // Stock, Option	Date and time
	DaysToExpiration string    `xml:"days_to_expiration"` // Option,	Days until option expiration date
	Div              float64   `xml:"div"`                // Stock,	Latest announced cash dividend
	Divexdate        string    `xml:"divexdate"`          // Stock,	Ex-dividend date of div(YYYYMMDD)
	Divfreq          string    `xml:"divfreq"`            // Stock,	Dividend frequency, A - Annual Dividend, S - Semi Annual Dividend, Q - Quarterly Dividend, M - Monthly Dividend, N - Not applicable or No Set Dividend Frequency.
	Divpaydt         string    `xml:"divpaydt"`           // Stock,	Dividend pay date of last announced div
	DollarValue      float64   `xml:"dollar_value"`       // Stock, Option	Total dollar value of shares traded today
	Eps              float64   `xml:"eps"`                // Stock,	Earnings per share
	Exch             string    `xml:"exch"`               // Stock, Option	exchange code
	ExchDesc         string    `xml:"exch_desc"`          // Stock, Option	exchange description
	Hi               float64   `xml:"hi"`                 // Stock, Option	High Trade Price for the trading day
	Iad              float64   `xml:"iad"`                // Stock,	Indicated annual dividend
	Idelta           string    `xml:"idelta"`             // Option,	Option risk measure of delta using implied volatility
	Igamma           string    `xml:"igamma"`             // Option,	Option risk measure of gamma using implied volatility
	ImpVolatility    string    `xml:"imp_volatility"`     // Option,	Implied volatility of option price based current stock price
	IncrVl           float64   `xml:"incr_vl"`            // Stock, Option	Volume of last trade
	Irho             string    `xml:"irho"`               // Option,	Option risk measure of rho using implied volatility
	IssueDesc        string    `xml:"issue_desc"`         // Option,	Issue description
	Itheta           string    `xml:"itheta"`             // Option,	Option risk measure of theta using implied volatility
	Ivega            string    `xml:"ivega"`              // Option,	Option risk measure of vega using implied volatility
	Last             float64   `xml:"last"`               // Stock, Option	Last trade price
	Lo               float64   `xml:"lo"`                 // Stock, Option	Low Trade Price for the trading day
	Name             string    `xml:"name"`               // Stock, Option	Company name
	OpDelivery       string    `xml:"op_delivery"`        // Option,	Option Settlement Designation – S Std N – Non Std X - NA
	OpFlag           float64   `xml:"op_flag"`            // Stock,	Security has options (1=Yes, 0=No).
	OpStyle          string    `xml:"op_style"`           // Option,	Option Style – values are “A” American and “E” European
	OpSubclass       string    `xml:"op_subclass"`        // Option,	Option class (0=Standard, 1=Leap, 3=Short Term)
	Openinterest     string    `xml:"openinterest"`       // Option,	Open interest of option contract
	Opn              float64   `xml:"opn"`                // Stock, Option	Open trade price
	OptVal           string    `xml:"opt_val"`            // Option,	Estimated Option Value – via Ju/Zhong or Black-Scholes
	Pchg             float64   `xml:"pchg"`               // Stock, Option	percentage change from prior day close
	PchgSign         string    `xml:"pchg_sign"`          // Stock, Option	prchg sign (e, u, d) as even, up, down
	Pcls             float64   `xml:"pcls"`               // Stock, Option	Prior day close
	Pe               float64   `xml:"pe"`                 // Stock,	Price earnings ratio
	Phi              float64   `xml:"phi"`                // Stock, Option	Prior day high value
	Plo              float64   `xml:"plo"`                // Stock, Option	Prior day low value
	Popn             float64   `xml:"popn"`               // Stock, Option	Prior day open
	PrAdp100         float64   `xml:"pr_adp_100"`         // Stock,	Prior Average Daily Price "100"trade days
	PrAdp200         float64   `xml:"pr_adp_200"`         // Stock,	Prior Average Daily Price "200"trade days
	PrAdp50          float64   `xml:"pr_adp_50"`          // Stock,	Prior Average Daily Price "50"trade days
	PrDate           string    `xml:"pr_date"`            // Stock, Option	Trade Date of Prior Last
	PrOpeninterest   string    `xml:"pr_openinterest"`    // Option,	Prior day's open interest
	Prbook           float64   `xml:"prbook"`             // Stock,	Book Value Price
	Prchg            float64   `xml:"prchg"`              // Stock, Option	Prior day change
	PremMult         string    `xml:"prem_mult"`          // Option,	Option premium multiplier
	PutCall          string    `xml:"put_call"`           // Option,	Option type (Put or Call)
	Pvol             float64   `xml:"pvol"`               // Stock, Option	Prior day total volume
	Qcond            float64   `xml:"qcond"`              // Option,	Condition code of quote
	Rootsymbol       string    `xml:"rootsymbol"`         // Option,	Option root symbol
	Secclass         float64   `xml:"secclass"`           // Stock, Option	Security class (0=stock, 1=option)
	Sesn             string    `xml:"sesn"`               // Stock, Option	Trading session as (pre, regular, &amp, post)
	Sho              int       `xml:"sho"`                // Stock,	Shares Outstanding
	Strikeprice      string    `xml:"strikeprice"`        // Option,	Option strike price (not extended by multiplier)
	Symbol           string    `xml:"symbol"`             // Stock, Option	Symbol from data provider
	Tcond            float64   `xml:"tcond"`              // Stock, Option	Trade condition code – (H) halted or (R) resumed
	Timestamp        int64     `xml:"timestamp"`          // Stock, Option	Timestamp
	TrNum            float64   `xml:"tr_num"`             // Stock, Option	Number of trades since market open
	Tradetick        string    `xml:"tradetick"`          // Stock, Option	Tick direction from prior trade – (e,u,d) even, up, down)
	Trend            string    `xml:"trend"`              // Stock, Option	Trend based on 10 prior ticks (e,u,d) even, up, down
	UnderCusip       string    `xml:"under_cusip"`        // Option,	An option's underlying cusip
	Undersymbol      string    `xml:"undersymbol"`        // Option,	An option's underlying symbol
	Vl               int       `xml:"vl"`                 // Stock, Option	Cumulative volume
	Volatility12     float64   `xml:"volatility12"`       // Stock,	one year volatility measure
	Vwap             float64   `xml:"vwap"`               // Stock, Option	Volume weighted average price
	Wk52Hi           float64   `xml:"wk52hi"`             // Stock, Option	52 week high
	Wk52Hidate       string    `xml:"wk52hidate"`         // Stock, Option	52 week high date
	Wk52Lo           float64   `xml:"wk52lo"`             // Stock, Option	52 week low
	Wk52Lodate       string    `xml:"wk52lodate"`         // Stock, Option	52 week low date
	Xdate            string    `xml:"xdate"`              // Option,	Expiration date of option in the format of (YYYYMMDD)
	Xday             string    `xml:"xday"`               // Option,	Expiration day of option
	Xmonth           string    `xml:"xmonth"`             // Option,	Expiration month of option
	Xyear            string    `xml:"xyear"`              // Option,	Expiration year of option
	Yield            float64   `xml:"yield"`              // Stock,	Dividend yield as %
}

type HHMM time.Duration

func (this *HHMM) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return errors.Stack(err)
	}

	chunks := strings.Split(s, ":")
	if len(chunks) != 2 {
		return errors.New("Expected a value in format HH:MM, got: " + s)
	}

	h, err := strconv.Atoi(chunks[0])
	if err != nil {
		return errors.Stack(err)
	}

	m, err := strconv.Atoi(chunks[1])
	if err != nil {
		return errors.Stack(err)
	}

	*this = HHMM(time.Hour*time.Duration(h) + time.Minute*time.Duration(m))

	return nil
}
