package brokie

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"strings"
	routerlog "testrouter/log"
)

var Brokieabi abi.ABI
var Brokieaddress = "0xd4032901D1B8d9CAcab6CaF62184EfDFe4A8c313"

func GetBrokieabi() abi.ABI {
	contractabi, err := abi.JSON(strings.NewReader(string(BrokieMetaData.ABI)))
	if err != nil {
		log.Fatal(err)
	}
	return contractabi
}

const (
	logerr = iota
	logapproval
	logautoLiquify
	logownershipTransferred
	logtransfer
)

func GetEventmap() map[string]interface{} {
	contractabi := GetBrokieabi()
	event := contractabi.Events
	eventmap := make(map[string]interface{}, 0)
	for k, v := range event {
		hash := hex.EncodeToString(crypto.Keccak256([]byte(v.Sig)))
		Eventmap[hash] = k
	}
	return eventmap
}

func ParseLog(log2 *types.Log) (interface{}, int, error) {
	if common.HexToAddress(Brokieaddress) != log2.Address {
		return nil, logerr, errors.New("不是brokie合约事件")
	}
	if log2.Topics[0].Hex() == Approvalhash {
		var l LogApproval
		l.Logheader.Txhash = log2.TxHash.Hex()
		l.Index = uint64(log2.Index)
		err := Brokieabi.UnpackIntoInterface(&l, "Approval", log2.Data)
		l.Owner = common.HexToAddress(log2.Topics[1].Hex())
		l.Spender = common.HexToAddress(log2.Topics[2].Hex())
		return l, logapproval, err
	}
	if log2.Topics[0].Hex() == Transferhash {
		var l LogTransfer
		l.Logheader.Txhash = log2.TxHash.Hex()
		l.Index = uint64(log2.Index)
		err := Brokieabi.UnpackIntoInterface(&l, "Transfer", log2.Data)
		l.From = common.HexToAddress(log2.Topics[1].Hex())
		l.To = common.HexToAddress(log2.Topics[2].Hex())
		return l, logtransfer, err
	}
	if log2.Topics[0].Hex() == AutoLiquifyhash {
		var l LogApproval
		l.Logheader.Txhash = log2.TxHash.Hex()
		l.Index = uint64(log2.Index)
		err := Brokieabi.UnpackIntoInterface(&l, "AutoLiquify", log2.Data)
		return l, logapproval, err
	}
	if log2.Topics[0].Hex() == OwnershipTransferredhash {
		var l LogApproval
		l.Logheader.Txhash = log2.TxHash.Hex()
		l.Index = uint64(log2.Index)
		err := Brokieabi.UnpackIntoInterface(&l, "OwnershipTransferred", log2.Data)
		return l, logapproval, err
	}
	return nil, logerr, nil
}
func Parse(log2 *types.Log) routerlog.Loginterface {
	lg, kind, err := ParseLog(log2)
	_ = kind
	if err != nil {
		log.Fatal("解析出错")
	}
	fmt.Println(kind)
	fmt.Println(lg)
	return lg.(routerlog.Loginterface)
}

var Eventmap map[string]interface{}

func init() {
	Brokieabi = GetBrokieabi()
	event := Brokieabi.Events
	Eventmap = make(map[string]interface{}, 0)
	for k, v := range event {
		hash := hex.EncodeToString(crypto.Keccak256([]byte(v.Sig)))
		Eventmap[hash] = k
	}
}
