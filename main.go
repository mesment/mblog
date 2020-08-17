package main

import (
	"fmt"
	zlog "github.com/mesment/mblog/pkg/logger"
	"github.com/mesment/mblog/pkg/setting"
	"github.com/mesment/mblog/routers"
	"os"
	"os/signal"
	"time"
	// "github.com/fvbock/endless"
	"context"
	"net/http"
)

func main() {
	// init logger fist
	log := zlog.Init()
	log.Infow("start server", "listening", setting.HTTPPort )

	router := routers.InitRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Infof("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit
	log.Infof("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown:%v", err)
	}
	log.Info("Server exiting")
}
