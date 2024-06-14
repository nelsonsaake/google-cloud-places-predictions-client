package mytrotro

import (
	"fmt"

	"google-cloud-places-predictions-client/env"

	"github.com/nelsonsaake/go/axios"
)

var (
	headers = map[string]string{}
)

func setAuthToken(token string) {
	headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
}

func client() *axios.Axios {
	client := axios.New()
	client.SetBaseUrl(env.MytrotroApiBaseUrl())
	client.AddHeaders(headers)
	return client
}
