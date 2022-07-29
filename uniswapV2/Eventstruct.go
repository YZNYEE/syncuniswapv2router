package uniswapv2

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testrouter/log"
)

const (
	burnhash     = "0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496"
	minthash     = "0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f"
	swaphash     = "0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"
	synchash     = "0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1"
	transferhash = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	approvalhash = "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"
)

var UniswapAddress = common.HexToAddress("0x4DfD0DD08a16448a7C7E0D0d832958ec215D876f")
var UniswapAddress2 = common.HexToAddress("0xBeB909a317e00cbCF737A84AEb64080E772548CD")

type Logmint struct {
	log.Logheader
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
}

type Mongomint struct {
	Txhash  string `bson:"txhash"`
	Index   uint64 `bson:"index"`
	Spender string `bson:"spender"`
	Amount0 string `bson:"amount0"`
	Amount1 string `bson:"amount1"`
}

func (receiver Logmint) Tomongo() (m Mongomint) {
	m.Txhash = receiver.Txhash
	m.Index = receiver.Index
	m.Amount0 = receiver.Amount0.String()
	m.Amount1 = receiver.Amount1.String()
	return m
}

type Logswap struct {
	log.Logheader
	Sender     common.Address
	To         common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
}

type Mongoswap struct {
	Txhash     string `bson:"txhash"`
	Index      uint64 `bson:"index"`
	Amount0In  string `bson:"amount0In"`
	Amount1In  string `bson:"amount1In"`
	Amount0Out string `bson:"amount0Out"`
	Amount1Out string `bson:"amount1Out"`
}

func (receiver Logswap) Tomongo() interface{} {
	var m Mongoswap
	m.Txhash = receiver.Txhash
	m.Index = receiver.Index
	m.Amount0In = receiver.Amount0In.String()
	m.Amount1In = receiver.Amount1In.String()
	m.Amount0Out = receiver.Amount0Out.String()
	m.Amount1Out = receiver.Amount1Out.String()
	return m
}

type Logsync struct {
	log.Logheader
	Reserve0 *big.Int
	Reserve1 *big.Int
}

type Mongosync struct {
	Txhash   string `bson:"txhash"`
	Index    uint64 `bson:"index"`
	Reserve0 string
	Reserve1 string
}

func (receiver Logsync) Tomongo() interface{} {
	var m Mongosync
	m.Txhash = receiver.Txhash
	m.Index = receiver.Index
	m.Reserve0 = receiver.Reserve0.String()
	m.Reserve1 = receiver.Reserve1.String()
	return m
}

type Logtransfer struct {
	log.Logheader
	From  common.Address
	To    common.Address
	Value *big.Int
}

type Mongotransfer struct {
	Txhash string `bson:"txhash"`
	Index  uint64 `bson:"index"`
	From   string `bson:"from"`
	To     string `bson:"to"`
	Value  string `bson:"value"`
}

func (receiver Logtransfer) Tomongo() interface{} {
	var m Mongotransfer
	m.Txhash = receiver.Txhash
	m.Index = receiver.Index
	m.From = receiver.From.Hex()
	m.To = receiver.To.Hex()
	return m
}

type MongoApproval struct {
	Txhash  string `bson:"txhash"`
	Index   uint64 `bson:"index"`
	Owner   string `bson:"owner"`
	Spender string `bson:"spender"`
	Value   string `bson:"value"`
}
type Logapproval struct {
	log.Logheader
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
}

//
////
func (receiver Logapproval) Tomongo(interface{}) {
	var m MongoApproval
	m.Txhash = receiver.Txhash
	m.Index = receiver.Index
	m.Spender = receiver.Spender.Hex()
	m.Owner = receiver.Owner.Hex()
	m.Value = receiver.Value.String()
}

type Logburn struct {
	log.Logheader
	Sender  common.Address
	To      common.Address
	Amount0 *big.Int
	Amount1 *big.Int
}

type Mongoburn struct {
	Txhash  string `bson:"txhash"`
	Index   uint64 `bson:"index"`
	Sender  string `bson:"sender"`
	To      string `bson:"to"`
	Amount0 string `bson:"amount0"`
	Amount1 string `bson:"amount1"`
}

func (receiver Logburn) Tomongo() interface{} {
	var m Mongoburn
	m.Txhash = receiver.Txhash
	m.Index = receiver.Index
	m.Sender = receiver.Sender.Hex()
	m.To = receiver.To.Hex()
	m.Amount0 = receiver.Amount0.String()
	m.Amount1 = receiver.Amount1.String()
	return m
}
