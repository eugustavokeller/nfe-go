package services

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"os"

	"software.sslmate.com/src/go-pkcs12"
)

func CarregarCertificado(caminhoCert, senha string) (*rsa.PrivateKey, *x509.Certificate, error) {
	pfxData, err := os.ReadFile(caminhoCert)
	if err != nil {
		return nil, nil, fmt.Errorf("erro ao carregar o arquivo .pfx: %w", err)
	}

	privateKey, cert, err := pkcs12.Decode(pfxData, senha)
	if err != nil {
		return nil, nil, fmt.Errorf("erro ao decodificar o arquivo .pfx: %w", err)
	}

	rsaKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, nil, fmt.Errorf("a chave privada não é do tipo RSA")
	}

	return rsaKey, cert, nil
}

func AssinarXML(xmlString string, privateKey *rsa.PrivateKey) (string, error) {
	hasher := sha1.New()
	_, err := hasher.Write([]byte(xmlString))
	if err != nil {
		return "", fmt.Errorf("erro ao calcular hash do XML: %w", err)
	}
	hash := hasher.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA1, hash)
	if err != nil {
		return "", fmt.Errorf("erro ao assinar o hash: %w", err)
	}

	signatureBase64 := base64.StdEncoding.EncodeToString(signature)
	return signatureBase64, nil
}
