package auth

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ed25519"
)

// verifySignature verifies the Ed25519 signature for given data and public key
func verifySignature(data string, signatureBase64 string, publicKeyBase58 string) bool {
	// Hash the data using SHA-512
	hash := sha512.Sum512([]byte(data))

	// Decode Base64 signature
	signature, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		fmt.Println("Failed to decode Base64 signature:", err)
		return false
	}

	// Decode Base58 public key
	publicKeyBytes := base58.Decode(publicKeyBase58)

	// Check if the length of the public key is valid
	if len(publicKeyBytes) != ed25519.PublicKeySize {
		fmt.Println("Invalid public key size")
		return false
	}

	// Create an Ed25519 public key from bytes
	pubKey := ed25519.PublicKey(publicKeyBytes)

	// Verify the signature
	return ed25519.Verify(pubKey, hash[:], signature)
}

// func main() {
//     // Example usage
//     data := "your_data_here"
//     signatureBase64 := "your_signature_base64_here"
//     publicKeyBase58 := "your_public_key_base58_here"

//     isValid := verifySignature(data, signatureBase64, publicKeyBase58)
//     fmt.Println("Signature valid:", isValid)
// }
