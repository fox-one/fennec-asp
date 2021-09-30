package api

import (
	"fennec/handler/ip"
	"fennec/handler/logger"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	endpoint = "https://api.mixin.one"
)

type (
	Server struct {
		debug     bool
		version   string
		startedAt time.Time
		proxy     *httputil.ReverseProxy

		client *mixin.Client
	}
)

func New(
	debug bool,
	version string,
	client *mixin.Client,
) *Server {
	mixinEndpoint, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}

	return &Server{
		debug:     debug,
		version:   version,
		startedAt: time.Now(),
		client:    client,
		proxy:     httputil.NewSingleHostReverseProxy(mixinEndpoint),
	}
}

func (s Server) Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(
		middleware.Recoverer,
		ip.WithClientIP,
		logger.Logger,
		s.TrimV1Prefix,
		s.MixinToken,
	)

	r.Get("/hc", s.healthCheck)
	r.Get("/ip", ip.Handle())
	r.Post("/users", http.HandlerFunc(s.proxy.ServeHTTP))

	{
		// for old version...
		r.Post("/api/v1/users", http.HandlerFunc(s.proxy.ServeHTTP))
	}

	return r
}

func (s Server) TrimV1Prefix(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api/v1")
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
