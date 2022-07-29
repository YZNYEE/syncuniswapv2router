package sniper

import (
	"fmt"
	"testing"
)

func TestParselog(t *testing.T) {
	for k, v := range Sniperabi.Events {
		fmt.Println(k)
		fmt.Println(v)
	}
}
