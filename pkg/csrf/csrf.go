package csrf

import (
	"crypto/sha256"
	"encoding/base64"
	"io"

	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
)

const (
	CSRFHeader = "X-CSRF-Token"
	// 32 bytes
	csrfSalt = "CX3pz55ZBIcNXikGR08X8ZNeiTYiYTNS"
)

// Create CSRF Token
func MakeToken(sid string, logger logger.Logger) string {
	hash := sha256.New()
	_, err := io.WriteString(hash, csrfSalt+sid)
	if err != nil {
		logger.Errorf("Make CSRF Token", err)
	}

	token := base64.RawStdEncoding.EncodeToString(hash.Sum(nil))
	return token
}

// Validate CSRF Token
func ValidateToken(token string, sid string, logger logger.Logger) bool {
	csrfToken := MakeToken(sid, logger)
	return token == csrfToken
}
