package weth

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

func Getabi() abi.ABI {
	contractabi, err := abi.JSON(strings.NewReader(string(WethMetaData.ABI)))
	if err != nil {
		log.Fatal(err)
	}
	return contractabi
}

var Eventmap map[string]interface{}
var Contractabi abi.ABI

func Parselog(log2 *types.Log) (interface{}, int, error) {
	if log2.Address != WETHAddress {
		return nil, LogErr, errors.New("地址不正确")
	}
	switch log2.Topics[0].Hex() {
	case Approvalhash:
		var log LogApproval
		log.Logheader.Txhash = log2.TxHash.Hex()
		log.Logheader.Index = uint64(log2.Index)
		Contractabi.UnpackIntoInterface(&log, "Approval", log2.Data)
		log.Src = common.HexToAddress(log2.Topics[1].Hex())
		log.Guy = common.HexToAddress(log2.Topics[2].Hex())
		return log, ApprovalLog, nil
	case Transferhash:
		var log LogTransfer
		log.Logheader.Txhash = log2.TxHash.Hex()
		log.Logheader.Index = uint64(log2.Index)
		Contractabi.UnpackIntoInterface(&log, "Transfer", log2.Data)
		log.Src = common.HexToAddress(log2.Topics[1].Hex())
		log.Dst = common.HexToAddress(log2.Topics[2].Hex())
		return log, TransferLog, nil
	case Deposithash:
		var log LogDeposit
		log.Logheader.Txhash = log2.TxHash.Hex()
		log.Logheader.Index = uint64(log2.Index)
		Contractabi.UnpackIntoInterface(&log, "Deposit", log2.Data)
		log.Dst = common.HexToAddress(log2.Topics[1].Hex())
		return log, DepositLog, nil
	case Withdrawalhash:
		var log LogWithdrawal
		log.Logheader.Txhash = log2.TxHash.Hex()
		log.Logheader.Index = uint64(log2.Index)
		Contractabi.UnpackIntoInterface(&log, "Withdrawal", log2.Data)
		log.Src = common.HexToAddress(log2.Topics[1].Hex())
		return log, WithdrawalLog, nil
	default:
		return nil, LogErr, errors.New("topic不准确")
	}
}
func Parse(log2 *types.Log) routerlog.Loginterface {
	lg, kind, err := Parselog(log2)
	_ = kind
	if err != nil {
		log.Fatal("解析出错")
	}
	return lg.(routerlog.Loginterface)
}

const (
	LogErr = iota
	ApprovalLog
	TransferLog
	DepositLog
	WithdrawalLog
)

func init() {
	Contractabi = Getabi()
	event := Contractabi.Events
	Eventmap = make(map[string]interface{}, 0)
	for k, v := range event {
		hash := hex.EncodeToString(crypto.Keccak256([]byte(v.Sig)))
		Eventmap[hash] = k
	}
}
