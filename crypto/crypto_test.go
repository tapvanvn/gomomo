package crypto_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/tapvanvn/gomomo/crypto"
)

func TestEncrypt(t *testing.T) {
	encrypted, err := crypto.AESEncrypt("test")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(encrypted))
}
