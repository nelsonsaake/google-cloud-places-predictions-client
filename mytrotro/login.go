package mytrotro

import (
	"google-cloud-places-predictions-client/env"
	"google-cloud-places-predictions-client/models"

	"github.com/nelsonsaake/go/do"
)

func Login() {

	req := map[string]any{
		"email":    env.MytrotroUsername(),
		"password": env.MytrotroPassword(),
	}

	payload := models.Auth{}

	res, err := client().Post("/auth/login", req)
	do.Die(err)

	err = res.Bind(&payload)
	do.Die(err)

	setAuthToken(payload.Token)
}
