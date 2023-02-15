package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"strings"
)

var (
	ErrKeyMustBePEMEncoded = errors.New("invalid key: Key must be a PEM encoded PKCS1 or PKCS8 key")
	ErrNotRSAPrivateKey    = errors.New("key is not a valid RSA private key")
	ErrNotRSAPublicKey     = errors.New("key is not a valid RSA public key")
)

const (
	publicKeyPemStart  = "-----BEGIN PUBLIC KEY-----"
	publicKeyPemEnd    = "-----END PUBLIC KEY-----"
	privateKeyPemStart = "-----BEGIN PRIVATE KEY-----"
	privateKeyPemEnd   = "-----END PRIVATE KEY-----"
)

func ParseRSAPublicKey(key string) (*rsa.PublicKey, error) {
	if !strings.HasPrefix(key, publicKeyPemStart) {
		key = publicKeyPemStart + "\n" + key + "\n" + publicKeyPemEnd
	}

	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode([]byte(key)); block == nil {
		return nil, ErrKeyMustBePEMEncoded
	}

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		if cert, err := x509.ParseCertificate(block.Bytes); err == nil {
			parsedKey = cert.PublicKey
		} else {
			return nil, err
		}
	}

	var pkey *rsa.PublicKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PublicKey); !ok {
		return nil, ErrNotRSAPublicKey
	}

	return pkey, nil
}

func ParseRSAPrivateKey(key string) (*rsa.PrivateKey, error) {
	if !strings.HasPrefix(key, privateKeyPemStart) {
		key = privateKeyPemStart + "\n" + key + "\n" + privateKeyPemEnd
	}

	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode([]byte(key)); block == nil {
		return nil, ErrKeyMustBePEMEncoded
	}

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
			return nil, err
		}
	}

	var pkey *rsa.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PrivateKey); !ok {
		return nil, ErrNotRSAPrivateKey
	}

	return pkey, nil
}

func SignSHA256WithRSA(priv *rsa.PrivateKey, s string) (string, error) {
	hashed := sha256.Sum256([]byte(s))
	signer, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signer), nil
}

func VerifySHA256WithRSA(pub *rsa.PublicKey, s, sign string) error {
	inSign, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	hashed := sha256.Sum256([]byte(s))
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], inSign)
}

func RsaEncrypt(pub *rsa.PublicKey, data []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}

func RsaDecrypt(priv *rsa.PrivateKey, data []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, priv, data)
}
