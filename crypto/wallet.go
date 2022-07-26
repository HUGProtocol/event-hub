package crypto

import (
	"github.com/ethereum/go-ethereum/accounts"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func VerifySig(msg string, sigHex string, addr ethcommon.Address) bool {
	data := []byte(msg)
	sig, err := hexutil.Decode(sigHex)
	if err != nil {
		return false
	}
	if len(sig) != crypto.SignatureLength {
		return false
	}
	if sig[crypto.RecoveryIDOffset] != 27 && sig[crypto.RecoveryIDOffset] != 28 {
		return false
	}
	sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1

	rpk, err := crypto.SigToPub(accounts.TextHash(data), sig)
	if err != nil {
		return false
	}
	signer := crypto.PubkeyToAddress(*rpk)
	if signer != addr {
		return false
	}
	return true
}
