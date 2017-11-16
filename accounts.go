package ally_api

import (
	"fmt"
	"github.com/jackmanlabs/errors"
	"time"
)

// 	GET accounts
func (client *Client) GetAccounts() (*AccountsResponse, error) {
	var (
		path   string            = "/accounts.xml"
		target *AccountsResponse = new(AccountsResponse)
	)

	err := client.get(path, target)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return target, nil
}

//GET accounts/balances
func (client *Client) GetAccountsBalances() (*AccountsBalancesResponse, error) {
	var (
		path   string                    = "/accounts/balances.xml"
		target *AccountsBalancesResponse = new(AccountsBalancesResponse)
	)

	err := client.get(path, target)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return target, nil
}

//	GET accounts/{id}
func (client *Client) GetAccount(id int) (*AccountResponse, error) {
	var (
		path   string           = fmt.Sprintf("/accounts/%d.xml", id)
		target *AccountResponse = new(AccountResponse)
	)

	err := client.get(path, target)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return target, nil
}

//GET accounts/{id}/balances
func (client *Client) GetAccountBalances(id int) (*AccountBalancesResponse, error) {
	var (
		path   string                   = fmt.Sprintf("/accounts/%d/balances.xml", id)
		target *AccountBalancesResponse = new(AccountBalancesResponse)
	)

	err := client.get(path, target)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return target, nil
}

//GET accounts/{id}/history
func (client *Client) GetAccountHistory(id int) (*AccountHistoryResponse, error) {
	var (
		path   string                  = fmt.Sprintf("/accounts/%d/history.xml", id)
		target *AccountHistoryResponse = new(AccountHistoryResponse)
	)

	err := client.get(path, target)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return target, nil
}

//GET accounts/{id}/holdings
func (client *Client) GetAccountHoldings(id int) (*AccountHoldingsResponse, error) {
	var (
		path   string                   = fmt.Sprintf("/accounts/%d/holdings.xml", id)
		target *AccountHoldingsResponse = new(AccountHoldingsResponse)
	)

	err := client.get(path, target)
	if err != nil {
		return nil, errors.Stack(err)
	}

	return target, nil
}

type AccountsResponse struct {
	AccountSummaries []AccountSummary `xml:"accounts>accountsummary"`
	Response
}

type AccountsBalancesResponse struct {
	AccountBalances []AccountBalanceSimple `xml:"accountbalance"`
	Response
	TotalBalance float64 `xml:"totalbalance>accountvalue"`
}

type AccountBalancesResponse struct {
	Response
	AccountBalance AccountBalance `xml:"accountbalance"`
}

type AccountResponse struct {
	AccountBalance  AccountBalance  `xml:"accountbalance"`
	AccountHoldings AccountHoldings `xml:"accountholdings"`
	Response
}

type AccountHoldingsResponse struct{ Response }

type AccountHistoryResponse struct {
	Response
	Transactions []AccountTransaction `xml:"transactions>transaction"`
}

type AccountTransaction struct {
	Activity    string                   `xml:"activity"`
	Amount      float64                  `xml:"amount"`
	Date        time.Time                `xml:"date"`
	Desc        string                   `xml:"desc"`
	Symbol      string                   `xml:"symbol"`
	Transaction AccountTransactionDetail `xml:"transaction"`
}

type AccountTransactionDetail struct {
	Accounttype     int                        `xml:"accounttype"`
	Commission      float64                    `xml:"commission"`
	Description     string                     `xml:"description"`
	Fee             float64                    `xml:"fee"`
	Price           float64                    `xml:"price"`
	Quantity        float64                    `xml:"quantity"`
	Secfee          float64                    `xml:"secfee"`
	Security        AccountTransactionSecurity `xml:"security"`
	Settlementdate  time.Time                  `xml:"settlementdate"`
	Side            int                        `xml:"side"`
	Source          string                     `xml:"source"`
	Tradedate       time.Time                  `xml:"tradedate"`
	Transactionid   int                        `xml:"transactionid"`
	Transactiontype string                     `xml:"transactiontype"`
}

type AccountTransactionSecurity struct {
	Cusip  int    `xml:"cusip"`
	Id     int    `xml:"id"`
	Sectyp string `xml:"sectyp"`
	Sym    string `xml:"sym"`
}

type AccountSummary struct {
	Account         int             `xml:"account"` // Account number
	AccountBalance  AccountBalance  `xml:"accountbalance"`
	AccountHoldings AccountHoldings `xml:"accountholdings"`
}

type AccountBalance struct {
	Account      int                `xml:"account"`
	AccountValue float64            `xml:"accountvalue"` // Account value
	BuyingPower  AccountBuyingPower `xml:"buyingpower"`  //
	FedCall      float64            `xml:"fedcall"`      // Value of any fed call on account
	HouseCall    float64            `xml:"housecall"`    // Value of any house call
	Money        AccountMoney       `xml:"money"`        //
	Securities   AccountSecurities  `xml:"securities"`   //
}

type AccountBalanceSimple struct {
	Account      int     `xml:"account"`
	AccountName  string  `xml:"accountname"`  // Used in the 'GET accounts/balances' response
	AccountValue float64 `xml:"accountvalue"` // Account value
}

