package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"os"
	"encoding/asn1"
)

// LoadPrivateKey loads a PEM-encoded ECDSA private key (PKCS#8)
func LoadPrivateKey(keyPEM string) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(keyPEM))
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing the private key")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	ecdsaPrivateKey, ok := privateKey.(*ecdsa.PrivateKey)
	if !ok {
		return nil, errors.New("not an ECDSA private key")
	}
	return ecdsaPrivateKey, nil
}

// LoadCertificate loads a PEM-encoded certificate and extracts the public key
func LoadCertificate(certPEM string) (*ecdsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(certPEM))
	if block == nil || block.Type != "CERTIFICATE" {
		return nil, fmt.Errorf("failed to decode PEM block containing the certificate")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse certificate: %v", err)
	}

	publicKey, ok := cert.PublicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("certificate does not contain an ECDSA public key")
	}
	return publicKey, nil
}


// SignData signs the provided data using the user's private key
func SignData(privateKey *ecdsa.PrivateKey, data []byte) (string, error) {
	// Compute the SHA-256 hash of the data
	hash := sha256.Sum256(data)

	// Sign the hash using the user's private key
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return "", fmt.Errorf("failed to sign data: %v", err)
	}

	// Encode the r and s values in ASN.1 DER format
	type ecdsaSignature struct {
		R, S *big.Int
	}
	encodedSignature, err := asn1.Marshal(ecdsaSignature{r, s})
	if err != nil {
		return "", fmt.Errorf("failed to encode signature: %v", err)
	}

	return base64.StdEncoding.EncodeToString(encodedSignature), nil
}

// ValidateSignature verifies that the provided signature matches the data
func ValidateSignature(publicKey *ecdsa.PublicKey, data []byte, signatureBase64 string) error {
	// Decode the Base64-encoded signature
	encodedSignature, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		return fmt.Errorf("failed to decode signature: %v", err)
	}

	// Parse the ASN.1 DER-encoded signature to extract r and s
	type ecdsaSignature struct {
		R, S *big.Int
	}
	var sig ecdsaSignature
	_, err = asn1.Unmarshal(encodedSignature, &sig)
	if err != nil {
		return fmt.Errorf("failed to decode ASN.1 signature: %v", err)
	}

	// Compute the SHA-256 hash of the data
	hash := sha256.Sum256(data)

	// Verify the signature
	if !ecdsa.Verify(publicKey, hash[:], sig.R, sig.S) {
		return errors.New("signature verification failed")
	}

	return nil
}



// this contain an example private key
// func main() {
// 	// Replace with your actual private key (in PKCS#8 PEM format)
// 	privateKeyPEM := `-----BEGIN PRIVATE KEY-----
// MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgJbcUINsTFhfpg1Cn
// IYT/jstATb/wjxARJhT5sgb7fgKhRANCAAReZjIwUl8z1jia8l2sIPY4vml2JCxb
// npvg14eAkp267Utso214cg642Lo/B2TAjLub+y3RDU255Ar5nGl1gWcD
// -----END PRIVATE KEY-----`

// 	// Replace with your actual user certificate
// 	userCertPEM := `-----BEGIN CERTIFICATE-----
// MIICsDCCAlagAwIBAgIUGHqHN/EyZwXxz6gkCmUMpSIP8CIwCgYIKoZIzj0EAwIw
// aDELMAkGA1UEBhMCVVMxFzAVBgNVBAgTDk5vcnRoIENhcm9saW5hMRQwEgYDVQQK
// EwtIeXBlcmxlZGdlcjEPMA0GA1UECxMGRmFicmljMRkwFwYDVQQDExBmYWJyaWMt
// Y2Etc2VydmVyMB4XDTI0MTIwMjIxMjYwMFoXDTI1MTIwMjIxMzEwMFowbjELMAkG
// A1UEBhMCVVMxFzAVBgNVBAgTDk5vcnRoIENhcm9saW5hMRQwEgYDVQQKEwtIeXBl
// cmxlZGdlcjEcMA0GA1UECxMGY2xpZW50MAsGA1UECxMEb3JnMTESMBAGA1UEAxMJ
// dGhhcmluZHUxMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEXmYyMFJfM9Y4mvJd
// rCD2OL5pdiQsW56b4NeHgJKduu1LbKNteHIOuNi6PwdkwIy7m/st0Q1NueQK+Zxp
// dYFnA6OB1zCB1DAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH/BAIwADAdBgNVHQ4E
// FgQU7b2VmlKZEq/aNEvMbldjIZCrpiEwHwYDVR0jBBgwFoAUCqCMEy1VPWuLfKE1
// f2nLvU+G0LswEgYDVR0RBAswCYIHdmFncmFudDBgBggqAwQFBgcIAQRUeyJhdHRy
// cyI6eyJoZi5BZmZpbGlhdGlvbiI6Im9yZzEiLCJoZi5FbnJvbGxtZW50SUQiOiJ0
// aGFyaW5kdTEiLCJoZi5UeXBlIjoiY2xpZW50In19MAoGCCqGSM49BAMCA0gAMEUC
// IQCzCWfTwajswmW5lbouwPF1IaEy4laN2LOnQadxqnfyiQIgVYbTpvrwErkfIVKQ
// rHA4IPtKPrDV/jGeBhUKiKq3a9M=
// -----END CERTIFICATE-----`

// 	data := []byte("example data to sign")

// 	// Load the private key
// 	privateKey, err := LoadPrivateKey(privateKeyPEM)
// 	if err != nil {
// 		fmt.Println("Failed to load private key:", err)
// 		os.Exit(1)
// 	}

// 	// Load the user's public key from the certificate
// 	publicKey, err := LoadCertificate(userCertPEM)
// 	if err != nil {
// 		fmt.Println("Failed to load user certificate:", err)
// 		os.Exit(1)
// 	}
// 	fmt.Println("Public Key Read:", publicKey)

// 	// Sign the data
// 	signature, err := SignData(privateKey, data)
// 	if err != nil {
// 		fmt.Println("Failed to sign data:", err)
// 		os.Exit(1)
// 	}
// 	fmt.Println("Signature (Base64):", signature)

// 	// Validate the signature
// 	err = ValidateSignature(publicKey, data, signature)
// 	if err != nil {
// 		fmt.Println("Signature validation failed:", err)
// 		os.Exit(1)
// 	}
// 	fmt.Println("Signature validation succeeded.")
// }
