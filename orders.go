package loopring

import (
	"net/http"
	"os"
)

var apiKey string = os.Getenv("LOOPRING_API_KEY")
var privateKey string = os.Getenv("LOOPRING_PRIVATE_KEY")
var publicKeyX string = os.Getenv("LOOPRING_PUBLIC_KEY_X")
var publicKeyY string = os.Getenv("LOOPRING_PUBLIC_KEY_Y")

func UpdateAPIKeyMessage(
	ed *EDDSAHashSigner,
	accountID string,
) *http.Request {
	accountIDParam := QueryParamPair{
		QueryParam: "accountID",
		Value:      accountID,
	}
	var signatureBase string = ConstructSignatureBase(http.MethodGet, "/apiKey", accountIDParam)

	hash := ed.Hash(signatureBase)

	signedHash, err := ed.Sign([]byte(privateKey), hash)
	if err != nil {
		panic(err)
	}
}

func CancelOrderMessage(
	ed *EDDSAHashSigner,
	accountID string,
	clientOrderId string,
) *http.Request {
	accountIDParam := QueryParamPair{
		QueryParam: "accountID",
		Value:      accountID,
	}
	clientOrderIdParam := QueryParamPair{
		QueryParam: "clientOrderId",
		Value:      clientOrderId,
	}
	var signatureBase = ConstructSignatureBase(http.MethodDelete, "/order", accountIDParam, clientOrderIdParam)

	hash := ed.Hash(signatureBase)

	signedHash, err := ed.Sign([]byte(privateKey), hash)
	if err != nil {
		panic(err)
	}
}

func SubmitOrderAMMMessage(
	*EDDSAHashSigner,
) {

}

func OffchainWithdrawalMessage(
	*EIP712Hasher,
	*ECDSASigner,
) *http.Request {
}

func InternalTransferMessage(
	*EIP712Hasher,
	*ECDSASigner,
) *http.Request {
	panic(ErrNotImplemented)
}
