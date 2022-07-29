package syn

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

const host = "https://mainnet.infura.io/v3/4065fc77233f4f18886d7236e9042ca6"

func Getlatestblocknum() int64 {
	n, er := client.BlockNumber(context.Background())
	if er != nil {
		log.Fatal(er)
	}
	return int64(n)
}
func GetBlockbynum(n int64) *types.Block {
	b, er := client.BlockByNumber(context.Background(), big.NewInt(n))
	if er != nil {
		log.Fatal("获取区块错误，区块号:", n)
	}
	return b
}

var client *ethclient.Client
var err error

func init() {
	client, err = ethclient.Dial(host)
	if err != nil {
		log.Fatal("加载节点出错!!!")
	}
}
