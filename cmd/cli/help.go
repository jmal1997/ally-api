package main

import "log"

func help() {
	log.Print(`

This application requires at least two arguments, an operation and a
resource.

Valid operations include REST operations, such as GET, PUT, POST, DELETE,
and PATCH. Capitalization is not required.

At the time of this writing, the following operation/resource combinations
are defined:

	ACCOUNT CALLS
	=============
	GET	accounts
	GET	accounts/balances
	GET	accounts/{id}
	GET	accounts/{id}/balances
	GET	accounts/{id}/history
	GET	accounts/{id}/holdings

	ORDER/TRADE CALLS
	=================
	GET		accounts/{id}/orders
	POST	accounts/{id}/orders
	POST	accounts/{id}/orders/preview

	MARKET CALLS
	============
	GET	market/clock
	GET	market/ext/quotes
	GET	market/news/search
	GET	market/news/{id}
	GET	market/options/search
	GET	market/options/strikes
	GET	market/options/expirations
	GET	market/timesales
	GET	market/toplists
	GET	market/quotes (streaming)

	MEMBER CALLS
	============
	GET	member/profile

	UTILITY CALLS
	=============
	GET	utility/status
	GET	utility/version

	WATCHLIST CALLS
	===============
	GET		watchlists
	POST	watchlists
	GET		watchlists/{id}
	DELETE	watchlists/{id}
	POST	watchlists/{id}/symbols
	DELETE	watchlists/{id}/symbols

	DEPRECATED
	==========
	About Deprecation
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

`)
}
