package lib

import (
	"fmt"
	"github.com/jackmanlabs/errors"
)

// 	GET accounts
func (client *Client) GetAccounts() (*AccountsResponse, error) {
	var (
		path     string            = "/accounts.xml"
		response *AccountsResponse = new(AccountsResponse)
	)

	err := client.get(root+path, response)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return response, nil
}

//GET accounts/balances
func (client *Client) GetAccountsBalances() (*AccountsBalancesResponse, error) {
	var (
		path     string                    = "/accounts/balances.xml"
		response *AccountsBalancesResponse = new(AccountsBalancesResponse)
	)

	err := client.get(root+path, response)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return response, nil
}

//	GET accounts/{id}
func (client *Client) GetAccount(id int) (*AccountResponse, error) {
	var (
		path     string           = fmt.Sprintf("/accounts/%d.xml", id)
		response *AccountResponse = new(AccountResponse)
	)

	err := client.get(root+path, response)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return response, nil
}

//GET accounts/{id}/balances
func (client *Client) GetAccountBalances(id int) (*AccountBalancesResponse, error) {
	var (
		path     string                   = fmt.Sprintf("/accounts/%d.xml", id)
		response *AccountBalancesResponse = new(AccountBalancesResponse)
	)

	err := client.get(root+path, response)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return response, nil
}

//GET accounts/{id}/history
func (client *Client) GetAccountHistory(id int) (*AccountHistoryResponse, error) {
	var (
		path     string                  = fmt.Sprintf("/accounts/%d.xml", id)
		response *AccountHistoryResponse = new(AccountHistoryResponse)
	)

	err := client.get(root+path, response)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return response, nil
}

//GET accounts/{id}/holdings
func (client *Client) GetAccountHoldings(id int) (*AccountHoldingsResponse, error) {
	var (
		path     string                   = fmt.Sprintf("/accounts/%d.xml", id)
		response *AccountHoldingsResponse = new(AccountHoldingsResponse)
	)

	err := client.get(root+path, response)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return response, nil
}

type AccountsResponse struct {
	AccountSummaries []AccountSummary `json:"accounts" xml:"accounts>accountsummary"`
	ResponseId       string           `json:"id" xml:"id,attr"`
	ElapsedTime      float64          `json:"elapsedtime" xml:"elapsedtime,omitempty"`
	Error            string           `json:"error" xml:"error,omitempty"`
}

type AccountsBalancesResponse struct {
	AccountBalances []AccountBalanceSimple `json:"accountbalance" xml:"accountbalance"`
	ResponseId      string                 `json:"id" xml:"id,attr" json:"id"`
	ElapsedTime     float64                `json:"elapsedtime" xml:"elapsedtime,omitempty"`
	Error           string                 `json:"error" xml:"error,omitempty"`
	TotalBalance    float64                `json:"totalbalance" xml:"totalbalance>accountvalue"`
}

type AccountBalancesResponse struct{}

type AccountResponse struct{}

type AccountHoldingsResponse struct{}

type AccountHistoryResponse struct{}

type AccountSummary struct {
	Account         int             `json:"account" xml:"account"` // Account number
	AccountBalance  AccountBalance  `json:"accountbalance" xml:"accountbalance"`
	AccountHoldings AccountHoldings `json:"accountholdings" xml:"accountholdings"`
}

type AccountBalance struct {
	Account      int                `json:"account" xml:"account"`
	AccountValue float64            `json:"accountvalue" xml:"accountvalue"` // Account value
	BuyingPower  AccountBuyingPower `json:"buyingpower" xml:"buyingpower"`   //
	FedCall      float64            `json:"fedcall" xml:"fedcall"`           // Value of any fed call on account
	HouseCall    float64            `json:"housecall" xml:"housecall"`       // Value of any house call
	Money        AccountMoney       `json:"money" xml:"money"`               //
	Securities   AccountSecurities  `json:"securities" xml:"securities"`     //
}

