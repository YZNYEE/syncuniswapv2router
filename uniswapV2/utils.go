package uniswapv2

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
	contractabi, err := abi.JSON(strings.NewReader(string(Uniswapv2pairABI)))
	if err != nil {
		log.Fatal(err)
	}
	return contractabi
}

var Eventmap map[string]interface{}
var Contractabi abi.ABI

//func Parselog(log2 *types.Log) (interface{}, int, error) {
//	if log2.Address != WETHAddress {
//		return nil, LogErr, errors.New("地址不正确")
//	}
//	switch log2.Topics[0].Hex() {
//	case Approvalhash:
//		var log LogApproval
//		Contractabi.UnpackIntoInterface(&log, "Approval", log2.Data)
//		log.Src = common.HexToAddress(log2.Topics[1].Hex())
//		log.Guy = common.HexToAddress(log2.Topics[2].Hex())
//		return log, ApprovalLog, nil
//	case Transferhash:
//		var log LogTransfer
//		Contractabi.UnpackIntoInterface(&log, "Transfer", log2.Data)
//		log.Src = common.HexToAddress(log2.Topics[1].Hex())
//		log.Dst = common.HexToAddress(log2.Topics[2].Hex())
//		return log, TransferLog, nil
//	case Deposithash:
//		var log LogDeposit
//		Contractabi.UnpackIntoInterface(&log, "Deposit", log2.Data)
//		log.Dst = common.HexToAddress(log2.Topics[1].Hex())
//		return log, DepositLog, nil
//	case Withdrawalhash:
//		var log LogWithdrawal
//		Contractabi.UnpackIntoInterface(&log, "Withdrawal", log2.Data)
//		log.Src = common.HexToAddress(log2.Topics[1].Hex())
//		return log, WithdrawalLog, nil
//	default:
//		return nil, LogErr, errors.New("topic不准确")
//	}
//}

const (
	Logerr = iota
	LogBurn
	LogMint
	LogSwap
	LogSync
	LogTransfer
	LogApproval
)

func Parselog(log2 *types.Log) (interface{}, int, error) {
	if log2.Address != UniswapAddress {
		return nil, 0, errors.New("合约地址不正确")
	}
	switch log2.Topics[0].Hex() {
	case burnhash:
		var logb Logburn
		Contractabi.UnpackIntoInterface(&logb, "Burn", log2.Data)
		//logb.Logheader.Txhash = log2.TxHash.Hex()
		//logb.Logheader.Index = log2.Index
		logb.Sender = common.HexToAddress(log2.Topics[1].Hex())
		logb.To = common.HexToAddress(log2.Topics[2].Hex())
		return logb, LogBurn, nil
	case approvalhash:
		var log Logapproval
		Contractabi.UnpackIntoInterface(&log, "Approval", log2.Data)
		log.Owner = common.HexToAddress(log2.Topics[1].Hex())
		log.Spender = common.HexToAddress(log2.Topics[2].Hex())
		return log, LogApproval, nil
	case transferhash:
		var log Logtransfer
		Contractabi.UnpackIntoInterface(&log, "Trandfer", log2.Data)
		log.Logheader.Txhash = log2.TxHash.Hex()
		log.Logheader.Index = uint64(log2.Index)
		log.From = common.HexToAddress(log2.Topics[1].Hex())
		log.To = common.HexToAddress(log2.Topics[2].Hex())
		return log, LogTransfer, nil
	case synchash:
		var logs Logsync
		logs.Logheader.Txhash = log2.TxHash.Hex()
		logs.Logheader.Index = uint64(log2.Index)
		Contractabi.UnpackIntoInterface(&logs, "Sync", log2.Data)
		return logs, LogSync, nil
	case swaphash:
		var logsw Logswap
		logsw.Logheader.Txhash = log2.TxHash.Hex()
		logsw.Logheader.Index = uint64(log2.Index)
		err := Contractabi.UnpackIntoInterface(&logsw, "Swap", log2.Data)
		if err != nil {
			log.Fatal(err)
		}
		logsw.Sender = common.HexToAddress(log2.Topics[1].Hex())
		logsw.To = common.HexToAddress(log2.Topics[2].Hex())
		return logsw, LogSwap, nil
	default:
		return nil, Logerr, errors.New("日志解析错误")
	}
	return nil, Logerr, errors.New("日志解析错误")
}

func Parse(log2 *types.Log) routerlog.Loginterface {
	lg, kind, err := Parselog(log2)
	_ = kind
	if err != nil {

		log.Fatal("解析出错,log:")
	}
	return lg.(routerlog.Loginterface)
}

func init() {
	Contractabi = Getabi()
	event := Contractabi.Events
	Eventmap = make(map[string]interface{}, 0)
	for k, v := range event {
		hash := hex.EncodeToString(crypto.Keccak256([]byte(v.Sig)))
		Eventmap[hash] = k
	}
}
