package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
)

func GenerateSecretKey() (key []byte, iv []byte, err error) {
	key = make([]byte, 24)
	_, err = rand.Read(key)
	if err != nil {
		return
	}
	iv = make([]byte, 16)
	_, err = rand.Read(key)
	return
}
func AESEncrypt(message string) ([]byte, error) {

	key, iv, err := GenerateSecretKey()
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if message == "" {
		return nil, errors.New("AESEncrypt: message cannot be empty")
	}
	ecb := cipher.NewCBCEncrypter(block, iv)
	content := []byte(message)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	return crypted, nil
}

//OP-Signature = base64UrlEncode(sha256withrsa(data + M-timestamp + openSecretKey))

/*
func AESDecrypt(crypt []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if len(crypt) == 0 {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCDecrypter(block, []byte(initialVector))
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)

	return PKCS5Trimming(decrypted)
}
*/
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
