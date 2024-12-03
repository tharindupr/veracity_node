package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)


caCertPEM := `-----BEGIN CERTIFICATE-----
MIICFjCCAb2gAwIBAgIUIswrc/lQfEJHNc+G+JEMlGH3nmgwCgYIKoZIzj0EAwIw
aDELMAkGA1UEBhMCVVMxFzAVBgNVBAgTDk5vcnRoIENhcm9saW5hMRQwEgYDVQQK
EwtIeXBlcmxlZGdlcjEPMA0GA1UECxMGRmFicmljMRkwFwYDVQQDExBmYWJyaWMt
Y2Etc2VydmVyMB4XDTI0MTIwMjIwNDcwMFoXDTM5MTEyOTIwNDcwMFowaDELMAkG
A1UEBhMCVVMxFzAVBgNVBAgTDk5vcnRoIENhcm9saW5hMRQwEgYDVQQKEwtIeXBl
cmxlZGdlcjEPMA0GA1UECxMGRmFicmljMRkwFwYDVQQDExBmYWJyaWMtY2Etc2Vy
dmVyMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEln7DvZoHL3RqUZhMGBO+rP5s
Am4jblnkzuflBgqogzfdbUiJsn1CrUbKaGd3E5QSNIqqtbHNyhb2+t4g9VfHBKNF
MEMwDgYDVR0PAQH/BAQDAgEGMBIGA1UdEwEB/wQIMAYBAf8CAQEwHQYDVR0OBBYE
FAqgjBMtVT1ri3yhNX9py71PhtC7MAoGCCqGSM49BAMCA0cAMEQCIHENZKSRN2SN
Mer1KzqjhKKb8sdEoedWsDljYxtUYW1rAiAe/9RKo0R8nfon8wu66dRxhNhFumGw
OYMsXrAmO71IhQ==
-----END CERTIFICATE-----`

// LoadCertificate loads a PEM-encoded certificate
func LoadCertificate(certPEM string) (*x509.Certificate, error) {
	block, _ := pem.Decode([]byte(certPEM))
	if block == nil || block.Type != "CERTIFICATE" {
		return nil, fmt.Errorf("failed to decode PEM block containing the certificate")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse certificate: %v", err)
	}
	return cert, nil
}

// VerifyUserCertificate verifies that the user certificate is signed by the CA
func VerifyUserCertificate(userCertPEM, caCertPEM string) error {
	// Load the user certificate
	userCert, err := LoadCertificate(userCertPEM)
	if err != nil {
		return fmt.Errorf("failed to load user certificate: %v", err)
	}

	// Load the CA certificate
	caCert, err := LoadCertificate(caCertPEM)
	if err != nil {
		return fmt.Errorf("failed to load CA certificate: %v", err)
	}

	// Create a certificate pool and add the CA certificate
	certPool := x509.NewCertPool()
	certPool.AddCert(caCert)

	// Verify the user certificate
	opts := x509.VerifyOptions{
		Roots: certPool,
	}

	if _, err := userCert.Verify(opts); err != nil {
		return fmt.Errorf("certificate verification failed: %v", err)
	}

	return nil
}




// func main() {
// 	// Replace with your actual certificate strings
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

// 	caCertPEM := `-----BEGIN CERTIFICATE-----
// MIICFjCCAb2gAwIBAgIUIswrc/lQfEJHNc+G+JEMlGH3nmgwCgYIKoZIzj0EAwIw
// aDELMAkGA1UEBhMCVVMxFzAVBgNVBAgTDk5vcnRoIENhcm9saW5hMRQwEgYDVQQK
// EwtIeXBlcmxlZGdlcjEPMA0GA1UECxMGRmFicmljMRkwFwYDVQQDExBmYWJyaWMt
// Y2Etc2VydmVyMB4XDTI0MTIwMjIwNDcwMFoXDTM5MTEyOTIwNDcwMFowaDELMAkG
// A1UEBhMCVVMxFzAVBgNVBAgTDk5vcnRoIENhcm9saW5hMRQwEgYDVQQKEwtIeXBl
// cmxlZGdlcjEPMA0GA1UECxMGRmFicmljMRkwFwYDVQQDExBmYWJyaWMtY2Etc2Vy
// dmVyMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEln7DvZoHL3RqUZhMGBO+rP5s
// Am4jblnkzuflBgqogzfdbUiJsn1CrUbKaGd3E5QSNIqqtbHNyhb2+t4g9VfHBKNF
// MEMwDgYDVR0PAQH/BAQDAgEGMBIGA1UdEwEB/wQIMAYBAf8CAQEwHQYDVR0OBBYE
// FAqgjBMtVT1ri3yhNX9py71PhtC7MAoGCCqGSM49BAMCA0cAMEQCIHENZKSRN2SN
// Mer1KzqjhKKb8sdEoedWsDljYxtUYW1rAiAe/9RKo0R8nfon8wu66dRxhNhFumGw
// OYMsXrAmO71IhQ==
// -----END CERTIFICATE-----`

// 	err := VerifyUserCertificate(userCertPEM, caCertPEM)
// 	if err != nil {
// 		fmt.Println("Certificate verification failed:", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println("Certificate verification succeeded.")
// }