type AccountBalanceSimple struct {
	Account      int     `json:"account" xml:"account"`
	AccountName  string  `json:"accountname" xml:"accountname,omitempty"` // Used in the 'GET accounts/balances' response
	AccountValue float64 `json:"accountvalue" xml:"accountvalue"`         // Account value
}

type AccountMoney struct {
	AccruedInterest   float64 `json:"accruedinterest" xml:"accruedinterest"`     // Amount of any accrued interest on the account
	Cash              float64 `json:"cash" xml:"cash"`                           // cash
	CashAvailable     float64 `json:"cashavailable" xml:"cashavailable"`         // cash available
	MarginBalance     float64 `json:"marginbalance" xml:"marginbalance"`         // Margin balance (- indicates debit balance, + indicates credit balance)
	Mmf               float64 `json:"mmf" xml:"mmf"`                             // Money market fund
	Total             float64 `json:"total" xml:"total"`                         // Total cash balance. Path: /response/accounts/accountsummary/accountbalance/money/
	UnclearedDeposits float64 `json:"uncleareddeposits" xml:"uncleareddeposits"` // uncleared deposits
	UnsettledFunds    float64 `json:"unsettledfunds" xml:"unsettledfunds"`       // unsettle funds
	Yield             float64 `json:"yield" xml:"yield"`                         // Yield
}

type AccountSecurities struct {
	LongOptions  float64 `json:"longoptions" xml:"longoptions"`   // Long option market value
	LongStocks   float64 `json:"longstocks" xml:"longstocks"`     // Long stock market value
	Options      float64 `json:"options" xml:"options"`           // Total option market value. Path: /response/accounts/accountsummary/accountbalance/securities/
	ShortOptions float64 `json:"shortoptions" xml:"shortoptions"` // Short option market value
	ShortStocks  float64 `json:"shortstocks" xml:"shortstocks"`   // Short stock market value
	Stocks       float64 `json:"stocks" xml:"stocks"`             // Total stock market value
	Total        float64 `json:"total" xml:"total"`               // Total market value (stock & option). Path: /response/accounts/accountsummary/accountbalance/securities/
}

type AccountHoldings struct {
	Holdings        []AccountHolding           `json:"holding" xml:"holding"`
	TotalSecurities float64                    `json:"totalsecurities" xml:"totalsecurities"` // Total account market value
	DisplayData     AccountHoldingsDisplayData `json:"displaydata" xml:"displaydata"`
}

type AccountHoldingsDisplayData struct {
	TotalSecurities string `json:"totalsecurities" xml:"totalsecurities"` // Total account market value
}

type AccountBuyingPower struct {
	CashAvailableForWithdrawal float64 `json:"cashavailableforwithdrawal" xml:"cashavailableforwithdrawal,omitempty"` // Cash available for withdrawal (cash & margin accounts only, n/a for retirement accounts)
	DayTrading                 float64 `json:"daytrading" xml:"daytrading,omitempty"`                                 //
	EquityPercentage           float64 `json:"equitypercentage" xml:"equitypercentage,omitempty"`                     // Percentage of equity (margin accounts only)
	Options                    float64 `json:"options" xml:"options,omitempty"`                                       // Options buying power. Path: /response/accounts/accountsummary/accountbalance/buyingpower/
	SodDayTrading              float64 `json:"soddaytrading" xml:"soddaytrading,omitempty"`                           //
	SodOptions                 float64 `json:"sodoptions" xml:"sodoptions,omitempty"`                                 // Start of day options buying power
	SodStock                   float64 `json:"sodstock" xml:"sodstock,omitempty"`                                     // Start of day stock buying power
	Stock                      float64 `json:"stock" xml:"stock,omitempty"`                                           // Stock buying power
}

