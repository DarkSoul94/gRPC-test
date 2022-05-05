package main

import (
	"github/DarkSoul94/gRPC-test/server/grpcserver"
	"github/DarkSoul94/gRPC-test/server/httpserver"
	"github/DarkSoul94/gRPC-test/server/pkg/config"
	"log"
	"os"
	"os/signal"

	"github.com/spf13/viper"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatal(err)
	}

	appHTTP := httpserver.NewApp()
	go appHTTP.Run(viper.GetString("app.http_port"))

	appgRPC := grpcserver.NewApp()
	go appgRPC.Run(viper.GetString("app.grpc_port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	appHTTP.Stop()
	appgRPC.Stop()
}
