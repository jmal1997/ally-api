package lib

import (
	"encoding/xml"
	"github.com/jackmanlabs/errors"
	"io"
	"net/http"
	"os"
)

func GetAccounts(client *http.Client) (*AccountsResponse, error) {

	url := "https://api.tradeking.com/v1/accounts.xml"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Stack(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Stack(err)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		return nil, errors.Stack(err)
	}

	var accountsResponse *AccountsResponse = new(AccountsResponse)
	xml.NewDecoder(resp.Body).Decode(&accountsResponse)

	return accountsResponse, nil
}

type AccountsResponse struct {
	Accounts struct {
		AccountSummary AccountSummary `xml:"accountsummary"`
	} `xml:"accounts"`
	ResponseId  string  `xml:"id,attr" json:"id"`
	ElapsedTime float64 `xml:"elapsedtime,omitempty"`
	Error       string  `xml:"error,omitempty"`
}

type AccountSummary struct {
	Account         int             `xml:"account"` // Account number
	AccountBalance  AccountBalance  `xml:"accountbalance"`
	AccountHoldings AccountHoldings `xml:"accountholdings"`
}

type AccountBalance struct {
	Account      int                `xml:"account"`
	AccountValue float64            `xml:"accountvalue"` // Account value
	BuyingPower  AccountBuyingPower `xml:"buyingpower"`
	FedCall      float64            `xml:"fedcall"`   // Value of any fed call on account
	HouseCall    float64            `xml:"housecall"` // Value of any house call
	Money        AccountMoney       `xml:"money"`
	Securities   AccountSecurities  `xml:"securities"`
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
	Holdings        []AccountHolding `xml:"holding"`
	TotalSecurities float64          `xml:"totalsecurities"` // Total account market value
	DisplayData     struct {
		TotalSecurities string `xml:"totalsecurities"` // Total account market value
	} `xml:"displaydata"`
}

type AccountBuyingPower struct {
	CashAvailableForWithdrawal float64 `xml:"cashavailableforwithdrawal,omitempty"` // Cash available for withdrawal (cash & margin accounts only, n/a for retirement accounts)
	DayTrading                 float64 `xml:"daytrading,omitempty"`
	EquityPercentage           float64 `xml:"equitypercentage,omitempty"` // Percentage of equity (margin accounts only)
	Options                    float64 `xml:"options,omitempty"`          // Options buying power. Path: /response/accounts/accountsummary/accountbalance/buyingpower/
	SodDayTrading              float64 `xml:"soddaytrading,omitempty"`
	SodOptions                 float64 `xml:"sodoptions,omitempty"` // Start of day options buying power
	SodStock                   float64 `xml:"sodstock,omitempty"`   // Start of day stock buying power
	Stock                      float64 `xml:"stock,omitempty"`      // Stock buying power
}

type AccountHolding struct {
	AccountType       int                       `xml:"accounttype"` // Holdings attribute for where asset as held, "1"= cash, "2"= margin long, "5"=margin short.
	CostBasis         float64                   `xml:"costbasis"`   // Holding cost basis
	DisplayData       AccountHoldingDisplayData `xml:"displaydata"`
	GainLoss          float64                   `xml:"gainloss"` // Holding gain/loss overall
	Instrument        AccountHoldingInstrument  `xml:"instrument"`
	MarketValue       float64                   `xml:"marketvalue"`       // Holding market value
	MarketValueChange float64                   `xml:"marketvaluechange"` // Holding market value change
	Price             float64                   `xml:"price"`             // Instrument price
	PurchasePrice     float64                   `xml:"purchaseprice"`     // Holding purchase price
	Qty               float64                   `xml:"qty"`               // Holding quantity
	Underlying        struct{}                  `xml:"underlying"`
	Quote             struct {
		Change    float64 `xml:"change"`    // Holding asset change for the day
		LastPrice float64 `xml:"lastprice"` // Instrument last price
	} `xml:"quote"`
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

type Account struct {
	CashBalance       string `xml:"cashbalance"`  // cash balance
	CashMarketValue   string `xml:"cashmv"`       // Value of cash market value
	Cfi               string `xml:"cfi"`          // Put or call code(FIXML)
	MarginMarketValue string `xml:"marginmv"`     // Margin market value
	MaturityDate      string `xml:"matdt"`        // Instrument maturity date (FIXML)
	Mmy               string `xml:"mmy"`          // Instrument maturity year/month (FIXML)
	Multiplier        string `xml:"mult"`         // Instrument multiplier
	OpenBuyValue      string `xml:"openbuyvalue"` // Open buy value
	PutOrCall         string `xml:"putcall"`      // put or call
	ShortBalance      string `xml:"shortbalance"` // short balance (credit for sell?)
	ShortMarketValue  string `xml:"shortmv"`      // short market value
	StrikePrice       string `xml:"strkpx"`       // Instrument strike price (FIXML)
}
