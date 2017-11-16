package main

import (
	"encoding/xml"
	"github.com/jackmanlabs/ally-api"
	"github.com/jackmanlabs/errors"
	"log"
	"os"
)

func main() {

	var todo map[string]interface{} = map[string]interface{}{
		"GetAccount.xml":            new(ally_api.AccountResponse),
		"GetAccountBalances.xml":    new(ally_api.AccountBalancesResponse),
		"GetAccountHistory.xml":     new(ally_api.AccountHistoryResponse),
		"GetAccountHoldings.xml":    new(ally_api.AccountHoldingsResponse),
		"GetAccountOrders.xml":      new(ally_api.AccountOrdersResponse),
		"GetAccounts.xml":           new(ally_api.AccountsResponse),
		"GetAccountsBalances.xml":   new(ally_api.AccountsBalancesResponse),
		"GetClock.xml":              new(ally_api.ClockResponse),
		"GetExtQuotes.xml":          new(ally_api.ExtQuotesResponse),
		"GetNewsSearch.xml":         new(ally_api.NewsSearchResponse),
		"GetOptionsExpirations.xml": new(ally_api.OptionsExpirationsResponse),
		"GetOptionsSearch.xml":      new(ally_api.OptionsSearchResponse),
		"GetOptionsStrikes.xml":     new(ally_api.OptionsStrikesResponse),
		"GetProfile.xml":            new(ally_api.ProfileResponse),
		"GetStatus.xml":             new(ally_api.StatusResponse),
		"GetTimeSales.xml":          new(ally_api.TimeSalesResponse),
		"GetVersion.xml":            new(ally_api.VersionResponse),
		"GetWatchlists.xml":         new(ally_api.WatchListsResponse),
		//"GetTopLists.xml":           new(ally_api.TopListsResponse),
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
