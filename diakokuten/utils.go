package diakokuten

import (
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"strings"
	routerlog "testrouter/log"
)

var Diakokutenabi abi.ABI
var Diakokutenaddress = "0x3Be79d67a911e48Ea249f331bf7d3894A418ad56"

func GetDiakokutenabi() abi.ABI {
	contractabi, err := abi.JSON(strings.NewReader(string(DiakokutenABI)))
	if err != nil {
		log.Fatal(err)
	}
	return contractabi
}

const (
	Logerr = iota
	Logtransfer
	Logapproval
)

func ParseLog(log2 *types.Log) (interface{}, int, error) {
	if common.HexToAddress(Diakokutenaddress) != log2.Address {
		return nil, Logerr, errors.New("不是diakokuten合约事件")
	}
	if log2.Topics[0].Hex() == Transferhash {
		var logtransfer LogTransfer
		logtransfer.Logheader.Txhash = log2.TxHash.Hex()
		logtransfer.Logheader.Index = uint64(log2.Index)
		err := Diakokutenabi.UnpackIntoInterface(&logtransfer, "Transfer", log2.Data)
		logtransfer.From = common.HexToAddress(log2.Topics[1].Hex())
		logtransfer.To = common.HexToAddress(log2.Topics[2].Hex())
		return logtransfer, Logtransfer, err
	}
	if log2.Topics[0].Hex() == Approvalhash {
		var log LogApproval
		err := Diakokutenabi.UnpackIntoInterface(&log, "Approval", log2.Data)
		log.Owner = common.HexToAddress(log2.Topics[1].Hex())
		log.Spender = common.HexToAddress(log2.Topics[2].Hex())
		return log, Logapproval, err
	}
	return nil, Logerr, nil
}
func Parse(log2 *types.Log) routerlog.Loginterface {
	lg, kind, err := ParseLog(log2)
	_ = kind
	if err != nil {
		log.Fatal("解析出错")
	}
	return lg.(routerlog.Loginterface)
}

var Eventmap map[string]interface{}

func GetEventmap() map[string]interface{} {
	contractabi := GetDiakokutenabi()
	event := contractabi.Events
	eventmap := make(map[string]interface{}, 0)
	for k, v := range event {
		hash := hex.EncodeToString(crypto.Keccak256([]byte(v.Sig)))
		Eventmap[hash] = k
	}
	return eventmap
}

func init() {
	contractabi := GetDiakokutenabi()
	event := contractabi.Events
	Eventmap = make(map[string]interface{}, 0)
	for k, v := range event {
		hash := hex.EncodeToString(crypto.Keccak256([]byte(v.Sig)))
		Eventmap[hash] = k
	}
	Diakokutenabi = GetDiakokutenabi()
}