type AccountMoney struct {
	AccruedInterest   float64 `xml:"accruedinterest"`   // Amount of any accrued interest on the account
	Cash              float64 `xml:"cash"`              // cash
	CashAvailable     float64 `xml:"cashavailable"`     // cash available
	MarginBalance     float64 `xml:"marginbalance"`     // Margin balance (- indicates debit balance, + indicates credit balance)
	Mmf               float64 `xml:"mmf"`               // Money market fund
	Total             float64 `xml:"total"`             // Total cash balance. Path: /response/accounts/accountsummary/accountbalance/money/
	UnclearedDeposits float64 `xml:"uncleareddeposits"` // uncleared deposits
	UnsettledFunds    float64 `xml:"unsettledfunds"`    // unsettle funds
	Yield             float64 `xml:"yield"`             // Yield
}

type AccountSecurities struct {
	LongOptions  float64 `xml:"longoptions"`  // Long option market value
	LongStocks   float64 `xml:"longstocks"`   // Long stock market value
	Options      float64 `xml:"options"`      // Total option market value. Path: /response/accounts/accountsummary/accountbalance/securities/
	ShortOptions float64 `xml:"shortoptions"` // Short option market value
	ShortStocks  float64 `xml:"shortstocks"`  // Short stock market value
	Stocks       float64 `xml:"stocks"`       // Total stock market value
	Total        float64 `xml:"total"`        // Total market value (stock & option). Path: /response/accounts/accountsummary/accountbalance/securities/
}

type AccountHoldings struct {
	Holdings        []AccountHolding           `xml:"holding"`
	TotalSecurities float64                    `xml:"totalsecurities"` // Total account market value
	DisplayData     AccountHoldingsDisplayData `xml:"displaydata"`
}

type AccountHoldingsDisplayData struct {
	TotalSecurities string `xml:"totalsecurities"` // Total account market value
}

type AccountBuyingPower struct {
	CashAvailableForWithdrawal float64 `xml:"cashavailableforwithdrawal,omitempty"` // Cash available for withdrawal (cash & margin accounts only, n/a for retirement accounts)
	DayTrading                 float64 `xml:"daytrading,omitempty"`                 //
	EquityPercentage           float64 `xml:"equitypercentage,omitempty"`           // Percentage of equity (margin accounts only)
	Options                    float64 `xml:"options,omitempty"`                    // Options buying power. Path: /response/accounts/accountsummary/accountbalance/buyingpower/
	SodDayTrading              float64 `xml:"soddaytrading,omitempty"`              //
	SodOptions                 float64 `xml:"sodoptions,omitempty"`                 // Start of day options buying power
	SodStock                   float64 `xml:"sodstock,omitempty"`                   // Start of day stock buying power
	Stock                      float64 `xml:"stock,omitempty"`                      // Stock buying power
}

type AccountHolding struct {
	AccountType       int                       `xml:"accounttype"`       // Holdings attribute for where asset as held, "1"= cash, "2"= margin long, "5"=margin short.
	CostBasis         float64                   `xml:"costbasis"`         // Holding cost basis
	DisplayData       AccountHoldingDisplayData `xml:"displaydata"`       //
	GainLoss          float64                   `xml:"gainloss"`          // Holding gain/loss overall
	Instrument        AccountHoldingInstrument  `xml:"instrument"`        //
	MarketValue       float64                   `xml:"marketvalue"`       // Holding market value
	MarketValueChange float64                   `xml:"marketvaluechange"` // Holding market value change
	Price             float64                   `xml:"price"`             // Instrument price
	PurchasePrice     float64                   `xml:"purchaseprice"`     // Holding purchase price
	Qty               float64                   `xml:"qty"`               // Holding quantity
	Quote             AccountHoldingQuote       `xml:"quote"`
	Underlying        struct{}                  `xml:"underlying"`
}

type AccountHoldingQuote struct {
	Change    float64 `xml:"change"`    // Holding asset change for the day
	LastPrice float64 `xml:"lastprice"` // Instrument last price
}

type AccountHoldingDisplayData struct {
	AccountType       string `xml:"accounttype"`       // Holdings attribute for where asset as held, "1"= cash, "2"= margin long, "5"=margin short.
	AssetClass        string `xml:"assetclass"`        // Holding asset class type
	Change            string `xml:"change"`            // Holding asset change for the day
	CostBasis         string `xml:"costbasis"`         // Holding cost basis
	Desc              string `xml:"desc"`              // Instrument description
	LastPrice         string `xml:"lastprice"`         // Instrument last price
	MarketValue       string `xml:"marketvalue"`       // Holding market value
	MarketValueChange string `xml:"marketvaluechange"` // Holding market value change
	Qty               string `xml:"qty"`               // Holding quantity
	Symbol            string `xml:"symbol"`            // Holding underlying symbol
}

type AccountHoldingInstrument struct {
	Cusip        string  `xml:"cusip"`  // Instrument cusip
	Desc         string  `xml:"desc"`   // Instrument description
	Factor       float64 `xml:"factor"` // Instrument factor
	SecurityType string  `xml:"sectyp"` // Instrument security type (FIXML)
	Symbol       string  `xml:"sym"`    // Instrument underlying symbol (FIXML)
}
