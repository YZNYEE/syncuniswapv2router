package entity

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testrouter/log"
)

type PairCreated struct {
	log.Logheader
	Token0 common.Address
	Token1 common.Address
	pair   common.Address
	Value  *big.Int
}

type Approval struct {
	log.Logheader
	owner   common.Address
	spender common.Address
	value   *big.Int
}

type Transfer struct {
	log.Logheader
	from  common.Address
	to    common.Address
	value *big.Int
}

type Mint struct {
	log.Logheader
	sender  common.Address
	amount0 *big.Int
	amount1 *big.Int
}

type Burn struct {
	log.Logheader
	sender  common.Address
	amount0 *big.Int
	amount1 *big.Int
	to      common.Address
}
