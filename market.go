package ally_api

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

type ClockResponse struct {
	ProfileResponse
}

type Quote struct {
	Adp100           *float64  `xml:"adp_100"`            // Stock,	Average Daily Price - 100 day
	Adp200           *float64  `xml:"adp_200"`            // Stock,	Average Daily Price - 200 day
	Adp50            *float64  `xml:"adp_50"`             // Stock,	Average Daily Price - 50 day
	Adv21            *int      `xml:"adv_21"`             // Stock,	Average Daily Volume - 21 day
	Adv30            *int      `xml:"adv_30"`             // Stock,	Average Daily Volume - 30 day
	Adv90            *int      `xml:"adv_90"`             // Stock,	Average Daily Volume - 90 day
	Ask              *float64  `xml:"ask"`                // Stock, Option	Ask price
	AskTime          *string   `xml:"ask_time"`           // Stock, Option	Time of latest ask
	Asksz            *int      `xml:"asksz"`              // Stock, Option	Size of latest ask (in 100's)
	Basis            *int      `xml:"basis"`              // Stock, Option	Reported precision (quotation decimal places)
	Beta             *float64  `xml:"beta"`               // Stock,	Beta volatility measure
	Bid              *float64  `xml:"bid"`                // Stock, Option	Bid price
	BidTime          *string   `xml:"bid_time"`           // Stock, Option	Time of latest bid
	Bidsz            *int      `xml:"bidsz"`              // Stock, Option	Size of latest bid (in 100's)
	Bidtick          *string   `xml:"bidtick"`            // Stock,	Tick direction since last bid
	Chg              *float64  `xml:"chg"`                // Stock, Option	Change since prior day close (cl)
	ChgSign          *string   `xml:"chg_sign"`           // Stock, Option	Change sign (e, u, d) as even, up, down
	ChgT             *float64  `xml:"chg_t"`              // Stock, Option	change in text format
	Cl               *float64  `xml:"cl"`                 // Stock, Option	previous close
	ContractSize     *int      `xml:"contract_size"`      // Option,	contract size for option
	Cusip            *int      `xml:"cusip"`              // Stock,	Cusip
	Date             *string   `xml:"date"`               // Stock, Option	Trade date of last trade
	Datetime         time.Time `xml:"datetime"`           // Stock, Option	Date and time
	DaysToExpiration *int      `xml:"days_to_expiration"` // Option,	Days until option expiration date
	Div              *float64  `xml:"div"`                // Stock,	Latest announced cash dividend
	Divexdate        *string   `xml:"divexdate"`          // Stock,	Ex-dividend date of div(YYYYMMDD)
	Divfreq          *string   `xml:"divfreq"`            // Stock,	Dividend frequency, A - Annual Dividend, S - Semi Annual Dividend, Q - Quarterly Dividend, M - Monthly Dividend, N - Not applicable or No Set Dividend Frequency.
	Divpaydt         *string   `xml:"divpaydt"`           // Stock,	Dividend pay date of last announced div
	DollarValue      *float64  `xml:"dollar_value"`       // Stock, Option	Total dollar value of shares traded today
	Eps              *float64  `xml:"eps"`                // Stock,	Earnings per share
	Exch             *string   `xml:"exch"`               // Stock, Option	exchange code
	ExchDesc         *string   `xml:"exch_desc"`          // Stock, Option	exchange description
	Hi               *float64  `xml:"hi"`                 // Stock, Option	High Trade Price for the trading day
	Iad              *float64  `xml:"iad"`                // Stock,	Indicated annual dividend
	Idelta           *int      `xml:"idelta"`             // Option,	Option risk measure of delta using implied volatility
	Igamma           *int      `xml:"igamma"`             // Option,	Option risk measure of gamma using implied volatility
	ImpVolatility    *int      `xml:"imp_volatility"`     // Option,	Implied volatility of option price based current stock price
	IncrVl           *int      `xml:"incr_vl"`            // Stock, Option	Volume of last trade
	Irho             *int      `xml:"irho"`               // Option,	Option risk measure of rho using implied volatility
	IssueDesc        *int      `xml:"issue_desc"`         // Option,	Issue description
	Itheta           *int      `xml:"itheta"`             // Option,	Option risk measure of theta using implied volatility
	Ivega            *int      `xml:"ivega"`              // Option,	Option risk measure of vega using implied volatility
	Last             *float64  `xml:"last"`               // Stock, Option	Last trade price
	Lo               *float64  `xml:"lo"`                 // Stock, Option	Low Trade Price for the trading day
	Name             *string   `xml:"name"`               // Stock, Option	Company name
	OpDelivery       *int      `xml:"op_delivery"`        // Option,	Option Settlement Designation – S Std N – Non Std X - NA
	OpFlag           *int      `xml:"op_flag"`            // Stock,	Security has options (1=Yes, 0=No).
	OpStyle          *int      `xml:"op_style"`           // Option,	Option Style – values are “A” American and “E” European
	OpSubclass       *int      `xml:"op_subclass"`        // Option,	Option class (0=Standard, 1=Leap, 3=Short Term)
	Openinterest     *int      `xml:"openinterest"`       // Option,	Open interest of option contract
	Opn              *float64  `xml:"opn"`                // Stock, Option	Open trade price
	OptVal           *int      `xml:"opt_val"`            // Option,	Estimated Option Value – via Ju/Zhong or Black-Scholes
	Pchg             *float64  `xml:"pchg"`               // Stock, Option	percentage change from prior day close
	PchgSign         *string   `xml:"pchg_sign"`          // Stock, Option	prchg sign (e, u, d) as even, up, down
	Pcls             *float64  `xml:"pcls"`               // Stock, Option	Prior day close
	Pe               *float64  `xml:"pe"`                 // Stock,	Price earnings ratio
	Phi              *float64  `xml:"phi"`                // Stock, Option	Prior day high value
	Plo              *float64  `xml:"plo"`                // Stock, Option	Prior day low value
	Popn             *float64  `xml:"popn"`               // Stock, Option	Prior day open
	PrAdp100         *float64  `xml:"pr_adp_100"`         // Stock,	Prior Average Daily Price "100"trade days
	PrAdp200         *float64  `xml:"pr_adp_200"`         // Stock,	Prior Average Daily Price "200"trade days
	PrAdp50          *float64  `xml:"pr_adp_50"`          // Stock,	Prior Average Daily Price "50"trade days
	PrDate           *string   `xml:"pr_date"`            // Stock, Option	Trade Date of Prior Last
	PrOpeninterest   *int      `xml:"pr_openinterest"`    // Option,	Prior day's open interest
	Prbook           *float64  `xml:"prbook"`             // Stock,	Book Value Price
	Prchg            *float64  `xml:"prchg"`              // Stock, Option	Prior day change
	PremMult         *int      `xml:"prem_mult"`          // Option,	Option premium multiplier
	PutCall          *int      `xml:"put_call"`           // Option,	Option type (Put or Call)
	Pvol             *int      `xml:"pvol"`               // Stock, Option	Prior day total volume
	Qcond            *int      `xml:"qcond"`              // Option,	Condition code of quote
	Rootsymbol       *int      `xml:"rootsymbol"`         // Option,	Option root symbol
	Secclass         *int      `xml:"secclass"`           // Stock, Option	Security class (0=stock, 1=option)
	Sesn             *string   `xml:"sesn"`               // Stock, Option	Trading session as (pre, regular, &amp, post)
	Sho              *int      `xml:"sho"`                // Stock,	Shares Outstanding
	Strikeprice      *int      `xml:"strikeprice"`        // Option,	Option strike price (not extended by multiplier)
	Symbol           *string   `xml:"symbol"`             // Stock, Option	Symbol from data provider
	Tcond            *string   `xml:"tcond"`              // Stock, Option	Trade condition code – (H) halted or (R) resumed
	Timestamp        *int      `xml:"timestamp"`          // Stock, Option	Timestamp
	TrNum            *int      `xml:"tr_num"`             // Stock, Option	Number of trades since market open
	Tradetick        *string   `xml:"tradetick"`          // Stock, Option	Tick direction from prior trade – (e,u,d) even, up, down)
	Trend            *string   `xml:"trend"`              // Stock, Option	Trend based on 10 prior ticks (e,u,d) even, up, down
	UnderCusip       *int      `xml:"under_cusip"`        // Option,	An option's underlying cusip
	Undersymbol      *int      `xml:"undersymbol"`        // Option,	An option's underlying symbol
	Vl               *int      `xml:"vl"`                 // Stock, Option	Cumulative volume
	Volatility12     *float64  `xml:"volatility12"`       // Stock,	one year volatility measure
	Vwap             *float64  `xml:"vwap"`               // Stock, Option	Volume weighted average price
	Wk52Hi           *float64  `xml:"wk52hi"`             // Stock, Option	52 week high
	Wk52Hidate       *string   `xml:"wk52hidate"`         // Stock, Option	52 week high date
	Wk52Lo           *float64  `xml:"wk52lo"`             // Stock, Option	52 week low
	Wk52Lodate       *string   `xml:"wk52lodate"`         // Stock, Option	52 week low date
	Xdate            *int      `xml:"xdate"`              // Option,	Expiration date of option in the format of (YYYYMMDD)
	Xday             *int      `xml:"xday"`               // Option,	Expiration day of option
	Xmonth           *int      `xml:"xmonth"`             // Option,	Expiration month of option
	Xyear            *int      `xml:"xyear"`              // Option,	Expiration year of option
	Yield            *float64  `xml:"yield"`              // Stock,	Dividend yield as %
}

