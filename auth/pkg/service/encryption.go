package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"math/big"
	"time"

	"golang.org/x/crypto/ssh"
)

// Generate SSH public key
func generateSSHPublicKey() (string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", err
	}

	publicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", err
	}

	return string(ssh.MarshalAuthorizedKey(publicKey)), nil
}

// Encrypt data using SSH public key
func encryptWithSSH(data []byte, publicKey string) ([]byte, error) {
	parsed, _, _, _, err := ssh.ParseAuthorizedKey([]byte(publicKey))
    if err != nil {
        panic(err)
    }

	parsedCryptoKey := parsed.(ssh.CryptoPublicKey)
	pubCrypto := parsedCryptoKey.CryptoPublicKey()
	pub := pubCrypto.(*rsa.PublicKey)

	encryptedData, err := rsa.EncryptPKCS1v15(
        rand.Reader,
        pub,
        []byte(data),
    )
    if err != nil {
        panic(err)
    }

	return encryptedData, nil
}

// Encrypt data using SSL/TLS
func encryptWithSSL(data []byte) ([]byte, error) {
	// Generate a self-signed SSL/TLS certificate
	cert, _, err := generateSSLCertificate() // return cert, private key
	if err != nil {
		return nil, err
	}

	// Create a TLS configuration
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	// Encrypt the data using SSL/TLS
	encryptedData, err := tlsEncrypt(data, config)
	if err != nil {
		return nil, err
	}

	return encryptedData, nil
}

// Generate a self-signed SSL/TLS certificate
func generateSSLCertificate() (tls.Certificate, *rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return tls.Certificate{}, nil, err
	}

	template := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0), // 10 years validity
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return tls.Certificate{}, nil, err
	}

	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})

	certificate, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return tls.Certificate{}, nil, err
	}

	return certificate, privateKey, nil
}

// Encrypt data using SSL/TLS
func tlsEncrypt(data []byte, config *tls.Config) ([]byte, error) {
	conn, err := tls.Dial("tcp", "localhost:443", config)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	_, err = conn.Write(data)
	if err != nil {
		return nil, err
	}

	encryptedData, err := ioutil.ReadAll(conn)
	if err != nil {
		return nil, err
	}

	return encryptedData, nil
}