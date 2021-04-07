package loopring

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"

	"github.com/iden3/go-iden3-crypto/poseidon"
	"github.com/katzenpost/core/crypto/eddsa"
)

type EDDSAHashSigner struct{}

func (s *EDDSAHashSigner) Hash(inputBI [6]*big.Int) *big.Int {
	hash, err := poseidon.PoseidonHash(inputBI)
	if err != nil {
		panic(err)
	}

	return hash
}
func (s *EDDSAHashSigner) Sign(privateKey *eddsa.PrivateKey, msg []byte) []byte {
	signatureBytes := privateKey.Sign(msg)
	return signatureBytes
}

type ECDSASigner struct{}

func (s *ECDSASigner) Sign(msg []byte) (r []byte, sig []byte) {
	curve := elliptic.P256()
	randReader := rand.Reader
	newPrivateKey, err := ecdsa.GenerateKey(curve, randReader)
	if err != nil {
		panic(err)
	}

	ro, si, err := ecdsa.Sign(randReader, newPrivateKey, msg)
	if err != nil {
		panic(err)
	}

	return ro.Bytes(), si.Bytes()
}

type EIP712Hasher struct{}

func (s *EIP712Hasher) Hash() {

}
