package main

import (
	"awawe/cmd/server"
	config "awawe/configuration"
	"awawe/infrastucture/datastore"
)

func main() {
	config.InitializeConfig()

	switch config.GetAppConfig().Command {
	case "migrate":
		datastore.MigrateMySQLDatabase(nil)
	case "server":
		server.StartServer()
	default:
		panic("App command is invalid!")
	}
}
