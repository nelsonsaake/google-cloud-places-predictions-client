package main

import (
	"fmt"
	"google-cloud-places-predictions-client/googlePlaces"
	"google-cloud-places-predictions-client/mytrotro"

	"github.com/nelsonsaake/go/cin"
	"github.com/nelsonsaake/go/do"
)

func init() {
	do.Die(do.LoadEnv())
}

func main() {
	//...

	mytrotro.Login()

	for {
		input := ""

		fmt.Print(">> ")
		cin.Line(&input)

		mytrotro.GetTrips("ho")
	}
}

func predictions(input string) {
	predictions, err := googlePlaces.GetPredictions(input)
	do.Die(err)

	for i, v := range predictions {
		fmt.Printf("%2v  | %s \n", i+1, v)
	}
}

func getTrips(input string) {

}
