package miniapp

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/tapvanvn/gomomo/crypto"
	"github.com/tapvanvn/gorouter/v2"
)

type ApiFormAccessKey struct {
	*gorouter.ApiForm //init
	client            *MiniAppClient
}

func NewApiFormAccessToken(client *MiniAppClient) *ApiFormAccessKey {
	frm := &ApiFormAccessKey{
		client:  client,
		ApiForm: gorouter.NewGetForm(),
	}
	return frm
}

func (frm *ApiFormAccessKey) Request(domain string, path string, indexes map[string]interface{}) (*gorouter.ApiResponse, error) {

	frm.Params["partnerUserId"] = indexes["partnerUserId"].(string)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	dataSegments := []string{}
	for key, val := range frm.Params {
		dataSegments = append(dataSegments, fmt.Sprintf("%s=%s", key, val))
	}
	signDataRaw := strings.Join(dataSegments, "&") + timestamp + frm.client.OpenSecret
	signature, _ := crypto.HeaderOPSignature(frm.client.PrivateKey, signDataRaw)

	fmt.Println("signDataRaw:", signDataRaw, "\n\tsign:", base64.URLEncoding.EncodeToString(signature))

	frm.Headers["Authorization"] = "Bearer " + indexes["authCode"].(string)
	frm.Headers["OP-Signature"] = base64.URLEncoding.EncodeToString(signature) //hex.EncodeToString(signature)
	frm.Headers["M-Timestamp"] = timestamp

	return frm.ApiForm.Request(domain, path, indexes)
}
