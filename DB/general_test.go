package DB

import (
	"fmt"
	"log"
	"testing"
)
import "testrouter/eth"
import "testrouter/utils"

func TestInsertUniswapV2(t *testing.T) {
	b, err := eth.Getblockbynum(15223547)
	if err != nil {
		log.Fatal(err)
	}
	out := utils.Parseblock(b)
	for _, v := range out {
		fmt.Println("1111111", v)
		InsertUniswapV2(*v)
	}
}
