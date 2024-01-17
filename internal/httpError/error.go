package httpError

import (
	"github.com/Nolions/wraperr"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	// ErrInsufficientBalance 402
	ErrInsufficientBalance = wraperr.New(http.StatusPaymentRequired, Code(4020), "Insufficient balance", nil)
	// ErrMemberDisable 403
	ErrMemberDisable = wraperr.New(http.StatusForbidden, Code(4030), "This member is disabled", nil)
	// ErrForbiddenMemberType 403
	ErrForbiddenMemberType = wraperr.New(http.StatusForbidden, Code(4031), "This member type is not allowed access to the resource", nil)
	// ErrPageNotFount 404
	ErrPageNotFount = wraperr.New(http.StatusNotFound, Code(4040), http.StatusText(http.StatusNotFound), nil)
	// ErrAccessKeyIDNotFound 404
	ErrAccessKeyIDNotFound = wraperr.New(http.StatusNotFound, Code(4041), "Access key id not found", nil)
	// ErrMemberNotFound 404
	ErrMemberNotFound = wraperr.New(http.StatusNotFound, Code(4042), "Member not found", nil)
	// ErrTxOrderNotFound 404
	ErrTxOrderNotFound = wraperr.New(http.StatusNotFound, Code(4043), "Transaction order not found", nil)
	// ErrSiteDomainNotFound 404
	ErrSiteDomainNotFound = wraperr.New(http.StatusNotFound, Code(4044), "Site domain not found", nil)
	// ErrOrderNotFound 404
	ErrOrderNotFound = wraperr.New(http.StatusNotFound, Code(4046), "Order not found", nil)
	// ErrSiteNotFound 404
	ErrSiteNotFound = wraperr.New(http.StatusNotFound, Code(4047), "Site not found", nil)
	// ErrGuestNotFound 404
	ErrGuestNotFound = wraperr.New(http.StatusNotFound, Code(4048), "Guest not found", nil)
	// ErrMethodNoAllow 405
	ErrMethodNoAllow = wraperr.New(http.StatusMethodNotAllowed, Code(4051), http.StatusText(http.StatusMethodNotAllowed), nil)
	// ErrInvalidTimeAfter 422
	ErrInvalidTimeAfter = wraperr.New(http.StatusUnprocessableEntity, Code(4220), "The end at must be a date after start at", nil)
	// ErrInvalidTimeDiff 422
	ErrInvalidTimeDiff = wraperr.New(http.StatusUnprocessableEntity, Code(4220), "Time difference must be less or equal '7' days", nil)
	// ErrInvalidPage 422
	ErrInvalidPage = wraperr.New(http.StatusUnprocessableEntity, Code(4220), "Page must be greater or equal '1'", nil)
	// ErrInvalidPerPage 422
	ErrInvalidPerPage = wraperr.New(http.StatusUnprocessableEntity, Code(4220), "Per page must be greater or equal '1' and less or equal '500'", nil)
	// ErrUsernameExisted 422
	ErrUsernameExisted = wraperr.New(http.StatusUnprocessableEntity, Code(4221), "The username has already been taken", nil)
	// ErrDuplicateOrderID 422
	ErrDuplicateOrderID = wraperr.New(http.StatusUnprocessableEntity, Code(4222), "Duplicate order id", nil)
)

func HandleNotFound(c *gin.Context) error {
	return ErrPageNotFount
}

func HandleNoAllowMethod(c *gin.Context) error {
	return ErrMethodNoAllow
}

func HandleUnauthorized(code int, err error) error {
	return wraperr.New(http.StatusUnauthorized, Code(code), err.Error(), err)
}

func FailCallError(code int, err error) error {
	return wraperr.New(http.StatusUnauthorized, Code(6000), err.Error(), err)
}
