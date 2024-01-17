package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Nolions/api-temp-php/api"
	"github.com/Nolions/api-temp-php/config"
	"github.com/redpkg/formula/log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	version  string = "unknown"
	confPath string
)

func main() {
	flag.StringVar(&confPath, "c", "app.conf.yaml", "default config path")
	flag.Parse()

	conf, err := config.Init(confPath)
	if err != nil {
		panic(fmt.Sprintf("Error using config file: [%s]", confPath))
	}

	if err = log.Init(conf.Log); err != nil {
		log.Fatal().Msgf("Error init log: [%v]", err)
	}

	log.Info().Msgf("(%s) version: %s", conf.Project, version)

	httpSrv := api.New(context.Background(), conf.Project, conf.App, conf.Redis, conf.DB)
	httpSrv.Run()

	wg := sync.WaitGroup{}

	shutdown(conf.Project, &wg, httpSrv)
}

func shutdown(project string, wg *sync.WaitGroup, server *api.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	s := <-quit
	log.Info().Msgf("get a signal %s. (%s) Server is shutting down ...", s.String(), project)

	//wait for consumers done.
	wg.Wait()

	// close http server with timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("http server shutdown: %v", err)
	}

	log.Info().Msgf("(%s) Server is exit.", project)
}
