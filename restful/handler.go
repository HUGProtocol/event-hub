package restful

import (
	"encoding/json"
	"errors"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"event-hub/crypto"
	"event-hub/log"
	"net/http"
)

const (
	DefaultRespStatus    = 100
	RespError            = 500
	JWTError             = 401
	WalletSignatureError = 402
	NewAddrJwtError      = 136
)

const (
	BoundMsg = "Bound this address to my HUG account"
	LoginMsg = "Login HUG as current account"
)

var (
	Web3AuthJwtErr     = errors.New("parse web3auth jwt token error")
	WriteResponseErr   = errors.New("write response error")
	WalletSigVerifyErr = errors.New("signature verify failed")
)

type Resp struct {
	Status int    `json:"status"`
	Value  string `json:"value"`
}

func NewResp() Resp {
	return Resp{
		Status: DefaultRespStatus,
		Value:  "",
	}
}

func AutoResponse(writer http.ResponseWriter, resp Resp) {
	b, _ := json.Marshal(resp)
	_, err := writer.Write(b)
	if err != nil {
		log.Error(WriteResponseErr, err)
		return
	}
}

func (s *Service) RegisterWeb3AuthHandler(writer http.ResponseWriter, request *http.Request) {
	resp := NewResp()
	defer AutoResponse(writer, resp)
	jwtStr := request.FormValue("jwt")
	profile, err := crypto.ParseWeb3AuthJose(jwtStr)
	if err != nil {
		resp.Status = JWTError
		resp.Value = err.Error()
		return
	}
	//todo:db

}

func (s *Service) BoundAssetAddress(writer http.ResponseWriter, request *http.Request) {
	resp := NewResp()
	defer AutoResponse(writer, resp)
	//check jwt token
	jwtStr := request.FormValue("jwt")
	profile, err := crypto.ParseWeb3AuthJose(jwtStr)
	if err != nil {
		resp.Status = JWTError
		resp.Value = err.Error()
		return
	}

	sigHex := request.FormValue("sig")
	addressStr := request.FormValue("address")
	address := ethcommon.HexToAddress(addressStr)
	if ok := crypto.VerifySig(BoundMsg, sigHex, address); !ok {
		resp.Status = WalletSignatureError
		resp.Value = WalletSigVerifyErr.Error()
		return
	}
	//todo:db

}

type AddrJwtResp struct {
	Address string `json:"address"`
	JWT     string `json:"jwt"`
}

//wallet signature login
func (s *Service) AssetAddressLogin(writer http.ResponseWriter, request *http.Request) {
	resp := NewResp()
	defer AutoResponse(writer, resp)
	addressStr := request.FormValue("address")
	sigHex := request.FormValue("sig")
	address := ethcommon.HexToAddress(addressStr)
	if ok := crypto.VerifySig(LoginMsg, sigHex, address); ok {
		addrJWT, err := crypto.NewAddressJWT(address)
		if err != nil {
			resp.Status = NewAddrJwtError
			resp.Value = err.Error()
			return
		}
		respJwt := AddrJwtResp{
			Address: addressStr,
			JWT:     addrJWT,
		}
		data, _ := json.Marshal(&respJwt)
		resp.Value = string(data)
		return
	}
}

func (s *Service) GetEvents(writer http.ResponseWriter, request *http.Request) {

}

func (s *Service) GetReviews(writer http.ResponseWriter, request *http.Request) {

}

func (s *Service) PutReview(writer http.ResponseWriter, request *http.Request) {

}
