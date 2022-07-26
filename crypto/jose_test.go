package crypto

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
	"gotest.tools/assert"
	"testing"
)

func TestAddressJWT(t *testing.T) {
	addressStr := "0x01Fd2f9b86E31dF9e4D2Ab13f444Ac076A689db8"
	token, err := NewAddressJWT(ethcommon.HexToAddress(addressStr))
	if err != nil {
		t.Fatal(err)
	}
	addressParse, err := ParseAddressJWT(token)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, addressStr, addressParse.String())
}

