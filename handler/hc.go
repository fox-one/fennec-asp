package handler

import (
	"fennec/handler/render"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func HealthCheckHandler(version string) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.NoCache)
	r.Handle("/", healthCheckHandlerFunc(version))

	return r
}

func healthCheckHandlerFunc(version string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uptime := time.Since(time.Now()).Truncate(time.Microsecond)
		render.JSON(w, render.H{
			"uptime":  uptime.String(),
			"version": version,
		})
	}
}
