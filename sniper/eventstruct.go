package sniper

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testrouter/log"
)

const (
	Approvalhash                        = "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"
	ExcludeFromFeeshash                 = "0x9d8f7706ea1113d1a167b526eca956215946dd36cc7df39eb16180222d8b5df7"
	ExcludeMultipleAccountsFromFeeshash = "0x7fdaf542373fa84f4ee8d662c642f44e4c2276a217d7d29e548b6eb29a233b35"
	OwnershipTransferredhash            = "0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0"
	SetAutomatedMarketMakerPairhash     = "0xffa9187bf1f18bf477bd0ea1bcbb64e93b6a98132473929edfce215cd9b16fab"
	SwapAndLiquifyhash                  = "0x17bbfb9a6069321b6ded73bd96327c9e6b7212a5cd51ff219cd61370acafb561"
	Transferhash                        = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
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

type LogExcludeFromFees struct {
	log.Logheader
	Account    common.Address
	IsExcluded bool
}

type MongoExcludeFromFees struct {
	Txhash     string `bson:"txhash"`
	Index      uint64 `bson:"index"`
	Account    string `bson:"account"`
	IsExcluded bool   `bson:"isExcluded"`
}

func (f LogExcludeFromFees) Tomongo() interface{} {
	var m MongoExcludeFromFees
	m.Account = f.Account.Hex()
	m.IsExcluded = f.IsExcluded
	return m
}

type LogMultipleAccountsFromFees struct {
	log.Logheader
	Accounts   []common.Address
	IsExcluded bool
}

type MongoMultipleAccountsFromFees struct {
	Txhash     string   `bson:"txhash"`
	Index      uint64   `bson:"index"`
	Account    []string `bson:"account"`
	IsExcluded bool     `bson:"isExcluded"`
}

func (f LogMultipleAccountsFromFees) Tomongo() interface{} {
	var m MongoMultipleAccountsFromFees
	am := make([]string, 0)
	for _, a := range f.Accounts {
		am = append(am, a.Hex())
	}
	m.Account = am
	m.IsExcluded = f.IsExcluded
	return m
}

type LogOwnershipTransferred struct {
	log.Logheader
	PreviousOwner common.Address
	NewOwner      common.Address
}

type MongoOwnershipTransferred struct {
	Txhash        string `bson:"txhash"`
	Index         uint64 `bson:"index"`
	PreviousOwner string `bson:"owner"`
	NewOwner      string `bson:"newOwner"`
}

func (t LogOwnershipTransferred) Tomongo() interface{} {
	var m MongoOwnershipTransferred
	m.NewOwner = t.NewOwner.Hex()
	m.PreviousOwner = t.PreviousOwner.Hex()
	return m
}

type LogSetAutomatedMarketMakerPair struct {
	log.Logheader
	Pair  common.Address
	Value string
}

type MongoSetAutomatedMarketMakerPair struct {
	Txhash string
	Index  uint64
	Pair   string
	Value  string
}

func (p LogSetAutomatedMarketMakerPair) Tomongo() interface{} {
	var m MongoSetAutomatedMarketMakerPair
	m.Txhash = p.Txhash
	m.Index = p.Index
	m.Pair = p.Pair.Hex()
	m.Value = p.Value
	return m
}

type LogSwapAndLiquify struct {
	log.Logheader
	TokensSwapped      *big.Int
	EthReceived        *big.Int
	TokensIntoLiqudity *big.Int
}
type MongoSwapAndLiquify struct {
	Txhash             string
	Index              uint64
	TokensSwapped      string
	EthReceived        string
	TokensIntoLiqudity string
}

func (l LogSwapAndLiquify) Tomongo() interface{} {
	var m MongoSwapAndLiquify
	m.Txhash = l.Txhash
	m.Index = l.Index
	m.EthReceived = l.EthReceived.String()
	m.TokensIntoLiqudity = l.TokensIntoLiqudity.String()
	m.TokensSwapped = l.TokensSwapped.String()
	return m
}
