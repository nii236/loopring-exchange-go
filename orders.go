package loopring

import (
	"net/http"
	"os"
)

var apiKey string = os.Getenv("LOOPRING_API_KEY")
var privateKey string = os.Getenv("LOOPRING_PRIVATE_KEY")
var publicKeyX string = os.Getenv("LOOPRING_PUBLIC_KEY_X")
var publicKeyY string = os.Getenv("LOOPRING_PUBLIC_KEY_Y")

type Order struct{}

func (ed *EDDSAHashSigner) serialize_data(order Order) []int {
	// Extract Order Data
	panic(ErrNotImplemented)
}

func UpdateAPIKeyMessage(
	ed *EDDSAHashSigner,
	accountID string,
) *http.Request {
	panic(ErrNotImplemented)
}

func CancelOrderMessage(
	ed *EDDSAHashSigner,
	accountID string,
	clientOrderId string,
) *http.Request {
	panic(ErrNotImplemented)
}

func SubmitOrderAMMMessage(
	*EDDSAHashSigner,
) {
	panic(ErrNotImplemented)
}

func OffchainWithdrawalMessage(
	*EIP712Hasher,
	*ECDSASigner,
) *http.Request {
	panic(ErrNotImplemented)
}

func InternalTransferMessage(
	*EIP712Hasher,
	*ECDSASigner,
) *http.Request {
	panic(ErrNotImplemented)
}
