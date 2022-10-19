package httpError

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gitlab.com/ht-co/wraperr"
)

type HandlerFunc func(c *gin.Context) error

func Code(code int) int {
	return code
}

type ErrorResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ErrHandler(h HandlerFunc) gin.HandlerFunc {
	statusCode := http.StatusInternalServerError

	resp := &ErrorResp{
		Code:    5000,
		Message: "Internal server error",
	}

	return func(c *gin.Context) {
		var err error
		err = h(c)
		if err == nil {
			return
		}
		switch e := err.(type) {
		case *wraperr.Error:
			statusCode = e.StatusCode
			resp.Code = e.Code
			resp.Message = e.Message
		case validator.ValidationErrors:
			statusCode = http.StatusUnprocessableEntity
			resp.Code = Code(4220)
			resp.Message = e.Error()
		default:
			//statusCode = e.Code
			//resp.Code = Code(e.Code * 10)
			//resp.Message = fmt.Sprintf("%v", e.Message)
		}

		c.JSON(statusCode, resp)
		c.Abort()
	}
}
