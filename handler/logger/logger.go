package logger

import (
	"fennec/handler/request"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
)

func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		next.ServeHTTP(ww, r)

		end := time.Now()
		ua := r.Header.Get("User-Agent")
		reqId := r.Header.Get("X-Request-Id")
		ip := request.ClientIPFrom(r.Context())
		log.Println(fmt.Sprintf(`%s "%s %s" %d %d %v %s "%s"`, ip, r.Method, r.RequestURI, ww.Status(), ww.BytesWritten(), end.Sub(start), reqId, ua))
	}

	return http.HandlerFunc(fn)
}
