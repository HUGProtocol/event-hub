package crypto

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/MicahParks/keyfunc"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const Web3authJwksUrl = "https://api.openlogin.com/jwks"
var HmacKey = ""
const HmacExpirationDuration = time.Hour

var Web3authJwks *keyfunc.JWKS

func init() {
	var err error
	Web3authJwks, err = keyfunc.Get(Web3authJwksUrl, keyfunc.Options{})
	if err != nil {
		panic(fmt.Sprintf("get web3auth jwts from url %v err %v\n", Web3authJwksUrl, err.Error()))
	}
}

type Web3AuthProfile struct {
	Name         string
	ProfileImage string
	VerifierId   string
	PublicKey    *ecdsa.PublicKey
}

func ParseWeb3AuthJose(jwtStr string) (Web3AuthProfile, error) {
	token, err := jwt.Parse(jwtStr, Web3authJwks.Keyfunc)
	if err != nil {
		return Web3AuthProfile{}, err
	}
	tokenMap := token.Claims.(jwt.MapClaims)
	profile := Web3AuthProfile{}
	if name, ok := tokenMap["name"]; ok {
		profile.Name = fmt.Sprintf("%v", name)
	}
	if image, ok := tokenMap["profileImage"]; ok {
		profile.ProfileImage = fmt.Sprintf("%v", image)
	}
	if id, ok := tokenMap["verifierId"]; ok {
		profile.VerifierId = fmt.Sprintf("%v", id)
	}
	if wallets, ok := tokenMap["wallets"]; ok {
		if walletsList, os := wallets.([]interface{}); os {
			if len(walletsList) != 0 {
				wallet := walletsList[0]
				if walletMap, is := wallet.(map[string]interface{}); is {
					pkStr := fmt.Sprintf("%v", walletMap["public_key"])
					pk, e := crypto.DecompressPubkey(ethcommon.Hex2Bytes(pkStr))
					if e == nil {
						profile.PublicKey = pk
					}
				}
			}
		}
	}
	return profile, nil
}

type HugCustomClaims struct {
	Address string `json:"address"`
	jwt.RegisteredClaims
}

func NewAddressJWT(address ethcommon.Address) (string, error) {
	claims := HugCustomClaims{
		Address: address.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(HmacExpirationDuration)),
			Issuer:    "test",
			ID:        address.String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(HmacKey))
	return ss, err
}

func ParseAddressJWT(tokenStr string) (ethcommon.Address, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &HugCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(HmacKey), nil
	})

	if claims, ok := token.Claims.(*HugCustomClaims); ok && token.Valid {
		return ethcommon.HexToAddress(claims.Address), nil
	} else {
		return ethcommon.Address{}, err
	}
}
