package loopring

import (
	"net/http"
)

func UpdateAPIKeyMessage(
	*EDDSAHashSigner,
) *http.Request {
	panic(ErrNotImplemented)
}
func CancelOrderMessage(
	*EDDSAHashSigner,
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
