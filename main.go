package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testrouter/syn"
)

var Host = "https://mainnet.infura.io/v3/4065fc77233f4f18886d7236e9042ca6"

var contractaddress = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"

type Token struct {
	symbol  string
	address string
	decimal int
}

func main() {
	syn.Sync()
}

func parsedarguments(inputs string) {
	fmt.Println(len(inputs))
	fmt.Println(inputs[0:8])
	fmt.Println(inputs[8 : 8+63])
	fmt.Println(inputs[8+63 : 8+63+63])
	//for cur := 8; cur+63 < len(inputs); cur += 63 {
	//	fmt.Println(inputs[cur : cur+63])
	//}
}

func getPaths(routes []Token) []common.Address {
	cAddres := []common.Address{}
	for _, token := range routes {
		cAddres = append(cAddres, common.HexToAddress(token.address))
	}
	return cAddres
}
