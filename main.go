package main

import (
	"encoding/json"
	"flag"
	"github.com/dghubble/oauth1"
	"github.com/jackmanlabs/ally-api/lib"
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
	jlog.Log(args)

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

	// right now,
	if len(args) == 3 {
		id, err = strconv.Atoi(args[2])
		if err != nil {
			log.Fatal(errors.Stack(err))
		}
	}

	if strings.Contains(path, "{id}") && len(args) != 3 {
		log.Fatal("The resource you specified requires a single parameter.")
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

	oauthConfig := oauth1.NewConfig(config.ConsumerKey, config.ConsumerSecret)
	oauthToken := oauth1.NewToken(config.Token, config.TokenSecret)
	oauthClient := oauth1.NewClient(oauth1.NoContext, oauthConfig, oauthToken)

	var response interface{}

	switch {

	case op == "DELETE" && path == "watchlists/{id}":
		response, err = lib.DeleteWatchlist(oauthClient, id)
	case op == "DELETE" && path == "watchlists/{id}/symbols":
		response, err = lib.DeleteWatchlistSymbols(oauthClient, id)
	case op == "GET" && path == "accounts":
		response, err = lib.GetAccounts(oauthClient)
	case op == "GET" && path == "accounts/balances":
		response, err = lib.GetAccountsBalances(oauthClient)
	case op == "GET" && path == "accounts/{id}":
		response, err = lib.GetAccount(oauthClient, id)
	case op == "GET" && path == "accounts/{id}/balances":
		response, err = lib.GetAccountBalances(oauthClient, id)
	case op == "GET" && path == "accounts/{id}/history":
		response, err = lib.GetAccountHistory(oauthClient, id)
	case op == "GET" && path == "accounts/{id}/holdings":
		response, err = lib.GetAccountHoldings(oauthClient, id)
	case op == "GET" && path == "accounts/{id}/orders":
		response, err = lib.GetAccountOrders(oauthClient, id)
	case op == "GET" && path == "market/clock":
		response, err = lib.GetClock(oauthClient)
	case op == "GET" && path == "market/ext/quotes":
		response, err = lib.GetExtQuotes(oauthClient)
	case op == "GET" && path == "market/news/search":
		response, err = lib.GetNewsSearch(oauthClient)
	case op == "GET" && path == "market/news/{id}":
		response, err = lib.GetNews(oauthClient, id)
	case op == "GET" && path == "market/options/expirations":
		response, err = lib.GetOptionsExpirations(oauthClient)
	case op == "GET" && path == "market/options/search":
		response, err = lib.GetOptionsSearch(oauthClient)
	case op == "GET" && path == "market/options/strikes":
		response, err = lib.GetOptionsStrikes(oauthClient)
	case op == "GET" && path == "market/quotes":
		response, err = lib.GetQuotes(oauthClient)
	case op == "GET" && path == "market/timesales":
		response, err = lib.GetTimeSales(oauthClient)
	case op == "GET" && path == "market/toplists":
		response, err = lib.GetTopLists(oauthClient)
	case op == "GET" && path == "member/profile":
		response, err = lib.GetProfile(oauthClient)
	case op == "GET" && path == "utility/status":
		response, err = lib.GetStatus(oauthClient)
	case op == "GET" && path == "utility/version":
		response, err = lib.GetVersion(oauthClient)
	case op == "GET" && path == "watchlists":
		response, err = lib.GetWatchlists(oauthClient)
	case op == "GET" && path == "watchlists/{id}":
		response, err = lib.GetWatchlist(oauthClient, id)
	case op == "POST" && path == "accounts/{id}/orders":
		response, err = lib.PostAccountOrders(oauthClient, id)
	case op == "POST" && path == "accounts/{id}/orders/preview":
		response, err = lib.PostAccountOrderPreview(oauthClient, id)
	case op == "POST" && path == "watchlists":
		response, err = lib.PostWatchlists(oauthClient)
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
