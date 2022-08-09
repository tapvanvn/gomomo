package crypto

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
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

func HeaderOPSignature(privateKey *rsa.PrivateKey, message string) ([]byte, error) {

	h := sha256.New()
	h.Write([]byte(message))
	d := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, d)
	if err != nil {
		return nil, err
	}

	//fmt.Printf("Signature in byte: %v\n\n", signature)
	return signature, nil

	//encodedSig := base64.StdEncoding.EncodeToString(signature)

	//fmt.Printf("Encoded signature: %v\n\n", encodedSig)
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
