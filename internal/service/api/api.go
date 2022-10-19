package api

import (
	"context"
	"fmt"
	"github.com/Nolions/api-temp-php/internal/httpError"
	"net/http"

	"github.com/Nolions/api-temp-php/config"
	"github.com/gin-gonic/gin"
	"github.com/redpkg/formula/db"
	"github.com/redpkg/formula/log"
	"github.com/redpkg/formula/redis"
)

type Server struct {
	httpServer *http.Server
	handler    Handler
}

func New(cachePrefix string, appConf config.App, apiConf config.Api, redisConf redis.Config, dbConf db.Config) *Server {
	handler := newHandler(cachePrefix, appConf, apiConf, redisConf, dbConf)

	e := engine(appConf.Mode)

	handler.router(e)

	addr := fmt.Sprintf(":%s", appConf.Addr)
	srv := &http.Server{
		Addr:    addr,
		Handler: e,
	}

	return &Server{
		httpServer: srv,
		handler:    handler,
	}
}

func (srv Server) Run() {
	log.Info().Msgf("http srv start listening on %s", srv.httpServer.Addr)
	go func() {
		err := srv.httpServer.ListenAndServe()

		if err == http.ErrServerClosed {
			log.Info().Msgf("api %v.", err)
			return
		}

		if err != nil {
			log.Fatal().Msgf("http server listen err: %v", err)
		}
	}()

}

func engine(mode string) *gin.Engine {
	e := &gin.Engine{}

	gin.SetMode(mode)
	e = gin.New()
	e.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/healthz"),
		gin.Recovery(),
	)
	//e = gin.Default()
	e.Use(gin.Recovery())
	e.HandleMethodNotAllowed = true
	e.NoMethod(httpError.ErrHandler(httpError.HandleNoAllowMethod))
	e.NoRoute(httpError.ErrHandler(httpError.HandleNotFound))

	return e
}

func (srv Server) Shutdown(ctx context.Context) error {
	srv.httpServer.Close()
	return srv.httpServer.Shutdown(ctx)
}
