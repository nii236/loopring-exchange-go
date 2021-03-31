package loopring

import (
	"crypto"
	"crypto/ed25519"
)

type EDDSAHashSigner struct{}

func (s *EDDSAHashSigner) Hash(data string) []byte {
	h := crypto.SHA256.New()
	h.Write([]byte(data))
	hash := h.Sum(nil)
	return hash
}
func (s *EDDSAHashSigner) Sign(privateKey []byte, msg []byte) ([]byte, error) {
	signatureBytes := ed25519.Sign(privateKey, msg)
	return signatureBytes, nil
}

type ECDSASigner struct{}

func (s *ECDSASigner) Sign() {}

type EIP712Hasher struct{}

func (s *EIP712Hasher) Hash() {}
