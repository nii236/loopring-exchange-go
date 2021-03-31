package loopring

import (
	"net/http"
)

func UpdateAPIKeyMessage(
	ed *EDDSAHashSigner,
	accountID string,
	privateKey string,
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
	privateKey string,
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
	panic(ErrNotImplemented)
}
func InternalTransferMessage(
	*EIP712Hasher,
	*ECDSASigner,
) *http.Request {
	panic(ErrNotImplemented)
}
