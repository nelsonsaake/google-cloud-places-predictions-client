package pre

import (
	"google-cloud-places-predictions-client/env"
	"net/url"

	"github.com/nelsonsaake/go/axios"
	"github.com/nelsonsaake/go/do"
	"github.com/nelsonsaake/go/pretty"
)

var (
	baseUrl = env.GoogleCloudPlacesPredictionApiBaseUrl()
	key     = env.GoogleCloudPlacesPredictionApiKey()
	input   = ""
	types   PlaceType
)

func GetPredictions(input string) {

	q := url.Values{}
	q.Add("key", key)
	q.Add("input", input)
	q.Add("types", string(types))

	url := url.URL{
		Path:     baseUrl,
		RawQuery: q.Encode(),
	}

	client := &axios.Axios{}

	resp, err := client.Get(url.String())
	do.Die(err)

	json, err := resp.Json()
	do.Die(err)

	pretty.Print(json)
}
