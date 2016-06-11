package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/mgerb42/goskeleton/controller/api"
	"github.com/mgerb42/goskeleton/db"
	"github.com/mgerb42/goskeleton/route"
)

//structure for application configurations
type Config struct {
	Database db.DatabaseInfo `json:"database"`
	Api      api.ApiInfo     `json:"api"`
	Server   ServerConfig    `json:"server"`
}

//configuration settings for the web
type ServerConfig struct {
	Port string `json:"port"`
}

func main() {

	//read config file and store configurations
	configurations := readConfig()

	//connect to database if configured
	if configurations.Database.Enabled == true {
		db.Configure(configurations.Database)
		db.Mongo.Connect()
	}

	//store password for api key
	api.Configure(configurations.Api)

	//start up the web server
	log.Println("Starting Server...")
	log.Println(http.ListenAndServe(":"+configurations.Server.Port, route.Routes()))
}

//read the config file and return JsonObject struct
func readConfig() Config {

	log.Println("Reading config file...")

	file, e := ioutil.ReadFile("./config.json")

	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	fmt.Printf("%s\n", string(file))

	//m := new(Dispatch)
	//var m interface{}
	var result Config

	//conver to struct
	err := json.Unmarshal(file, &result)

	if err != nil {
		fmt.Println(err)
	}

	return result
}
