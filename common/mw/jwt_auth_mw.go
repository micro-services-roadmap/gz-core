package mw

import (
	"github.com/zeromicro/go-zero/rest/handler"
	"net/http"
)

// JwtAuthMiddleware : with jwt on the verification, no jwt on the verification
type JwtAuthMiddleware struct {
	secret string
}

func NewCommonJwtAuthMiddleware(secret string) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{
		secret: secret,
	}
}

// Deprecated: this will validate by kong.
func (m *JwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header.Get("Authorization")) > 0 {
			//has jwt Authorization
			authHandler := handler.Authorize(m.secret)
			authHandler(next).ServeHTTP(w, r)
			return
		} else {
			//no jwt Authorization
			next(w, r)
		}
	}
}
