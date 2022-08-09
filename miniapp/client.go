package miniapp

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
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
	PrivateKey *rsa.PrivateKey
}

func NewMiniAppClient(isDev bool, config *config.ClientConfig) (*MiniAppClient, error) {
	client := &MiniAppClient{
		router:     &gorouter.Router{},
		OpenSecret: config.OpenSecret,
	}
	ppkData, err := base64.URLEncoding.DecodeString(config.OpenPrivateKey)
	if err != nil {
		fmt.Println(config.OpenPrivateKey)
		return nil, err
	}
	key, err := x509.ParsePKCS8PrivateKey(ppkData)
	if err != nil {
		return nil, err
	}
	client.PrivateKey = key.(*rsa.PrivateKey)
	if client.PrivateKey == nil {
		return nil, errors.New("cannot load private key")
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
	builder.AddOneLine("gateway/open/v1/msd/users")
	handers["gateway/open/v1/msd/users"] = gorouter.EndpointDefine{
		Handles:   nil,
		ApiFormer: NewApiFormGetUserConsents(client),
	}
	client.router.Init("", builder.Export(), handers)
	return client, nil
}

//RequestAccessKey request an access key from authCode
func (client *MiniAppClient) RequestAccessToken(partnerUserID string, authCode string) (*entity.AccessToken, error) {
	res, err := client.router.Request(client.domain, "gateway/open/v1/oauth/accessToken", map[string]interface{}{
		"partnerUserId": partnerUserID,
		"authCode":      authCode,
	})
	fmt.Println("partnerUserID:", partnerUserID, "\nauthCode:", authCode)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Base.Body)
	defer res.Close()
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	accessToken := &entity.AccessToken{}
	if err := json.Unmarshal(body, accessToken); err != nil {
		return nil, err
	}
	return accessToken, nil
}

//RequestGetUserConsents request for infomation
func (client *MiniAppClient) RequestGetUserConsents(partnerUserID string, attributes []common.Attribute, accessToken string) error {
	res, err := client.router.Request(client.domain, "gateway/open/v1/oauth/accessToken", map[string]interface{}{
		"partnerUserId": partnerUserID,
		"attributes":    attributes,
		"accessToken":   accessToken,
	})
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Base.Body)
	defer res.Close()
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
