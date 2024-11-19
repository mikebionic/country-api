package main

import (
	"countryApi/config"
	"countryApi/database"
	app "countryApi/internal"
	"fmt"
)

func main() {
	config.InitConfig()
	database.InitDB()
	server := app.InitApp()
	address := fmt.Sprintf("%v:%v", config.ENV.API_HOST, config.ENV.API_PORT)
	server.Run(address)
}
