package helper

import (
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
)

// NewUUIDV4 generate uuid v4 no hyphens
func NewUUIDV4() string {
	return hex.EncodeToString(uuid.NewV4().Bytes())
}
