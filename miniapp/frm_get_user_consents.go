package miniapp

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/tapvanvn/gomomo/common"
	"github.com/tapvanvn/gomomo/crypto"
	"github.com/tapvanvn/gorouter/v2"
)

type ApiFormGetUserConsents struct {
	*gorouter.ApiForm //init
	client            *MiniAppClient
}

func NewApiFormGetUserConsents(client *MiniAppClient) *ApiFormGetUserConsents {
	frm := &ApiFormGetUserConsents{
		client:  client,
		ApiForm: gorouter.NewGetForm(),
	}
	return frm
}

func (frm *ApiFormGetUserConsents) Request(domain string, path string, indexes map[string]interface{}) (*gorouter.ApiResponse, error) {

	frm.Params["partnerUserId"] = indexes["partnerUserId"].(string)

	attrs := indexes["attributes"].([]common.Attribute)
	if attrs == nil {
		return nil, errors.New("attributes is need to make getUserConsents request.")
	}
	frm.Params["fields"] = common.AttributesToString(attrs)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	dataSegments := []string{}
	for key, val := range frm.Params {
		dataSegments = append(dataSegments, fmt.Sprintf("%s=%s", key, val))
	}

	signature, _ := crypto.AESEncrypt(strings.Join(dataSegments, "&") + timestamp + frm.client.OpenSecret)

	frm.Headers["Authorization"] = "Bearer " + indexes["accessToken"].(string)
	frm.Headers["OP-Signature"] = hex.EncodeToString(signature)
	frm.Headers["M-Timestamp"] = timestamp

	return frm.ApiForm.Request(domain, path, indexes)
}
