package brokie

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testrouter/log"
)

const (
	Approvalhash             = "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"
	AutoLiquifyhash          = "0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163"
	OwnershipTransferredhash = "0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163"
	Transferhash             = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
)

type MongoTransfer struct {
	Txhash string `bson:"txhash"`
	Index  uint64 `bson:"index"`
	From   string `bson:"from"`
	To     string `bson:"to"`
	Value  string `bson:"value"`
}

type LogTransfer struct {
	log.Logheader
	From  common.Address
	To    common.Address
	Value *big.Int
}

func (receiver LogTransfer) Tomongo() interface{} {
	var m MongoTransfer
	m.Txhash = receiver.Txhash
	m.Index = receiver.Index
	m.From = receiver.From.Hex()
	m.To = receiver.To.Hex()
	m.Value = receiver.Value.String()
	return m
}

type LogApproval struct {
	log.Logheader
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
}

type MongoApproval struct {
	Txhash  string `bson:"txhash"`
	Index   uint64 `bson:"index"`
	Owner   string `bson:"owner"`
	Spender string `bson:"spender"`
	Value   string `bson:"value"`
}

func (receiver LogApproval) Tomongo() interface{} {
	//var m MongoApproval
	var m MongoApproval
	m.Txhash = receiver.Logheader.Txhash
	m.Index = receiver.Logheader.Index
	m.Owner = receiver.Owner.Hex()
	m.Spender = receiver.Spender.Hex()
	m.Value = receiver.Value.String()
	return m
}

type LogAutoLiquify struct {
	log.Logheader
	AmountETH *big.Int
	amountBOG *big.Int
}

type MongoAutoLiquify struct {
	Txhash    string `bson:"txhash"`
	Index     uint64 `bson:"index"`
	AmountETH string `bson:"amountETH"`
	AmountBOG string `bson:"amountBOG"`
}

func (l LogAutoLiquify) Tomongo() interface{} {
	var m MongoAutoLiquify
	m.Txhash = l.Txhash
	m.Index = l.Index
	m.AmountBOG = l.amountBOG.String()
	m.AmountETH = l.AmountETH.String()
	return m
}

type LogOwnershipTransferred struct {
	log.Logheader
	Owner common.Address
}

type MongoOwnershipTransferred struct {
	Txhash string `bson:"txhash"`
	Index  uint64 `bson:"index"`
	Owner  string `bson:"owner"`
}

func (l LogOwnershipTransferred) Tomongo() interface{} {
	var m MongoOwnershipTransferred
	m.Txhash = l.Txhash
	m.Index = l.Index
	return m
}
