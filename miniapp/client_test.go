package miniapp_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/tapvanvn/gomomo/config"
	"github.com/tapvanvn/gomomo/miniapp"
	"github.com/tapvanvn/gorouter/v2"
)

var testDomain = "https://api.mservice.com.vn/openapi"

func initRoute(appClient *miniapp.MiniAppClient) *gorouter.Router {
	var router = &gorouter.Router{}

	builder := gorouter.NewStructureBuilder()
	var handers = map[string]gorouter.EndpointDefine{}

	builder.AddOneLine("gateway/open/v1/oauth/accessToken")
	handers["gateway/open/v1/oauth/accessToken"] = gorouter.EndpointDefine{
		Handles:   nil,
		ApiFormer: miniapp.NewApiFormAccessKey(appClient),
	}
	router.Init("", builder.Export(), handers)
	return router
}

func Test1(t *testing.T) {
	config := &config.ClientConfig{
		OpenSecret:     "",
		OpenPrivateKey: "",
		OpenPublicKey:  "",
	}
	appClient := miniapp.NewMiniAppClient(true, config)
	router := initRoute(appClient)
	res, err := router.Request(testDomain, "gateway/open/v1/oauth/accessToken", map[string]interface{}{
		"partnerUserId": "",
		"authCode":      "",
	})
	if err != nil {
		t.Error(err)
	}
	body, err := ioutil.ReadAll(res.Base.Body)

	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(body))
}
