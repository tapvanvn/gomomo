package miniapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/tapvanvn/gomomo/common"
	"github.com/tapvanvn/gomomo/config"
	"github.com/tapvanvn/gomomo/entity"
	"github.com/tapvanvn/gorouter/v2"
)

type MiniAppClient struct {
	router     *gorouter.Router
	OpenSecret string //The secret key provided by momo when register miniapp
	domain     string
}

func NewMiniAppClient(isDev bool, config *config.ClientConfig) *MiniAppClient {
	client := &MiniAppClient{
		router:     &gorouter.Router{},
		OpenSecret: config.OpenSecret,
	}
	if isDev {
		client.domain = common.DomainDev
	} else {
		client.domain = common.DomainProd
	}
	builder := gorouter.NewStructureBuilder()
	var handers = map[string]gorouter.EndpointDefine{}
	builder.AddOneLine("gateway/open/v1/oauth/accessToken")
	handers["gateway/open/v1/oauth/accessToken"] = gorouter.EndpointDefine{
		Handles:   nil,
		ApiFormer: NewApiFormAccessToken(client),
	}
	client.router.Init("", builder.Export(), handers)
	return client
}

//RequestAccessKey request an access key from authCode
func (client *MiniAppClient) RequestAccessToken(partnerUserID string, authCode string) (*entity.AccessToken, error) {
	res, err := client.router.Request(client.domain, "gateway/open/v1/oauth/accessToken", map[string]interface{}{
		"partnerUserId": partnerUserID,
		"authCode":      authCode,
	})
	fmt.Println("partnerUserID:", partnerUserID, "\nauCode:", authCode)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Base.Body)
	defer res.Close()
	if err != nil {
		return nil, err
	}

	accessToken := &entity.AccessToken{}
	if err := json.Unmarshal(body, accessToken); err != nil {
		return nil, err
	}
	return accessToken, nil
}
