package mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/ssoyyoung.p/clothForYD/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const dbName = "LooksLike"
const colNameCloth = "clothList"

func getDBinfo() model.DBinfo {
	fileName := "mongoDB/info.json"

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	var dbInfo model.DBinfo
	resByte, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(resByte, &dbInfo)
	if err != nil {
		panic(err)
	}

	return dbInfo
}

func connectDB() (client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)

	dbInfo := getDBinfo()

	clientOpt := options.Client().ApplyURI("mongodb://" + dbInfo.Hostname + dbInfo.Port).SetAuth(options.Credential{
		Username: dbInfo.Username,
		Password: dbInfo.Password,
	})

	client, err := mongo.Connect(ctx, clientOpt)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())

	return client, ctx, cancle
}

// GetClothList func
func GetClothList() []bson.M {
	//var clothes model.Cloth
	var clothes []bson.M

	client, ctx, cancel := connectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	findOptions := options.Find()
	findOptions.SetSort(bson.M{})

	res, err := client.Database(dbName).Collection(colNameCloth).Find(ctx, bson.M{}, findOptions)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	if err = res.All(ctx, &clothes); err != nil {
		panic(err)
	}

	return clothes

}
