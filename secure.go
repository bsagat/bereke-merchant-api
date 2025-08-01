package bereke_merchant

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"net/http"
	"os"
)

func (a *api) signAndSetHeaders(req *http.Request, bodyStr string) error {
	xHash := calculateSHA256Base64(bodyStr)

	privKey, err := loadEncryptedPrivateKey(a.certPath, []byte(a.certPassphrase))
	if err != nil {
		return err
	}

	xSignature, err := signSHA256(privKey, xHash)
	if err != nil {
		return err
	}

	req.Header.Set("X-Hash", xHash)
	req.Header.Set("X-Signature", xSignature)
	return nil
}

func calculateSHA256Base64(data string) string {
	hash := sha256.Sum256([]byte(data))
	return base64.StdEncoding.EncodeToString(hash[:])
}

// Декодирует и расшифровывает PEM-зашифрованный ключ
func loadEncryptedPrivateKey(path string, password []byte) (*rsa.PrivateKey, error) {
	pemBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("failed to decode PEM")
	}

	der, err := x509.DecryptPEMBlock(block, password)
	if err != nil {
		return nil, err
	}

	key, err := x509.ParsePKCS1PrivateKey(der)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// Подписывает SHA256-хэш
func signSHA256(privateKey *rsa.PrivateKey, hashBase64 string) (string, error) {
	hashBytes, err := base64.StdEncoding.DecodeString(hashBase64)
	if err != nil {
		return "", err
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashBytes)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}
