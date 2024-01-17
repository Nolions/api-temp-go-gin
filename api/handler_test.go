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

var mockEngine *gin.Engine

func init() {
	//mockEngine = engine("test")
	//mockEngine.s
}

var app Handler
var resp *httptest.ResponseRecorder
var c *gin.Context
var e *gin.Engine

func setup() {
	resp = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(resp)
	e = engine("")
	app = Handler{Ctx: c}
}

func TestPageNotFount(t *testing.T) {
	setup()
	req, _ := http.NewRequest("GET", "/test", nil)
	e.ServeHTTP(resp, req)

	//httpError.ErrHandler(func(c *gin.Context) error {
	//	return httpError.ErrPageNotFount
	//})

	str, _ := json.Marshal(httpError.ErrorResp{
		Code:    4040,
		Message: http.StatusText(http.StatusNotFound),
	})

	assert.Equal(t, http.StatusNotFound, resp.Code)
	assert.Contains(t, string(str), resp.Body.String())
}

func TestMethodNoAllow(t *testing.T) {
	mockEngine.GET("/test", httpError.ErrHandler(func(c *gin.Context) error {
		return httpError.ErrMethodNoAllow
	}))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/test", nil)

	mockEngine.ServeHTTP(w, req)

	str, _ := json.Marshal(httpError.ErrorResp{
		Code:    4051,
		Message: http.StatusText(http.StatusMethodNotAllowed),
	})

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	assert.Contains(t, string(str), w.Body.String())
}

func TestHealthz(t *testing.T) {
	h := Handler{}
	mockEngine.GET("/healthz", h.healthz)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/healthz", nil)
	mockEngine.ServeHTTP(w, r)

	str, _ := json.Marshal(struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	})

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, string(str), w.Body.String())
}
