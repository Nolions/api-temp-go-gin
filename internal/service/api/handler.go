package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Nolions/api-temp-php/config"
	"github.com/Nolions/api-temp-php/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/redpkg/formula/db"
	"github.com/redpkg/formula/log"
	"github.com/redpkg/formula/redis"
)

type Handler struct {
	CachePrefix string
	rep         repository.Repository
}

func newHandler(cachePrefix string, appConf config.App, configApi config.Api, redisConf redis.Config, dbConf db.Config) Handler {
	r, err := redis.New(redisConf)
	if err != nil {
		log.Fatal().Msgf("Failed to new redis: [%v]", err)
	}

	d, err := db.New(dbConf)
	if appConf.Mode == gin.DebugMode || appConf.Mode == gin.TestMode {
		d.ShowSQL()
	}

	if err != nil {
		log.Fatal().Msgf("Failed to new db: [%v]", err)
	}

	rep := repository.New(d, r, cachePrefix)

	return Handler{
		rep: rep,
	}
}

func (handler Handler) router(e *gin.Engine) {
	e.GET("/healthz", handler.healthz)
}

func (handler Handler) healthz(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (handler Handler) validate(req interface{}) error {
	validate := validator.New()
	err := validate.RegisterValidation("amount", validateInteger)
	err = validate.Struct(req)
	if err != nil {
		return err
	}

	return nil
}

func validateInteger(fl validator.FieldLevel) bool {
	amount := fl.Field().Float()
	dp := strings.Split(fmt.Sprintf("%v", amount), ".")

	if len(dp) < 2 {
		return true
	}

	return false
}
