package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func NewAesCbc(iv, key string) (res AesCBCCrypto, err error) {
	if len(iv) != 16 {
		err = fmt.Errorf("invalid iv : %s", iv)
		return
	}
	if len(key) != 32 {
		err = fmt.Errorf("invalid key : %s", key)
		return
	}
	res = AesCBCCrypto{
		initialVector: []byte(iv),
		key:           []byte(key),
	}
	return
}

type AesCBCCrypto struct {
	initialVector []byte
	key           []byte
}

func (o AesCBCCrypto) Encrypt(plainTxt []byte) (result []byte, err error) {

	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
		}
	}()

	block, err := aes.NewCipher(o.key)
	if err != nil {
		return
	}

	bPlaintext := PKCS5Padding(plainTxt, block.BlockSize())
	result = make([]byte, len(bPlaintext))
	cbc := cipher.NewCBCEncrypter(block, o.initialVector)
	cbc.CryptBlocks(result, bPlaintext)
	return
}

func (o AesCBCCrypto) Decrypt(encryptedTxt []byte) (result []byte, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
		}
	}()
	block, err := aes.NewCipher(o.key)
	if err != nil {
		return
	}
	cbc := cipher.NewCBCDecrypter(block, o.initialVector)
	result = make([]byte, len(encryptedTxt))
	cbc.CryptBlocks(result, encryptedTxt)
	result = PKCS5UnPadding(result)
	return
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
