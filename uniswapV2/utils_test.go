package uniswapv2

import (
	"fmt"
	"testing"
)

func TestParselog(t *testing.T) {
	for k, v := range Contractabi.Events {
		fmt.Println(k)
		fmt.Println(v)
	}
}
