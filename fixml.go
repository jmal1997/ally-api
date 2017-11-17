package ally_api

import "encoding/xml"

type Fixml struct {
	XMLName xml.Name   `xml:"FIXML"`
	Order   FixmlOrder `xml:"Order"`

	// Only used for closing short positions, "Buy to Cover" orders should
	// include this attribute as AcctTyp = "5".
	AcctTyp string `xml:"AcctTyp"`

	// Abbreviation for "classification of financial instrument", used for
	// options to distinguish "OC" for call option or "OP" for put option.
	CFI string `xml:"CFI"`

	// Used for trailing stop orders. Value of ExecInst = "a" needs to be
	// passed.
	ExecInst string `xml:"ExecInst"`

	// Expiration of the option in the format of YYYYMM.
	MMY string `xml:"MMY"`

	// Represents the expiration date of a option. Needs to be in the format of
	// "YYYY‐MM‐ DDT00:00:00.000‐05:00". For single leg orders, this attribute
	// tag changes from Mat to MatDt.
	Mat string `xml:"Mat"`

	// Represents the expiration date of a option. Needs to be in the format of
	// "YYYY‐MM‐ DDT00:00:00.000‐05:00". For multiple leg orders, this attribute
	// tag changes from MatDt to Mat.
	MatDt string `xml:"MatDt"`

	// Used for trailing stop orders. Value of OfstTyp = "0" needs to be passed.
	// The offset value of "0" denotes a "price" offset from the PegPxTyp field
	// below. The offset value of "1" denotes a "basis point" offset from the
	// PegPxTyp field below (used as a percentage offset).
	OfstTyp string `xml:"OfstTyp"`

	// Used for trailing stop orders. Signed value needs to be passed for amount
	// of offset value combined with the PegPxTyp & OfstTyp fields. Negative
	// values are normally used for sell trailing stops so the trigger trails
	// below current price. Positive values are normally used for buy trailing
	// stops so the trigger trails above the current price. For example,
	// assuming an OfstTyp = "0", a sell order with a OfstVal of ‐.50 will
	// trigger if the current price falls by more than .50 of its last highest
	// value since the order was placed. OfstType = "1" would require the signed
	// value for a percentage. For example, OfstVal = "5" would represent a 5%
	// increase in price before a buy trailing stop is triggered.
	OfstVal string `xml:"OfstVal"`

	// Order ID that needs to be passed for any change or cancel requests. Note:
	// for Multi‐leg orders, use tag OrigClOrdID instead of OrigID.
	OrigID string `xml:"OrigID"`

	// Used for trailing stop orders defining type of peg (price used) for
	// trailing. In this case, PegPxTyp = "1" references "last price" of
	// security.
	PegPxTyp string `xml:"PegPxTyp"`

	// Used for options, option legs require and attribute of "O" for opening or
	// "C" for closing.
	PosEfct string `xml:"PosEfct"`

	// Price for price type if needed. This attribute would be required for
	// limits (Typ = "2") or stop limits (Typ = "4").
	Px string `xml:"Px"`

	// Strike price of option contract. This tag changes from Strk to StrkPx for
	// single leg orders.
	Strk string `xml:"Strk"`

	// Strike price of option contract. This tag changes from StrkPx to Strk for
	// multi‐leg orders.
	StrkPx string `xml:"StrkPx"`
}

type FixmlOrder struct {
	OrderQuantity FixmlOrderQuantity `xml:"OrdQty"`
	Instrument    FixmlInstrument    `xml:"Instrmt"`

	// Account number needs to be passed with all order requests.
	Acct int `xml:"Acct,attr"`

	// Time in force, possible values include "0" ‐ Day Order, "1" ‐ GTC Order,
	// "7" ‐ Market on Close. Not applicable when Typ = "1" (market order).
	TmInForce int `xml:"TmInForce,attr"`

	// Price Type as "1" ‐ Market, "2" ‐ Limit", "3" ‐ Stop, "4" Stop Limit, or
	// "P" for trailing stop.
	Typ string `xml:"Typ,attr"`

	// Side of market as "1" ‐ Buy, "2" ‐ Sell, "5" ‐ Sell Short. Buy to cover
	// orders are attributed as buy orders with Side = "1".
	Side int `xml:"Side,attr"`
}

type FixmlInstrument struct {
	XMLName xml.Name `xml:"Instrmt"`

	// Security type attribute is needed."CS" for common stock or "OPT" for
	// option.
	SecTyp string `xml:"SecTyp,attr"`

	// Ticker symbol of underlying security. This is utilized for stock, option,
	// & multi‐leg orders.
	Sym string `xml:"Sym,attr"`
}

type FixmlOrderQuantity struct {
	XMLName xml.Name `xml:"OrdQty"`
	Qty     int      `xml:"Qty,attr"`
}
