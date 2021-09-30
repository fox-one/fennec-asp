package ip

import (
	"fennec/handler/render"
	"fennec/handler/request"
	"net"
	"net/http"
	"strings"
)

func WithClientIP(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(
			request.WithClientIP(r.Context(), getClientIP(r)),
		))
	}

	return http.HandlerFunc(fn)
}

func Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.Data(w, getClientIP(r))
	}
}

func getClientIP(req *http.Request) string {
	clientIP := req.Header.Get("X-Forwarded-For")
	clientIP = strings.TrimSpace(strings.Split(clientIP, ",")[0])
	if clientIP == "" {
		clientIP = strings.TrimSpace(req.Header.Get("X-Real-Ip"))
	}

	if clientIP != "" {
		return clientIP
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(req.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
