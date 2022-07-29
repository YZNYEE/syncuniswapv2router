package entity

type UniswapV2TX struct {
	Amountin     string   `bson:"amountin"`
	AmountOutMin string   `bson:"amountOutMin"`
	Path         []string `bson:"path"`
	To           string   `bson:"to"`
	Deadline     string   `bson:"deadline"`
	Name         string   `bson:"name"`
	Hash         string   `bson:"hash"`
	BlockNum     string   `bson:"blockNum"`
}
