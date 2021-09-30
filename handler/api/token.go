package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/fox-one/pkg/uuid"
)

func (s Server) MixinToken(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r.Header.Del("X-Forwarded-For")
		r.RemoteAddr = ""
		r.Host = ""

		reqID := r.Header.Get("X-Request-Id")
		if _, err := uuid.FromString(reqID); err != nil {
			reqID = uuid.New()
			r.Header.Set("X-Request-Id", reqID)
		}

		var body []byte
		if r.Body != nil {
			body, _ = ioutil.ReadAll(r.Body)
			r.Body.Close()
			r.Body = ioutil.NopCloser(bytes.NewReader(body))
		}

		sig := mixin.SignRaw(r.Method, r.URL.String(), body)
		token := s.client.Signer.SignToken(sig, reqID, time.Minute)
		r.Header.Set("Authorization", "Bearer "+token)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
