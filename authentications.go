package flip

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"os"

	"github.com/fari-99/go-flip/constants"
)

func CheckCallback(token string) (valid bool, err error) {
	flipValidationToken := os.Getenv("FLIP_VALIDATION_TOKEN")
	if flipValidationToken == "" {
		panic("please set flip validation key [FLIP_VALIDATION_TOKEN] to ENV")
	}

	if token != flipValidationToken {
		return false, nil
	}

	return true, nil
}

func GenerateRSAKey() (privateKey string, publicKey string, err error) {
	flipPrivateKey := os.Getenv("FLIP_PRIVATE_KEY") // in base64
	flipPublicKey := os.Getenv("FLIP_PUBLIC_KEY")   // in base64

	if flipPrivateKey != "" || flipPublicKey != "" {
		return flipPrivateKey, flipPublicKey, nil
	}

	randomness := rand.Reader
	key, err := rsa.GenerateKey(randomness, constants.RSAKeyBitTotal)
	if err != nil {
		return "", "", err
	}

	keyPrivate := x509.MarshalPKCS1PrivateKey(key)
	keyPublic := x509.MarshalPKCS1PublicKey(&key.PublicKey)

	// encode to base 64
	flipPrivateKey = base64.RawURLEncoding.EncodeToString(keyPrivate)
	flipPublicKey = base64.RawURLEncoding.EncodeToString(keyPublic)

	return flipPrivateKey, flipPublicKey, nil
}

func GetSignature(payload any) (signature string, err error) {
	payloadMarshal, _ := json.Marshal(payload)
	flipPrivateKey, _, err := GenerateRSAKey()
	if err != nil {
		return "", err
	}

	privateKeyMarshal, err := base64.RawURLEncoding.DecodeString(flipPrivateKey)
	if err != nil {
		return "", err
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyMarshal)
	if err != nil {
		return "", err
	}

	result, err := rsa.SignPKCS1v15(nil, privateKey, 0, payloadMarshal)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
