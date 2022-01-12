package main

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/storage"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	tableCli storage.TableServiceClient
)

const (
	account      = "minefinderstorageaccount"
	key          = "vqgU7HTQ39+i05unpYXB2/sjw6YVBVCzcgbDlcq1UH04qw2V8TxKMOoaobMkjUuM587C/0NZjbtdBobd83rhGg=="
	fullmetadata = "application/json;odata=fullmetadata"
	tablename    = "MineFinder"
)

func insert(player string, playtime string, fieldSize string, mineCount string) {
	client, err := storage.NewBasicClient(account, key)
	if err != nil {
		fmt.Printf("%s: \n", err)
	}
	tableCli = client.GetTableService()
	fmt.Println(tableCli)
	table := tableCli.GetTableReference(tablename)
	entity := table.GetEntityReference("1", time.Now().String())
	props := map[string]interface{}{
		"player":    player,
		"playTime":  playtime,
		"fieldSize": fieldSize,
		"mineCount": mineCount,
	}
	entity.Properties = props
	err = entity.Insert(fullmetadata, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Inserted! ")
	}
}

func allGameResults(c *gin.Context) {
	client, err := storage.NewBasicClient(account, key)
	if err != nil {
		fmt.Printf("%s: \n", err)
	}
	tableCli = client.GetTableService()
	fmt.Println(tableCli)
	table := tableCli.GetTableReference(tablename)
	entities, err := table.QueryEntities(30, fullmetadata, nil)
	var messages []byte

	messages, err = json.Marshal(entities.Entities)
	var mess string
	err = json.Unmarshal(messages, &mess)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, entities.Entities)

}

//type T struct {
//	PartitionKey string `json:"PartitionKey"`
//	RowKey       string `json:"RowKey"`
//	FieldSize    string `json:"fieldSize"`
//	MineCount    string `json:"mineCount"`
//	PlayTime     string `json:"playTime"`
//	Player       string `json:"player"`
//}
//
//func data(c *gin.Context) {
//	var datas []T
//	obj := T{PartitionKey: "buhehe", RowKey: "key", FieldSize: "5", MineCount: "5"}
//	datas = append(datas, obj)
//
//	c.JSON(http.StatusOK, datas)
//}

func addResult(c *gin.Context) {
	player := c.Request.URL.Query()["player"][0]
	playtime := c.Request.URL.Query()["playtime"][0]
	fieldSize := c.Request.URL.Query()["fieldSize"][0]
	mineCount := c.Request.URL.Query()["mineCount"][0]
	fmt.Println("inserting")
	insert(player, playtime, fieldSize, mineCount)
}

func handleRequests() {
	r := gin.Default()
	//r.GET("/data.json", data)
	r.GET("/addresult", addResult)
	r.GET("/allresults", allGameResults)
	fmt.Println("prepare inserting")
	r.Use(static.Serve("/", static.LocalFile("./frontend", false)))
	err := r.Run(":8081")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	handleRequests()

}
