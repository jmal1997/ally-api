package lib

import (
	"net/http"
	"github.com/jackmanlabs/errors"
	"encoding/json"
)

// https://api.tradeking.com/v1/accounts.xml
/*
accounttype	Holdings attribute for where asset as held, "1"= cash, "2"= margin long, "5"=margin short.
accountvalue	Account value
accruedinterest	Amount of any accrued interest on the account
assetclass	Holding asset class type
cash	cash
cashavailable	cash available
cashavailableforwithdrawal	Cash available for withdrawal (cash & margin accounts only, n/a for retirement accounts)
cashbalance	cash balance
cashmv	Value of cash market value
cfi	Put or call code(FIXML)
change	Holding asset change for the day
costbasis	Holding cost basis
cusip	Instrument cusip
desc	Instrument description
equitypercentage	Percentage of equity (margin accounts only)
factor	Instrument factor
fedcall	Value of any fed call on account
gainloss	Holding gain/loss overall
housecall	Value of any house call
lastprice	Instrument last price
longoptions	Long option market value
longstocks	Long stock market value
marginbalance	Margin balance (- indicates debit balance, + indicates credit balance)
marginmv	Margin market value
marketvalue	Holding market value
marketvaluechange	Holding market value change
matdt	Instrument maturity date (FIXML)
mmf	Money market fund
mmy	Instrument maturity year/month (FIXML)
mult	Instrument multiplier
openbuyvalue	Open buy value
options	Total option market value. Path: /response/accounts/accountsummary/accountbalance/securities/
options	Options buying power. Path: /response/accounts/accountsummary/accountbalance/buyingpower/
price	Instrument price
purchaseprice	Holding purchase price
putcall	put or call
qty	Holding quantity
sectyp	Instrument security type (FIXML)
shortbalance	short balance (credit for sell?)
shortmv	short market value
shortoptions	Short option market value
shortstocks	Short stock market value
sodoptions	Start of day options buying power
sodstock	Start of day stock buying power
stock	Stock buying power
stocks	Total stock market value
strkpx	Instrument strike price (FIXML)
sym	Instrument underlying symbol (FIXML)
symbol	Holding underlying symbol
total	Total cash balance. Path: /response/accounts/accountsummary/accountbalance/money/
total	Total market value (stock & option). Path: /response/accounts/accountsummary/accountbalance/securities/
totalsecurities	Total account market value
uncleareddeposits	uncleared deposits
unsettledfunds	unsettle funds
yield	Yield
account	Account number
 */
func GetAccounts(client *http.Client) error {

	url := "https://api.tradeking.com/v1/accounts.xml"
	req, err := http.NewRequest("GET", url,nil)
	if err != nil{
		return errors.Stack(err)
	}

	resp, err := client.Do(req)
	if err != nil{
		return errors.Stack(err)
	}

	var accounts Accounts
	json.NewDecoder(resp.Body).Decode()

}

type Accounts struct{

	accounttype	Holdings attribute for where asset as held, "1"= cash, "2"= margin long, "5"=margin short.
	accountvalue	Account value
	accruedinterest	Amount of any accrued interest on the account
	assetclass	Holding asset class type
	cash	cash
	cashavailable	cash available
	cashavailableforwithdrawal	Cash available for withdrawal (cash & margin accounts only, n/a for retirement accounts)
	cashbalance	cash balance
	cashmv	Value of cash market value
	cfi	Put or call code(FIXML)
	change	Holding asset change for the day
	costbasis	Holding cost basis
	cusip	Instrument cusip
	desc	Instrument description
	equitypercentage	Percentage of equity (margin accounts only)
	factor	Instrument factor
	fedcall	Value of any fed call on account
	gainloss	Holding gain/loss overall
	housecall	Value of any house call
	lastprice	Instrument last price
	longoptions	Long option market value
	longstocks	Long stock market value
	marginbalance	Margin balance (- indicates debit balance, + indicates credit balance)
	marginmv	Margin market value
	marketvalue	Holding market value
	marketvaluechange	Holding market value change
	matdt	Instrument maturity date (FIXML)
	mmf	Money market fund
	mmy	Instrument maturity year/month (FIXML)
	mult	Instrument multiplier
	openbuyvalue	Open buy value
	options	Total option market value. Path: /response/accounts/accountsummary/accountbalance/securities/
	options	Options buying power. Path: /response/accounts/accountsummary/accountbalance/buyingpower/
	price	Instrument price
	purchaseprice	Holding purchase price
	putcall	put or call
	qty	Holding quantity
	sectyp	Instrument security type (FIXML)
	shortbalance	short balance (credit for sell?)
	shortmv	short market value
	shortoptions	Short option market value
	shortstocks	Short stock market value
	sodoptions	Start of day options buying power
	sodstock	Start of day stock buying power
	stock	Stock buying power
	stocks	Total stock market value
	strkpx	Instrument strike price (FIXML)
	sym	Instrument underlying symbol (FIXML)
	symbol	Holding underlying symbol
	total	Total cash balance. Path: /response/accounts/accountsummary/accountbalance/money/
	total	Total market value (stock & option). Path: /response/accounts/accountsummary/accountbalance/securities/
	totalsecurities	Total account market value
	uncleareddeposits	uncleared deposits
	unsettledfunds	unsettle funds
	yield	Yield
	account	Account number

}