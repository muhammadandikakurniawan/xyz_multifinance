package crypto

type PkgCrypto interface {
	Encrypt(plainTxt []byte) ([]byte, error)
	Decrypt(encryptedTxt []byte) ([]byte, error)
}