type AccountHolding struct {
	AccountType       int                       `json:"accounttype" xml:"accounttype"`             // Holdings attribute for where asset as held, "1"= cash, "2"= margin long, "5"=margin short.
	CostBasis         float64                   `json:"costbasis" xml:"costbasis"`                 // Holding cost basis
	DisplayData       AccountHoldingDisplayData `json:"displaydata" xml:"displaydata"`             //
	GainLoss          float64                   `json:"gainloss" xml:"gainloss"`                   // Holding gain/loss overall
	Instrument        AccountHoldingInstrument  `json:"instrument" xml:"instrument"`               //
	MarketValue       float64                   `json:"marketvalue" xml:"marketvalue"`             // Holding market value
	MarketValueChange float64                   `json:"marketvaluechange" xml:"marketvaluechange"` // Holding market value change
	Price             float64                   `json:"price" xml:"price"`                         // Instrument price
	PurchasePrice     float64                   `json:"purchaseprice" xml:"purchaseprice"`         // Holding purchase price
	Qty               float64                   `json:"qty" xml:"qty"`                             // Holding quantity
	Quote             AccountHoldingQuote       `json:"quote" xml:"quote"`
	Underlying        struct{}                  `json:"underlying" xml:"underlying"`
}

type AccountHoldingQuote struct {
	Change    float64 `json:"change" xml:"change"`       // Holding asset change for the day
	LastPrice float64 `json:"lastprice" xml:"lastprice"` // Instrument last price
}

type AccountHoldingDisplayData struct {
	AccountType       string `json:"accounttype" xml:"accounttype"`             // Holdings attribute for where asset as held, "1"= cash, "2"= margin long, "5"=margin short.
	AssetClass        string `json:"assetclass" xml:"assetclass"`               // Holding asset class type
	Change            string `json:"change" xml:"change"`                       // Holding asset change for the day
	CostBasis         string `json:"costbasis" xml:"costbasis"`                 // Holding cost basis
	Desc              string `json:"desc" xml:"desc"`                           // Instrument description
	LastPrice         string `json:"lastprice" xml:"lastprice"`                 // Instrument last price
	MarketValue       string `json:"marketvalue" xml:"marketvalue"`             // Holding market value
	MarketValueChange string `json:"marketvaluechange" xml:"marketvaluechange"` // Holding market value change
	Qty               string `json:"qty" xml:"qty"`                             // Holding quantity
	Symbol            string `json:"symbol" xml:"symbol"`                       // Holding underlying symbol
}

type AccountHoldingInstrument struct {
	Cusip        string  `json:"cusip" xml:"cusip"`   // Instrument cusip
	Desc         string  `json:"desc" xml:"desc"`     // Instrument description
	Factor       float64 `json:"factor" xml:"factor"` // Instrument factor
	SecurityType string  `json:"sectyp" xml:"sectyp"` // Instrument security type (FIXML)
	Symbol       string  `json:"sym" xml:"sym"`       // Instrument underlying symbol (FIXML)
}

type Account struct {
	CashBalance       string `json:"cashbalance" xml:"cashbalance"`   // cash balance
	CashMarketValue   string `json:"cashmv" xml:"cashmv"`             // Value of cash market value
	Cfi               string `json:"cfi" xml:"cfi"`                   // Put or call code(FIXML)
	MarginMarketValue string `json:"marginmv" xml:"marginmv"`         // Margin market value
	MaturityDate      string `json:"matdt" xml:"matdt"`               // Instrument maturity date (FIXML)
	Mmy               string `json:"mmy" xml:"mmy"`                   // Instrument maturity year/month (FIXML)
	Multiplier        string `json:"mult" xml:"mult"`                 // Instrument multiplier
	OpenBuyValue      string `json:"openbuyvalue" xml:"openbuyvalue"` // Open buy value
	PutOrCall         string `json:"putcall" xml:"putcall"`           // put or call
	ShortBalance      string `json:"shortbalance" xml:"shortbalance"` // short balance (credit for sell?)
	ShortMarketValue  string `json:"shortmv" xml:"shortmv"`           // short market value
	StrikePrice       string `json:"strkpx" xml:"strkpx"`             // Instrument strike price (FIXML)
}
