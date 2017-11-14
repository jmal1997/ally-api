package main

import (
	"encoding/json"
	"flag"
	"github.com/jackmanlabs/ally-api"
	"github.com/jackmanlabs/bucket/jlog"
	"github.com/jackmanlabs/errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	var (
		configPath *string = flag.String("config", "", "The path of the configuration file. Specifying a non-existent file will cause an empty one to be created.")
	)
	flag.Parse()

	if *configPath == "" {
		flag.Usage()
		log.Fatal("A config file path must be specified. If the file you specify does not exist, a new one will be created for you which you may populate.")
	}

	args := flag.Args()

	if len(args) < 2 {
		help()
		os.Exit(1)
	}

	var (
		err  error
		id   int
		op   string = args[0]
		path string = args[1]
	)

	if strings.Contains(path, "{id}") && len(args) != 3 {
		log.Fatal("The resource you specified requires a single parameter.")
	}

	if strings.Contains(path, "{id}") && len(args) == 3 {
		id, err = strconv.Atoi(args[2])
		if err != nil {
			log.Fatal(errors.Stack(err))
		}
	}

	if !strings.Contains(path, "{id}") && len(args) > 2 {
		args = args[2:]
	}

	configFile, err := os.Open(*configPath)
	if os.IsNotExist(err) {
		configFile, err = os.Create(*configPath)
		if err != nil {
			log.Print(errors.Stack(err))
			log.Fatal("Unable to create the new, empty config file.")
		}

		enc := json.NewEncoder(configFile)
		enc.SetIndent("", "\t")
		err = enc.Encode(Config{})
		if err != nil {
			log.Print(errors.Stack(err))
			log.Fatal("Failure to write empty configuration to new file.")
		}

		err = configFile.Close()
		if err != nil {
			log.Fatal(errors.Stack(err))
		}

		log.Print("New, empty config file successfully created: ", *configPath)
		os.Exit(0)

	} else if err != nil {
		log.Fatal(errors.Stack(err))
	}

	var config Config
	dec := json.NewDecoder(configFile)
	err = dec.Decode(&config)
	if err != nil {
		log.Print(errors.Stack(err))
		log.Fatal("Failure to read existing configuration.")
	}

	err = configFile.Close()
	if err != nil {
		log.Fatal(errors.Stack(err))
	}

	client := ally_api.NewClient(config.ConsumerKey, config.ConsumerSecret, config.Token, config.TokenSecret)

	var response interface{}

	switch {
	case op == "DELETE" && path == "watchlists/{id}":
		response, err = client.DeleteWatchlist(id)
	case op == "DELETE" && path == "watchlists/{id}/symbols":
		response, err = client.DeleteWatchlistSymbols(id)
	case op == "GET" && path == "accounts":
		response, err = client.GetAccounts()
	case op == "GET" && path == "accounts/balances":
		response, err = client.GetAccountsBalances()
	case op == "GET" && path == "accounts/{id}":
		response, err = client.GetAccount(id)
	case op == "GET" && path == "accounts/{id}/balances":
		response, err = client.GetAccountBalances(id)
	case op == "GET" && path == "accounts/{id}/history":
		response, err = client.GetAccountHistory(id)
	case op == "GET" && path == "accounts/{id}/holdings":
		response, err = client.GetAccountHoldings(id)
	case op == "GET" && path == "accounts/{id}/orders":
		response, err = client.GetAccountOrders(id)
	case op == "GET" && path == "market/clock":
		response, err = client.GetClock()
	case op == "GET" && path == "market/ext/quotes":
		response, err = client.GetExtQuotes(args)
	case op == "GET" && path == "market/news/search":
		response, err = client.GetNewsSearch()
	case op == "GET" && path == "market/news/{id}":
		response, err = client.GetNews(id)
	case op == "GET" && path == "market/options/expirations":
		response, err = client.GetOptionsExpirations()
	case op == "GET" && path == "market/options/search":
		response, err = client.GetOptionsSearch()
	case op == "GET" && path == "market/options/strikes":
		response, err = client.GetOptionsStrikes()
	case op == "GET" && path == "market/quotes":
		response, err = client.GetQuotes()
	case op == "GET" && path == "market/timesales":
		response, err = client.GetTimeSales()
	case op == "GET" && path == "market/toplists":
		response, err = client.GetTopLists()
	case op == "GET" && path == "member/profile":
		response, err = client.GetProfile()
	case op == "GET" && path == "utility/status":
		response, err = client.GetStatus()
	case op == "GET" && path == "utility/version":
		response, err = client.GetVersion()
	case op == "GET" && path == "watchlists":
		response, err = client.GetWatchlists()
	case op == "GET" && path == "watchlists/{id}":
		response, err = client.GetWatchlist(id)
	case op == "POST" && path == "accounts/{id}/orders":
		response, err = client.PostAccountOrders(id)
	case op == "POST" && path == "accounts/{id}/orders/preview":
		response, err = client.PostAccountOrderPreview(id)
	case op == "POST" && path == "watchlists":
		response, err = client.PostWatchlists()
	case op == "POST" && path == "watchlists/{id}/symbols":

		// case op == "GET" && path == "market/quotes":
		// case op == "GET" && path == "market/chains":
	}
	if err != nil {
		log.Fatal(errors.Stack(err))
	}

	jlog.Log(response)
}

type Config struct {
	ConsumerKey    string
	ConsumerSecret string
	Token          string
	TokenSecret    string
}
