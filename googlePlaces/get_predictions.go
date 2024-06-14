package googlePlaces

import (
	"fmt"
	"google-cloud-places-predictions-client/env"
	"net/url"

	"github.com/nelsonsaake/go/axios"
)

func makeUrl(input, types string) (string, error) {

	var (
		key     = env.GoogleCloudPlacesPredictionApiKey()
		baseUrl = env.GoogleCloudPlacesPredictionApiBaseUrl()
	)

	q := map[string]string{
		"key":   key,
		"input": input,
		"types": string(types),
	}

	values := url.Values{}
	for k, v := range q {
		values.Add(k, v)
	}

	url, err := url.Parse(baseUrl)
	if err != nil {
		return "", nil
	}

	url.RawQuery = values.Encode()
	return url.String(), nil
}

func die(err error) ([]string, error) {
	return nil, err
}

func getPredictions(input, types string) ([]string, error) {
	var (
		client  = &axios.Axios{}
		payload = &Resp{}
	)

	url, err := makeUrl(input, types)
	if err != nil {
		return die(err)
	}

	resp, err := client.Get(url)
	if err != nil {
		return die(err)
	}

	err = resp.Bind(payload)
	if err != nil {
		return die(err)
	}

	res := []string{}
	for _, v := range payload.Predictions {
		res = append(res, v.Description)
	}

	return res, nil
}

func GetPredictions(input string) ([]string, error) {

	input = fmt.Sprintf("%s, Ghana", input)

	var (
		scope = []PlaceType{Locality, Address, Sublocality}
		set   = map[string]string{}
	)

	for _, types := range scope {

		predictions, err := getPredictions(input, string(types))
		if err != nil {
			return die(err)
		}

		for _, v := range predictions {
			set[v] = v
		}

		if len(set) > 3 {
			break
		}
	}

	var res = []string{}
	for _, v := range set {
		res = append(res, v)
	}

	return res, nil
}
