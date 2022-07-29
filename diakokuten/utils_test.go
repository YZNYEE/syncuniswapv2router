package diakokuten

import (
	"fmt"
	"testing"
	"testrouter/log"
)

func TestParseLog(t *testing.T) {
	m := LogTransfer{}
	var mnull interface{}
	mnull = m
	out := mnull.(log.Loginterface)
	fmt.Println(out.Tomongo())
}
