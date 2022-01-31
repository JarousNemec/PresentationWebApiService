package main

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/storage"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

var tpl *template.Template
var (
	tableCli storage.TableServiceClient
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

func getSolarPanelState() interface{} {
	client, err := storage.NewBasicClient(account, key)
	if err != nil {
		fmt.Printf("%s: \n", err)
	}
	tableCli = client.GetTableService()
	fmt.Println(tableCli)
	table := tableCli.GetTableReference(sptablename)
	entities, err := table.QueryEntities(30, fullmetadata, nil)

	return entities.Entities[0].Properties["state"]
}

func getSolarPanelStateJSON(c *gin.Context) {
	client, err := storage.NewBasicClient(account, key)
	if err != nil {
		fmt.Printf("%s: \n", err)
	}
	tableCli = client.GetTableService()
	fmt.Println(tableCli)
	table := tableCli.GetTableReference(sptablename)
	entities, err := table.QueryEntities(30, fullmetadata, nil)

	c.JSON(http.StatusOK, entities.Entities)
}

func handleRequests() {

	router := gin.Default()
	authData := router.Group("/", AuthMiddleWare())
	{
		authData.GET("/addresult", addGameResult)
	}
	router.GET("/allresults", allGameResults)
	router.GET("/spstate", getSolarPanelStateJSON)

	router.Static("/assets/", "./templates/assets")
	router.Static("/images/", "./templates/images")

	router.Handle("GET", "/solarIndex", solarPanelsApp)

	router.Handle("POST", "/solarIndex", solarPanelsApp)

	router.Handle(http.MethodGet, "/mfLeaderBoard", mfLeaderBoardApp)

	err := router.Run(":8081")
	if err != nil {
		fmt.Println(err)
	}
}

type Content struct {
	Status bool
}

func solarPanelsApp(c *gin.Context) {
	state := getSolarPanelState()
	if c.Request.Method == http.MethodPost && authSettingSolarState(c) {
		insertSpState(!state.(bool))
		state = getSolarPanelState()
	}
	if state == true {
		status := Content{Status: true}
		err := tpl.ExecuteTemplate(c.Writer, "solarIndex.html", status)
		if err != nil {
			return
		}

	} else {
		status := Content{Status: false}
		err := tpl.ExecuteTemplate(c.Writer, "solarIndex.html", status)
		if err != nil {
			return
		}
	}

}

func authSettingSolarState(c *gin.Context) bool {
	if c.PostForm("code") == authorization["solerstate_set_code"] {
		return true
	}
	return false
}

func mfLeaderBoardApp(c *gin.Context) {
	err := tpl.ExecuteTemplate(c.Writer, "mfLeaderBoard.html", nil)
	if err != nil {
		return
	}

}

func main() {
	loadAuthorization()
	tpl = template.Must(template.ParseGlob("./templates/*.html"))
	handleRequests()
}
