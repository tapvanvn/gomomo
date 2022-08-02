package miniapp_test

import (
	"fmt"
	"testing"

	"github.com/tapvanvn/gomomo/config"
	"github.com/tapvanvn/gomomo/miniapp"
)

var testDomain = "https://api.mservice.com.vn/openapi"

func Test1(t *testing.T) {
	config := &config.ClientConfig{
		OpenSecret:     "",
		OpenPrivateKey: "",
		OpenPublicKey:  "",
	}
	appClient := miniapp.NewMiniAppClient(true, config)
	partnerUserID := ""
	authCode := ""
	token, err := appClient.RequestAccessToken(partnerUserID, authCode)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(token.Token)
}
