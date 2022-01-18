package main

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/storage"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var (
	tableCli storage.TableServiceClient
)

const (
	account      = "minefinderstorageaccount"
	key          = "vqgU7HTQ39+i05unpYXB2/sjw6YVBVCzcgbDlcq1UH04qw2V8TxKMOoaobMkjUuM587C/0NZjbtdBobd83rhGg=="
	fullmetadata = "application/json;odata=fullmetadata"
	mftablename  = "MineFinder"
	sptablename  = "solarpanels"
)

func insertGameResult(player string, playtime string, fieldSize string, mineCount string) {
	client, err := storage.NewBasicClient(account, key)
	if err != nil {
		fmt.Printf("%s: \n", err)
	}
	tableCli = client.GetTableService()
	fmt.Println(tableCli)
	table := tableCli.GetTableReference(mftablename)
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
	table := tableCli.GetTableReference(mftablename)
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

func addGameResult(c *gin.Context) {
	player := c.Request.URL.Query()["player"][0]
	playtime := c.Request.URL.Query()["playtime"][0]
	fieldSize := c.Request.URL.Query()["fieldSize"][0]
	mineCount := c.Request.URL.Query()["mineCount"][0]
	fmt.Println("inserting")
	insertGameResult(player, playtime, fieldSize, mineCount)
}

func insertSpState(state bool) {
	client, err := storage.NewBasicClient(account, key)
	if err != nil {
		fmt.Printf("%s: \n", err)
	}
	tableCli = client.GetTableService()
	fmt.Println(tableCli)
	table := tableCli.GetTableReference(sptablename)
	entity := table.GetEntityReference("1", "1")
	props := map[string]interface{}{
		"state": state,
	}
	entity.Properties = props
	err = entity.Update(true, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Updated! ")
	}
}

func getSolarPanelState(c *gin.Context) {
	client, err := storage.NewBasicClient(account, key)
	if err != nil {
		fmt.Printf("%s: \n", err)
	}
	tableCli = client.GetTableService()
	fmt.Println(tableCli)
	table := tableCli.GetTableReference(sptablename)
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

func setSolarPanelsState(c *gin.Context) {
	//TODO: nastavení stavu panelů; vymyslet způsob na bezpečný call, nemohl zavolat nikdo jiný
	fmt.Println("setting spstate...")
	data := c.Request.URL.Query()["state"][0]
	strings.ToLower(data)
	state := true
	if data == "true" {
		state = true
	}
	if data == "false" {
		state = false
	}
	insertSpState(state)
}

func handleRequests() {
	r := gin.Default()
	r.GET("/addresult", addGameResult)
	r.GET("/allresults", allGameResults)
	r.GET("/setspstate", setSolarPanelsState)
	r.GET("/spstate", getSolarPanelState)
	r.Use(static.Serve("/", static.LocalFile("./frontend", false)))
	err := r.Run(":8081")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	handleRequests()
}
