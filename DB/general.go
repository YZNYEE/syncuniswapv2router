package DB

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	_ "log"
	"math/big"
	_ "math/big"
	"strconv"
	"testrouter/entity"
)

var url string = "mongodb://localhost:27017"
var clientoptions *options.ClientOptions
var client *mongo.Client
var DB *mongo.Database
var err error
var Coll *mongo.Collection
var Colltx *mongo.Collection
var Colllog *mongo.Collection

func InsertUniswapV2(tx entity.UniswapV2TX) {
	Colltx.InsertOne(context.Background(), tx)
}

func InsertLog(log interface{}) {
	Colllog.InsertOne(context.Background(), log)
}

func InsertBlockNum(blocknum int64) {
	var block = entity.Mongounit{"blocknum", big.NewInt(blocknum).String()}
	Colltx.InsertOne(context.Background(), block)
}

func UpdateBlockNum(num int64) {
	Deleteblocknum()
	InsertBlockNum(num)
}

func Deleteblocknum() {
	Colltx.DeleteOne(context.TODO(), bson.D{{"key", "blocknum"}})
}

func GetBlockNum() int64 {
	var out entity.Mongounit
	Colltx.FindOne(context.Background(), bson.D{{"key", "blocknum"}}).Decode(&out)
	out1, err := strconv.ParseInt(out.Value, 10, 64)
	if err != nil {
		return -1
	}
	return out1
}

func init() {
	clientoptions = options.Client().ApplyURI(url)
	client, err = mongo.Connect(context.TODO(), clientoptions)
	if err != nil {
		log.Fatal(err)
	}
	DB = client.Database("uniswap")
	Coll = DB.Collection("V2")
	Colltx = DB.Collection("TX")
	Colllog = DB.Collection("Log")

}
