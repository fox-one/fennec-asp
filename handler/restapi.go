package handler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/oxtoacart/bpool"
)

const (
	mixinEndpoint = "https://api.mixin.one"
)

var (
	xAuthorization = http.CanonicalHeaderKey("Authorization")
)

func MixinProxy(auth *mixin.KeystoreAuth) http.Handler {
	endpoint, _ := url.Parse(mixinEndpoint)

	return &httputil.ReverseProxy{
		BufferPool: bpool.NewBytePool(16, 1024*8),
		Director: func(req *http.Request) {
			//remove "/api/v1"
			req.URL.Path = req.URL.Path[len("/api/v1"):]

			if token := req.Header.Get(xAuthorization); token == "" {
				var body []byte
				if req.Body != nil {
					body, _ = ioutil.ReadAll(req.Body)
					_ = req.Body.Close()
					req.Body = ioutil.NopCloser(bytes.NewReader(body))
				}

				fmt.Println(req.URL.String())
				sig := mixin.SignRaw(req.Method, req.URL.String(), body)
				token := auth.SignToken(sig, mixin.RandomTraceID(), time.Minute)
				req.Header.Set(xAuthorization, "Bearer "+token)
			}

			req.Header["X-Forwarded-For"] = nil
			req.Host = endpoint.Host
			req.URL.Host = endpoint.Host
			req.URL.Scheme = endpoint.Scheme
		},
		ModifyResponse: func(resp *http.Response) error {
			resp.Header.Set("Access-Control-Allow-Origin", "*, *")
			resp.Header.Set("Access-Control-Allow-Headers", "*")
			resp.Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
			resp.Header.Set("Access-Control-Allow-Credentials", "false")
			resp.Header.Del("X-Request-Id")

			return nil
		},
	}
}
