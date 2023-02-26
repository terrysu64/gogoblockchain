package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/ripemd160"
	"github.com/btcsuite/btcutil/base58"
	"crypto/sha256"
)

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	blockchainAddress string
}
	 

func NewWallet() *Wallet {
	//1. create ECDSA private and public keys (32 and 64 bytes respectively)
	w := new(Wallet)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privateKey = privateKey
	w.publicKey = &privateKey.PublicKey

	//2. preform SHA-256 hashing on the public key (32 bytes)
	h2 := sha256.New()
	h2.Write(w.publicKey.X.Bytes())
	h2.Write(w.publicKey.Y.Bytes())
	hashedPublicKey := h2.Sum(nil)

	//3. preform RIPEMD-160 hashing on the result of SHA-256 (20 bytes)
	h3 := ripemd160.New()
	h3.Write(hashedPublicKey)
	ripeHashedPublicKey := h3.Sum(nil)

	//4. add version byte in front of RIPEMD-160 hash (0x00 for Main Network)
	versionedRipeHashedPublicKey := make([]byte, 21)
	versionedRipeHashedPublicKey[0] = 0x00
	copy(versionedRipeHashedPublicKey[1:], ripeHashedPublicKey[:])

	//5. perform SHA-256 hash on the extended RIPEMD-160 result
	h4 := sha256.New()
	h4.Write(versionedRipeHashedPublicKey)
	hash1 := h4.Sum(nil)

	//6. perform SHA-256 hash on the result of the previous SHA-256 hash
	h5 := sha256.New()
	h5.Write(hash1)
	hash2 := h5.Sum(nil)

	//7. take the first 4 bytes of the second SHA-256 hash. This is the address checksum
	checksum := hash2[:4]

	//8. add the 4 checksum bytes from stage 7 at the end of extended RIPEMD-160 hash from stage 4. This is the 25-byte binary Bitcoin Address.
	binaryAddress := make([]byte, 25)
	copy(binaryAddress[:], versionedRipeHashedPublicKey[:])
	copy(binaryAddress[21:], checksum[:])

	//9. convert the result from a byte string into a base58 string using Base58Check encoding. This is the most commonly used Bitcoin Address format
	w.blockchainAddress = base58.Encode(binaryAddress)
	return w
}

func (w *Wallet) BlockchainAddress() string {
	return w.blockchainAddress
}

func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

//D is a component of the private key
func (w *Wallet) PrivateKeyString() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

//X and Y are mathematical components of the public key
func (w *Wallet) PublicKeyString() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}