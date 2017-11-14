package ally_api

import (
	"encoding/xml"
	"github.com/dghubble/oauth1"
	"github.com/jackmanlabs/errors"
	"net/http"
)

const (
	root string = "https://api.tradeking.com/v1"
)

type Client struct {
	http.Client
}

func NewClient(consumerKey, consumerSecret, oauthToken, oauthTokenSecret string) *Client {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(oauthToken, oauthTokenSecret)
	httpClient := oauth1.NewClient(oauth1.NoContext, config, token)

	client := new(Client)
	client.Client = *httpClient

	return client
}

func (client *Client) get(path string, target interface{}) error {

	resp, err := client.Get(root + path)
	if err != nil {
		return errors.Stack(err)
	}

	r, err := dump(resp.Body)
	if err != nil {
		return errors.Stack(err)
	}

	if resp.StatusCode != 200 {
		return errors.New("unexpected HTTP status: " + resp.Status)
	}

	err = xml.NewDecoder(r).Decode(target)
	if err != nil {
		return errors.Stack(err)
	}

	return nil
}
