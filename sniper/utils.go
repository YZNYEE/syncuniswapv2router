package sniper

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

var Sniperabi abi.ABI
var Sniperaddress = "0x788F39d599C790a2E7E5CcC6F9bf539063cB339A"

func GetBrokieabi() abi.ABI {
	contractabi, err := abi.JSON(strings.NewReader(string(SniperMetaData.ABI)))
	if err != nil {
		log.Fatal(err)
	}
	return contractabi
}

const (
	logerr = iota
	logsuccess
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
	if common.HexToAddress(Sniperaddress) != log2.Address {
		return nil, logerr, errors.New("不是sniper合约事件")
	}
	if log2.Topics[0].Hex() == Approvalhash {
		var l LogApproval
		l.Logheader.Txhash = log2.TxHash.Hex()
		l.Index = uint64(log2.Index)
		err := Sniperabi.UnpackIntoInterface(&l, "Approval", log2.Data)
		l.Owner = common.HexToAddress(log2.Topics[1].Hex())
		l.Spender = common.HexToAddress(log2.Topics[2].Hex())
		return l, logsuccess, err
	}
	if log2.Topics[0].Hex() == Transferhash {
		var l LogTransfer
		l.Logheader.Txhash = log2.TxHash.Hex()
		l.Index = uint64(log2.Index)
		err := Sniperabi.UnpackIntoInterface(&l, "Transfer", log2.Data)
		l.From = common.HexToAddress(log2.Topics[1].Hex())
		l.To = common.HexToAddress(log2.Topics[2].Hex())
		return l, logsuccess, err
	}
	if log2.Topics[0].Hex() == SetAutomatedMarketMakerPairhash {
		var l LogSetAutomatedMarketMakerPair
		l.Logheader.Txhash = log2.TxHash.Hex()
		l.Index = uint64(log2.Index)
		err := Sniperabi.UnpackIntoInterface(&l, "SetAutomatedMarketMakerPair", log2.Data)
		l.Pair = common.HexToAddress(log2.Topics[1].Hex())
		l.Value = log2.Topics[2].Hex()
		return l, logsuccess, err
	}
	if log2.Topics[0].Hex() == SwapAndLiquifyhash {
		var l LogSwapAndLiquify
		l.Logheader.Txhash = log2.TxHash.Hex()
		l.Index = uint64(log2.Index)
		err := Sniperabi.UnpackIntoInterface(&l, "SwapAndLiquify", log2.Data)
		return l, logsuccess, err
	}
	if log2.Topics[0].Hex() == ExcludeFromFeeshash {
		var l LogExcludeFromFees
		l.Logheader.Txhash = log2.TxHash.Hex()
		l.Index = uint64(log2.Index)
		err := Sniperabi.UnpackIntoInterface(&l, "ExcludeFromFees", log2.Data)
		l.Account = common.HexToAddress(log2.Topics[1].Hex())
		return l, logsuccess, err
	}
	if log2.Topics[0].Hex() == ExcludeMultipleAccountsFromFeeshash {
		var l LogExcludeFromFees
		l.Logheader.Txhash = log2.TxHash.Hex()
		l.Index = uint64(log2.Index)
		err := Sniperabi.UnpackIntoInterface(&l, "ExcludeMultipleAccountsFromFees", log2.Data)
		l.Account = common.HexToAddress(log2.Topics[1].Hex())
		return l, logsuccess, err
	}
	if log2.Topics[0].Hex() == OwnershipTransferredhash {
		var l LogExcludeFromFees
		l.Logheader.Txhash = log2.TxHash.Hex()
		l.Index = uint64(log2.Index)
		err := Sniperabi.UnpackIntoInterface(&l, "ExcludeMultipleAccountsFromFees", log2.Data)
		l.Account = common.HexToAddress(log2.Topics[1].Hex())
		return l, logsuccess, err
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
	Sniperabi = GetBrokieabi()
	event := Sniperabi.Events
	Eventmap = make(map[string]interface{}, 0)
	for k, v := range event {
		hash := hex.EncodeToString(crypto.Keccak256([]byte(v.Sig)))
		Eventmap[hash] = k
	}
}
