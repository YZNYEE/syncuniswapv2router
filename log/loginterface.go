package log

type Logheader struct {
	Txhash string `bson:"txhash"`
	Index  uint64 `bson:"index"`
}

type Loginterface interface {
	Tomongo() interface{}
}
