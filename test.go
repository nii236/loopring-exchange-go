package loopring

import (
	"crypto/rand"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/katzenpost/core/crypto/eddsa"
)

func main() {
	keyPair, err := eddsa.NewKeypair(rand.Reader)
	if err != nil {
		panic(err)
	}

	fmt.Println("Private Key <", len(keyPair.Bytes()), ">: ", keyPair.Bytes())
	fmt.Println("Public Key <", len(keyPair.PublicKey().Bytes()), ">: ", keyPair.PublicKey().Bytes())

	var newPrivateKey = "2bf7d29ae0293dd8b7538681341934a26ec5c98bd2f8c58e4d67bbede05d1b7"
	var newPublicKey = "2cba32a1bdb218aa380c1e42a8d7dcbf345ec9c914ecf37e34006b34260d198"
	keyPair.FromBytes([]byte(newPrivateKey))
	keyPair.PublicKey().FromBytes([]byte(newPublicKey))

	fmt.Println("Private Key <", len(keyPair.Bytes()), ">: ", keyPair.Bytes())
	fmt.Println("Public Key <", len(keyPair.PublicKey().Bytes()), ">: ", keyPair.PublicKey().Bytes())

	var hash = crypto.Keccak256Hash([]byte("bruh"))
	fmt.Println(hash)
	fmt.Println(hash.String())

	var newKeyPair = GenerateLoopringKeyPair()
	fmt.Println(newKeyPair.PublicKey())
}
