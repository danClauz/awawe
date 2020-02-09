package server

import (
	config "awawe/configuration"
	"awawe/infrastucture/router"
	"context"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func startServer(e *echo.Echo) {
	s := &http.Server{
		Addr:         config.GetServerConfig().Host + ":" + config.GetServerConfig().Port,
		ReadTimeout:  time.Duration(config.GetServerConfig().ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.GetServerConfig().WriteTimeout) * time.Second,
	}

	go func() {
		e.Logger.Fatal(e.StartServer(s))
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	fmt.Println("Server listen at http://" + config.GetServerConfig().Host + ":" + config.GetServerConfig().Port)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func StartServer() {
	startServer(router.NewRouter())
}