/*
I know this is ugly, BUT this is a real PITA situation. I don't want to hassle
the consumers of this API with the issue of nested values and what-have-you.
This solution is messy, but it makes it easy for the consumer.
*/
func (this *Quote) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var endElement xml.EndElement
	for endElement.Name.Local != start.Name.Local {

		t, err := d.Token()
		if err != nil {
			return errors.Stack(err)
		}

		if t == nil {
			break
		}

		switch el := t.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "adp_100":
				this.Adp100, err = naParseFloat(d, &el)
			case "adp_200":
				this.Adp200, err = naParseFloat(d, &el)
			case "adp_50":
				this.Adp50, err = naParseFloat(d, &el)
			case "adv_21":
				this.Adv21, err = naParseInt(d, &el)
			case "adv_30":
				this.Adv30, err = naParseInt(d, &el)
			case "adv_90":
				this.Adv90, err = naParseInt(d, &el)
			case "ask":
				this.Ask, err = naParseFloat(d, &el)
			case "ask_time":
				this.AskTime, err = naParseString(d, &el)
			case "asksz":
				this.Asksz, err = naParseInt(d, &el)
			case "basis":
				this.Basis, err = naParseInt(d, &el)
			case "beta":
				this.Beta, err = naParseFloat(d, &el)
			case "bid":
				this.Bid, err = naParseFloat(d, &el)
			case "bid_time":
				this.BidTime, err = naParseString(d, &el)
			case "bidsz":
				this.Bidsz, err = naParseInt(d, &el)
			case "bidtick":
				this.Bidtick, err = naParseString(d, &el)
			case "chg":
				this.Chg, err = naParseFloat(d, &el)
			case "chg_sign":
				this.ChgSign, err = naParseString(d, &el)
			case "chg_t":
				this.ChgT, err = naParseFloat(d, &el)
			case "cl":
				this.Cl, err = naParseFloat(d, &el)
			case "contract_size":
				this.ContractSize, err = naParseInt(d, &el)
			case "cusip":
				this.Cusip, err = naParseInt(d, &el)
			case "date":
				this.Date, err = naParseString(d, &el)
			case "datetime":
				err = d.DecodeElement(&this.Datetime, &el)
			case "days_to_expiration":
				this.DaysToExpiration, err = naParseInt(d, &el)
			case "div":
				this.Div, err = naParseFloat(d, &el)
			case "divexdate":
				this.Divexdate, err = naParseString(d, &el)
			case "divfreq":
				this.Divfreq, err = naParseString(d, &el)
			case "divpaydt":
				this.Divpaydt, err = naParseString(d, &el)
			case "dollar_value":
				this.DollarValue, err = naParseFloat(d, &el)
			case "eps":
				this.Eps, err = naParseFloat(d, &el)
			case "exch":
				this.Exch, err = naParseString(d, &el)
			case "exch_desc":
				this.ExchDesc, err = naParseString(d, &el)
			case "hi":
				this.Hi, err = naParseFloat(d, &el)
			case "iad":
				this.Iad, err = naParseFloat(d, &el)
			case "idelta":
				this.Idelta, err = naParseInt(d, &el)
			case "igamma":
				this.Igamma, err = naParseInt(d, &el)
			case "imp_volatility":
				this.ImpVolatility, err = naParseInt(d, &el)
			case "incr_vl":
				this.IncrVl, err = naParseInt(d, &el)
			case "irho":
				this.Irho, err = naParseInt(d, &el)
			case "issue_desc":
				this.IssueDesc, err = naParseInt(d, &el)
			case "itheta":
				this.Itheta, err = naParseInt(d, &el)
			case "ivega":
				this.Ivega, err = naParseInt(d, &el)
			case "last":
				this.Last, err = naParseFloat(d, &el)
			case "lo":
				this.Lo, err = naParseFloat(d, &el)
			case "name":
				this.Name, err = naParseString(d, &el)
			case "op_delivery":
				this.OpDelivery, err = naParseInt(d, &el)
			case "op_flag":
				this.OpFlag, err = naParseInt(d, &el)
			case "op_style":
				this.OpStyle, err = naParseInt(d, &el)
			case "op_subclass":
				this.OpSubclass, err = naParseInt(d, &el)
			case "openinterest":
				this.Openinterest, err = naParseInt(d, &el)
			case "opn":
				this.Opn, err = naParseFloat(d, &el)
			case "opt_val":
				this.OptVal, err = naParseInt(d, &el)
			case "pchg":
				this.Pchg, err = naParseFloat(d, &el)
			case "pchg_sign":
				this.PchgSign, err = naParseString(d, &el)
			case "pcls":
				this.Pcls, err = naParseFloat(d, &el)
			case "pe":
				this.Pe, err = naParseFloat(d, &el)
			case "phi":
				this.Phi, err = naParseFloat(d, &el)
			case "plo":
				this.Plo, err = naParseFloat(d, &el)
			case "popn":
				this.Popn, err = naParseFloat(d, &el)
			case "pr_adp_100":
				this.PrAdp100, err = naParseFloat(d, &el)
			case "pr_adp_200":
				this.PrAdp200, err = naParseFloat(d, &el)
			case "pr_adp_50":
				this.PrAdp50, err = naParseFloat(d, &el)
			case "pr_date":
				this.PrDate, err = naParseString(d, &el)
			case "pr_openinterest":
				this.PrOpeninterest, err = naParseInt(d, &el)
			case "prbook":
				this.Prbook, err = naParseFloat(d, &el)
			case "prchg":
				this.Prchg, err = naParseFloat(d, &el)
			case "prem_mult":
				this.PremMult, err = naParseInt(d, &el)
			case "put_call":
				this.PutCall, err = naParseInt(d, &el)
			case "pvol":
				this.Pvol, err = naParseInt(d, &el)
			case "qcond":
				this.Qcond, err = naParseInt(d, &el)
			case "rootsymbol":
				this.Rootsymbol, err = naParseInt(d, &el)
			case "secclass":
				this.Secclass, err = naParseInt(d, &el)
			case "sesn":
				this.Sesn, err = naParseString(d, &el)
			case "sho":
				this.Sho, err = naParseInt(d, &el)
			case "strikeprice":
				this.Strikeprice, err = naParseInt(d, &el)
			case "symbol":
				this.Symbol, err = naParseString(d, &el)
			case "tcond":
				this.Tcond, err = naParseString(d, &el)
			case "timestamp":
				this.Timestamp, err = naParseInt(d, &el)
			case "tr_num":
				this.TrNum, err = naParseInt(d, &el)
			case "tradetick":
				this.Tradetick, err = naParseString(d, &el)
			case "trend":
				this.Trend, err = naParseString(d, &el)
			case "under_cusip":
				this.UnderCusip, err = naParseInt(d, &el)
			case "undersymbol":
				this.Undersymbol, err = naParseInt(d, &el)
			case "vl":
				this.Vl, err = naParseInt(d, &el)
			case "volatility12":
				this.Volatility12, err = naParseFloat(d, &el)
			case "vwap":
				this.Vwap, err = naParseFloat(d, &el)
			case "wk52hi":
				this.Wk52Hi, err = naParseFloat(d, &el)
			case "wk52hidate":
				this.Wk52Hidate, err = naParseString(d, &el)
			case "wk52lo":
				this.Wk52Lo, err = naParseFloat(d, &el)
			case "wk52lodate":
				this.Wk52Lodate, err = naParseString(d, &el)
			case "xdate":
				this.Xdate, err = naParseInt(d, &el)
			case "xday":
				this.Xday, err = naParseInt(d, &el)
			case "xmonth":
				this.Xmonth, err = naParseInt(d, &el)
			case "xyear":
				this.Xyear, err = naParseInt(d, &el)
			case "yield":
				this.Yield, err = naParseFloat(d, &el)
			}
			if err != nil {
				return errors.Stack(err)
			}
		case xml.EndElement:
			endElement = el
		}
	}

	return nil
}

func naParseInt(d *xml.Decoder, startElement *xml.StartElement) (*int, error) {
	var s string
	err := d.DecodeElement(&s, startElement)
	if err != nil {
		return nil, errors.Stack(err)
	}

	if s == "na" {
		return nil, nil
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return &i, nil
}

func naParseFloat(d *xml.Decoder, startElement *xml.StartElement) (*float64, error) {
	var s string
	err := d.DecodeElement(&s, startElement)
	if err != nil {
		return nil, errors.Stack(err)
	}

	if s == "na" {
		return nil, nil
	}

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return &f, nil
}

func naParseString(d *xml.Decoder, startElement *xml.StartElement) (*string, error) {
	var s string
	err := d.DecodeElement(&s, startElement)
	if err != nil {
		return nil, errors.Stack(err)
	}

	if s == "na" {
		return nil, nil
	}

	return &s, nil
}
