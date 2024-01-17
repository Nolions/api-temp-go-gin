package api

import (
	"encoding/json"
	"github.com/Nolions/api-temp-php/internal/httpError"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var handler Handler
var resp *httptest.ResponseRecorder
var c *gin.Context
var e *gin.Engine

func setup() {
	resp = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(resp)
	e = engine("")
	handler = Handler{Ctx: c}
}

func TestPageNotFount(t *testing.T) {
	setup()
	req, _ := http.NewRequest("GET", "/test", nil)
	e.ServeHTTP(resp, req)

	str, _ := json.Marshal(httpError.ErrorResp{
		Code:    4040,
		Message: http.StatusText(http.StatusNotFound),
	})

	assert.Equal(t, http.StatusNotFound, resp.Code)
	assert.Contains(t, string(str), resp.Body.String())
}

func TestMethodNoAllow(t *testing.T) {
	setup()
	e.GET("/test", httpError.ErrHandler(func(c *gin.Context) error {
		return httpError.ErrMethodNoAllow
	}))

	req, _ := http.NewRequest(http.MethodPost, "/test", nil)

	e.ServeHTTP(resp, req)

	str, _ := json.Marshal(httpError.ErrorResp{
		Code:    4051,
		Message: http.StatusText(http.StatusMethodNotAllowed),
	})

	assert.Equal(t, http.StatusMethodNotAllowed, resp.Code)
	assert.Contains(t, string(str), resp.Body.String())
}

func TestHealthz(t *testing.T) {
	setup()

	e.GET("/healthz", handler.healthz)

	req, _ := http.NewRequest(http.MethodGet, "/healthz", nil)
	e.ServeHTTP(resp, req)

	actually := struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}

	b, _ := json.Marshal(actually)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, string(b), resp.Body.String())
}
