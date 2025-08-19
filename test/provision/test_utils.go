package test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

func generateSshKeys() (privKey []byte, pubKey []byte, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate key: %w", err) 
	}

	privKey = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	pub, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get public key: %w", err)
	}
	pubKey = ssh.MarshalAuthorizedKey(pub)

	return privKey, pubKey, nil
}

func storeSshKeys(privateKey []byte, publicKey []byte, targetPath string) (publicKeyPath string, err error) {
	privateKeyPath := fmt.Sprintf("%s/id_rsa", targetPath)
	publicKeyPath = fmt.Sprintf("%s/id_rsa.pub", targetPath)

	if err = ioutil.WriteFile(privateKeyPath, privateKey, 0600); err != nil {
		return "", fmt.Errorf("Failed to write private key: %w", err)
	}

	if err = ioutil.WriteFile(publicKeyPath, publicKey, 0644); err != nil {
		return "", fmt.Errorf("Failed to write public key: %w", err)
	}

	return publicKeyPath, nil
}
