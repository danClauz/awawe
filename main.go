package main

import (
	"awawe/cmd/server"
	config "awawe/configuration"
	"awawe/infrastructure/datastore"
)

//func startTrace() {
//	f, err := os.Create("trace.out")
//	if err != nil {
//		log.Fatalf("failed to create trace output file: %v", err)
//	}
//	defer func() {
//		if err := f.Close(); err != nil {
//			log.Fatalf("failed to close trace file: %v", err)
//		}
//	}()
//
//	if err := trace.Start(f); err != nil {
//		log.Fatalf("failed to start trace: %v", err)
//	}
//	defer trace.Stop()
//
//	// your program here
//}

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
