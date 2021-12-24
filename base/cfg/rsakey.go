package cfg

import (
	"crypto/rsa"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
)

var (
	RsaPrivateKey *rsa.PrivateKey
	RsaPublicKey  *rsa.PublicKey
)

// InitRsaKey init rsa key object
func InitRsaKey() {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(readRsaKeyFile(AppConfObj.RsaPrivateKey))
	if err != nil {
		Mlog.Panic("failed to parse rsa private key, error message: ", err.Error())
		return
	}
	RsaPrivateKey = privateKey

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(readRsaKeyFile(AppConfObj.RsaPublicKey))
	if err != nil {
		Mlog.Panic("failed to parse rsa public key, error message: ", err.Error())
		return
	}
	RsaPublicKey = publicKey
}

// readRsaKeyFile read the raw contents of the rsa key file
func readRsaKeyFile(filepath string) []byte {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		Mlog.Panic("failed to read rsa key file, error message: ", err.Error())
		return nil
	}

	return data
}
