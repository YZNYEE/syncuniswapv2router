package weth

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testrouter/log"
)

const (
	Approvalhash   = "8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"
	Transferhash   = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	Deposithash    = "0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c"
	Withdrawalhash = "0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65"
)

var WETHAddress = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")

type LogApproval struct {
	log.Logheader
	Src common.Address
	Guy common.Address
	Wad *big.Int
}

type LogTransfer struct {
	log.Logheader
	Src common.Address
	Dst common.Address
	Wad *big.Int
}

type Mongotransfer struct {
	Txhash string `bson:"txhash"`
	Index  uint64 `bson:"index"`
	Src    string `bson:"src"`
	Dst    string `bson:"dst"`
	Wad    string `bson:"wad"`
}

func (receiver LogTransfer) Tomongo() interface{} {
	var m Mongotransfer
	m.Txhash = receiver.Txhash
	m.Index = receiver.Index
	m.Src = receiver.Src.Hex()
	m.Dst = receiver.Dst.Hex()
	m.Wad = receiver.Wad.String()
	return m
}

type LogDeposit struct {
	log.Logheader
	Dst common.Address
}

type Mongodeposit struct {
	Txhash string `bson:"txhash"`
	Index  uint64 `bson:"index"`
	Dst    string `bson:"dst"`
}

func (receiver LogDeposit) Tomongo() interface{} {
	var m Mongodeposit
	m.Txhash = receiver.Txhash
	m.Index = receiver.Index
	m.Dst = receiver.Dst.Hex()
	return m
}

type LogWithdrawal struct {
	log.Logheader
	Src common.Address
	Wad *big.Int
}

type Mongowithdrawl struct {
	Txhash string `bson:"txhash"`
	Index  uint64 `bson:"index"`
	Src    string `bson:"src"`
	Wad    string `bson:"wad"`
}

func (receiver LogWithdrawal) Tomongo() interface{} {
	var m Mongowithdrawl
	m.Txhash = receiver.Txhash
	m.Index = receiver.Index
	m.Src = receiver.Src.Hex()
	m.Wad = receiver.Wad.String()
	return m
}
