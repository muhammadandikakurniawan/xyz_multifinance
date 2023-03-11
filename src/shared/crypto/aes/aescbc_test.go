package aes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_aesCbc(t *testing.T) {
	iv := "acfa7a047800b2f2"
	key := "acfa7a047800b2f221f2c4f7d626eafb"
	plainTxt := "test-plain-text-123-test-plain-text-123asdasdasd!@#$%^&*()"

	aesCrypto, _ := NewAesCbc(iv, key)

	encryptedTxt, err := aesCrypto.Encrypt([]byte(plainTxt))
	if err != nil {
		t.Fail()
	}
	decryptRes, err := aesCrypto.Decrypt(encryptedTxt)
	if err != nil {
		t.Fail()
	}

	decryptredTxt := string(decryptRes)
	assert.Equal(t, plainTxt, decryptredTxt)
}
