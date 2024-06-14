package mytrotro

import (
	"fmt"
	"google-cloud-places-predictions-client/models"

	"github.com/nelsonsaake/go/axios"
	"github.com/nelsonsaake/go/do"
)

func GetTrips(input string) {

	query := axios.QueryParam{"search": input}
	payload := models.PaginatedTripResponse{}

	res, err := client().Get("/trips", query)
	do.Die(err)

	err = res.Bind(&payload)
	do.Die(err)

	fmt.Println("---")
	for _, v := range payload.Data {
		fmt.Println(v)
	}
}
