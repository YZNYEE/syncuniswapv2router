package syn

import (
	"context"
	"errors"
	"fmt"
	_ "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"testrouter/DB"
	"testrouter/brokie"
	"testrouter/diakokuten"
	"testrouter/sniper"
	uniswapv2 "testrouter/uniswapV2"
	"testrouter/weth"
)

const uniswapv2routeraddress = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
const uniswappairaddress = "0x4DfD0DD08a16448a7C7E0D0d832958ec215D876f"
const wethaddress = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
const diakukotenaddress = "0x3Be79d67a911e48Ea249f331bf7d3894A418ad56"
const uniswapaddressv2 = "0xBeB909a317e00cbCF737A84AEb64080E772548CD"
const sniperaddress = "0x788F39d599C790a2E7E5CcC6F9bf539063cB339A"

func Syncblock(block *types.Block) {
	for _, tx := range block.Transactions() {
		err := Syntx(tx)
		if err != nil {
			continue
		}
	}
}

func Sync() {
	num := DB.GetBlockNum()
	curnum := Getlatestblocknum()
	if num == -1 {
		num = curnum - 1
	}
	for {
		if num == curnum {
			break
		}
		num++
		fmt.Println("正在同步区块:", num)
		block := GetBlockbynum(num)
		Syncblock(block)
		DB.UpdateBlockNum(num)

	}
}

func Syntx(transaction *types.Transaction) error {
	if transaction.To().Hex() != uniswapv2routeraddress {
		return errors.New("只支持uniswapv2")
	}
	reipts, err := client.TransactionReceipt(context.Background(), transaction.Hash())
	if err != nil {
		return err
	}
	for _, logts := range reipts.Logs {
		switch logts.Address.Hex() {
		case uniswappairaddress:
			out := uniswapv2.Parse(logts)
			DB.InsertLog(out.Tomongo())
		case wethaddress:
			out := weth.Parse(logts)
			DB.InsertLog(out.Tomongo())
		case diakukotenaddress:
			out := diakokuten.Parse(logts)
			DB.InsertLog(out.Tomongo())
		case brokie.Brokieaddress:
			out := brokie.Parse(logts)
			DB.InsertLog(out.Tomongo())
		case sniperaddress:
			out := sniper.Parse(logts)
			DB.InsertLog(out.Tomongo())
		default:
			fmt.Println("交易:", transaction.Hash())
			log.Fatal("暂不支持该合约:", logts.Address.Hex())
		}
	}
	return nil
}
