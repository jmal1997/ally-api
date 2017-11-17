# ally-api

Go API bindings for Ally Invest (formerly TradeKing).

Ally Invest documentation can be found here: https://www.ally.com/api/invest/documentation/getting-started/

## Status

This project is currently under development.

TODO:
* Define all response types to coincide with the API documentation.

The bindings take advantage of the XML output of the Ally Invest API. I do not
plan on supporting anything else. They specifically state in their API
documentation that XML is output by default. Furthermore, they only provide
sample XML output (not JSON). Therefore, I'm assuming XML is what Ally Invest
prefers to support.

## CLI Tool

At this stage of development, I recommend you look at the CLI client tool found
in the root of the project. It is the primary interface for testing the library.

The CLI tool, for convenience, relies on a configuration file. Simply specify a
configuration file. It doesn't have to exist. When you specify a non-existing
configuration file, the CLI tool will create one. Then, you can populate it, and
use the now-populated file.

```
jackman@dorito:~/gopath/src/github.com/jackmanlabs/ally-api$ go build cmd/cli

jackman@dorito:~/gopath/src/github.com/jackmanlabs/ally-api$ ./cli 
Usage of ./cli:
  -config string
        The path of the configuration file. Specifying a non-existent file will cause an empty one to be created.
2017/11/14 16:06:08 A config file path must be specified. If the file you specify does not exist, a new one will be created for you which you may populate.

jackman@dorito:~/gopath/src/github.com/jackmanlabs/ally-api$ ./cli -config config.json
2017/11/14 16:21:01 New, empty config file successfully created: config.json
```

Another execution of `./cli -config config.json` will result in a great deal of
usage information:

```
jackman@dorito:~/gopath/src/github.com/jackmanlabs/ally-api$ ./cli -config config.json
2017/11/14 16:23:45 

This application requires at least two arguments, an operation and a
resource.

Valid operations include REST operations, such as GET, PUT, POST, DELETE,
and PATCH. Capitalization is not required.

At the time of this writing, the following operation/resource combinations
are defined:

    ACCOUNT CALLS
    =============
    GET     accounts
    GET     accounts/balances
    GET     accounts/{id}
    GET     accounts/{id}/balances
    GET     accounts/{id}/history
    GET     accounts/{id}/holdings
    
    ORDER/TRADE CALLS
    =================
    GET     accounts/{id}/orders
    POST    accounts/{id}/orders
    POST    accounts/{id}/orders/preview
    
    MARKET CALLS
    ============
    GET     market/clock
    GET     market/ext/quotes
    GET     market/news/search
    GET     market/news/{id}
    GET     market/options/search
    GET     market/options/strikes
    GET     market/options/expirations
    GET     market/timesales
    GET     market/toplists
    GET     market/quotes (streaming)
    
    MEMBER CALLS
    ============
    GET     member/profile
    
    UTILITY CALLS
    =============
    GET     utility/status
    GET     utility/version
    
    WATCHLIST CALLS
    ===============
    GET     watchlists
    POST    watchlists
    GET     watchlists/{id}
    DELETE  watchlists/{id}
    POST    watchlists/{id}/symbols
    DELETE  watchlists/{id}/symbols
    
    DEPRECATED
    ==========
    GET market/chains
    GET market/quotes

Because of the way the library is written, and because the developer hasn't come
up with something better, when you want to make a call with a URL parameter,
append the value of the parameter to the end of the command, like so:

    GET watchlists/{id} 12

If the URL contains multiple parameters, the values you append should be in the
same order as they are found in the URL template.

This tool was written to test the library. Therefore, I don't want to encourage
anyone to think this tool's output will be consistent. At the time of this
writing, the tool does two things. First, it dumps the raw response body to a
file, named after the library function being called. Second, whatever was
successfully parsed is logged to stderr in the form of JSON.

```

## Rehash Tool

The Rehash tool is used to compare raw XML from the server and the data parsed
by the library.

## Re-Use of Ally Invest Documentation

In the process of creating the various data types to correspond with Ally's API,
the comments in those data types are derived from the descriptions in the Ally
documentation. This is done simply for convenience. Credit goes to Ally and
their API documentation writers.

I do not intend to update this documentation as it would be time-consuming to
track.

Furthermore, if Ally Invest is offended by the reuse of technical documentation,
I will be happy to make things right. 

## API discrepencies

### /market/ext/quotes
pchg example has a %, but the live API does not.
sho example has commas, live API does not.
tcond is numeric in live API.