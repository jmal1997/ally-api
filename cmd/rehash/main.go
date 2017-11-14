package main

import (
	"encoding/xml"
	"github.com/jackmanlabs/ally-api/lib"
	"github.com/jackmanlabs/errors"
	"log"
	"os"
)

func main() {

	var todo map[string]interface{} = map[string]interface{}{
		"GetAccount.xml":            new(lib.AccountResponse),
		"GetAccountBalances.xml":    new(lib.AccountBalancesResponse),
		"GetAccountHistory.xml":     new(lib.AccountHistoryResponse),
		"GetAccountHoldings.xml":    new(lib.AccountHoldingsResponse),
		"GetAccountOrders.xml":      new(lib.AccountOrdersResponse),
		"GetAccounts.xml":           new(lib.AccountsResponse),
		"GetAccountsBalances.xml":   new(lib.AccountsBalancesResponse),
		"GetClock.xml":              new(lib.ClockResponse),
		"GetExtQuotes.xml":          new(lib.ExtQuotesResponse),
		"GetNewsSearch.xml":         new(lib.NewsSearchResponse),
		"GetOptionsExpirations.xml": new(lib.OptionsExpirationsResponse),
		"GetOptionsSearch.xml":      new(lib.OptionsSearchResponse),
		"GetOptionsStrikes.xml":     new(lib.OptionsStrikesResponse),
		"GetProfile.xml":            new(lib.ProfileResponse),
		"GetStatus.xml":             new(lib.StatusResponse),
		"GetTimeSales.xml":          new(lib.TimeSalesResponse),
		"GetVersion.xml":            new(lib.VersionResponse),
		"GetWatchlists.xml":         new(lib.WatchListsResponse),
		//"GetTopLists.xml":           new(lib.TopListsResponse),
	}

	for inputName, target := range todo {
		inputFile, err := os.Open(inputName)
		if os.IsNotExist(err) {
			log.Fatal("Input file does not exist: ", inputName)
		} else if err != nil {
			log.Fatal(errors.Stack(err))
		}

		dec := xml.NewDecoder(inputFile)
		err = dec.Decode(target)
		if err != nil {
			log.Print(inputName)
			log.Fatal(errors.Stack(err))
		}

		err = inputFile.Close()
		if err != nil {
			log.Fatal(errors.Stack(err))
		}

		outputFile, err := os.Create(inputName + ".xml")
		if err != nil {
			log.Fatal(errors.Stack(err))
		}

		enc := xml.NewEncoder(outputFile)
		enc.Indent("", "\t")
		err = enc.EncodeElement(target, xml.StartElement{Name: xml.Name{Local: "response"}})
		if err != nil {
			log.Fatal(errors.Stack(err))
		}

		err = outputFile.Close()
		if err != nil {
			log.Fatal(errors.Stack(err))
		}

	}
}
