package weth

import "fmt"
import "testing"

func TestParselog(t *testing.T) {
	//fmt.Println(len(Contractabi.Events))
	for k, v := range Contractabi.Events {
		fmt.Println(k)
		fmt.Println(v)
	}
}
