package googlePlaces

type Resp struct {
	Predictions []struct {
		Description string `json:"description"`
	} `json:"predictions"`
}
