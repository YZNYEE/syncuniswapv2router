package brokie

import (
	"fmt"
	"testing"
)

func TestGetBrokieabi(t *testing.T) {
	for k, v := range Brokieabi.Events {
		fmt.Println(k)
		fmt.Println(v)
	}
}
