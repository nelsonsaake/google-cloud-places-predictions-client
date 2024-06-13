package main

import (
	"google-cloud-places-predictions-client/conio"

	"github.com/nelsonsaake/go/do"
)

func init() {
	do.Die(do.LoadEnv())
}

func main() {
	conio.AddListener()
	conio.StartListening()
}
