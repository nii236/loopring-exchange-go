package loopring

import "errors"

var ErrNotImplemented = errors.New("not implemented")

type EDDSAHashSigner struct{}

func (s *EDDSAHashSigner) Hash() {}
func (s *EDDSAHashSigner) Sign() {}

type ECDSASigner struct{}

func (s *ECDSASigner) Sign() {}

type EIP712Hasher struct{}

func (s *EIP712Hasher) Hash() {}
