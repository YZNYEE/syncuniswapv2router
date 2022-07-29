package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"strings"
	"testrouter/entity"
	"testrouter/router"
)

var Parse abi.ABI
var err error

var Tokenmap map[string]string
var Uniswap = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"

func Parsemethod(transaction *types.Transaction, blocknum string) *entity.UniswapV2TX {

	inputs := transaction.Data()
	meth, err := Parse.MethodById(inputs[:4])
	if err != nil {
		log.Fatal("解析方法出错:", err)
	}
	if meth.Name != "swapExactTokensForTokens" {
		fmt.Println("暂不支持其它方法")
		return nil
	}
	result, err := meth.Inputs.UnpackValues(inputs[4:])
	if err != nil {
		log.Fatal(err)
	}
	cur := ParseRes(result, meth.Name, transaction.Hash().Hex(), meth.Inputs)
	cur.BlockNum = blocknum
	return cur

}

func check(ok bool, head string) {
	if ok != false {
		log.Fatal(head, "解析出错:")
	}
}

func ParseRes(res []interface{}, name, hash string, inputs abi.Arguments) *entity.UniswapV2TX {
	amountin := res[0].(*big.Int)
	amountout := res[1].(*big.Int)
	path := res[2].([]common.Address)
	pathstore := make([]string, 0)
	for _, k := range path {
		pathstore = append(pathstore, Tokenmap[k.Hex()])
	}
	to := res[3].(common.Address)
	deadline := res[4].(*big.Int)
	return &entity.UniswapV2TX{Name: name, Amountin: amountin.String(), AmountOutMin: amountout.String(), To: to.Hex(), Deadline: deadline.String(), Path: pathstore, Hash: hash}
}

func Parseblock(block *types.Block) []*entity.UniswapV2TX {
	var out []*entity.UniswapV2TX
	out = make([]*entity.UniswapV2TX, 0)
	txs := block.Transactions()
	fmt.Println("txlen:", len(txs))
	for i := 0; i < len(txs); i++ {
		tx := txs[i]
		//判断是否是uniswapv2router合约
		if tx.To() == nil {
			continue
		}
		//fmt.Println("###########", i, tx.Hash())
		if tx.To().Hex() == Uniswap {
			//fmt.Println(i, tx.Hash().Hex())
			cur := Parsemethod(tx, block.Number().String())
			out = append(out, cur)
		}
	}
	return out
}

func init() {
	Parse, err = abi.JSON(strings.NewReader(router.RouterABI))
	if err != nil {
		log.Fatal("加载abi出错:", err)
	}
	Tokenmap = make(map[string]string, 16)
	Tokenmap["0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"] = "WETH"
	Tokenmap["0x80aB141F324C3d6F2b18b030f1C4E95d4d658778"] = "DEA"
	Tokenmap["0x3b62F3820e0B035cc4aD602dECe6d796BC325325"] = "DEUS"
	Tokenmap["0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"] = "USDC"
	Tokenmap["0xf1f955016EcbCd7321c7266BccFB96c68ea5E49b"] = "RLY"
	Tokenmap["0xdAC17F958D2ee523a2206206994597C13D831ec7"] = "USDT"
	Tokenmap["0x53fD2342B43eCD24AEf1535BC3797F509616Ce8c"] = "ANARCHY"

}
