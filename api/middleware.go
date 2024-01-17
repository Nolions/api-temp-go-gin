package api

import (
	"github.com/Nolions/api-temp-php/internal/httpError"
	"github.com/gin-gonic/gin"
	"github.com/tent/hawk-go"
)

func (handler Handler) hawkAuthentication(c *gin.Context) error {
	auth, err := hawk.NewAuthFromRequest(c.Request, handler.credentials(), nil)
	if err == httpError.ErrSiteDomainNotFound {
		return httpError.ErrSiteDomainNotFound
	} else if err == hawk.ErrNoAuth {
		return httpError.HandleUnauthorized(4010, err)
	} else if err == httpError.ErrAccessKeyIDNotFound {
		return httpError.HandleUnauthorized(4011, err)
	} else if err != nil {
		return c.Error(err)
	}

	if err = auth.Valid(); err != nil {
		return httpError.HandleUnauthorized(4012, err)
	}

	c.Set("auth", auth.Credentials.Data)
	c.Next()
	return nil
}

func (handler Handler) credentials() hawk.CredentialsLookupFunc {
	return func(c *hawk.Credentials) error {
		// TODO AccessKey verify

		return nil
	}
}
