package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

var host = "https://mainnet.infura.io/v3/4065fc77233f4f18886d7236e9042ca6"

var ETHClient *ethclient.Client

func Getclient() *ethclient.Client {
	c, err := ethclient.Dial(host)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func Getblockbynum(num int64) (*types.Block, error) {
	b, err := ETHClient.BlockByNumber(context.Background(), big.NewInt(num))
	return b, err
}

func init() {
	ETHClient = Getclient()
}
