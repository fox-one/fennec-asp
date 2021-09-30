package api

import (
	"fennec/handler/render"
	"net/http"
	"time"
)

func (s Server) healthCheck(w http.ResponseWriter, r *http.Request) {
	render.Data(w, map[string]interface{}{
		"uptime":  time.Since(s.startedAt).String(),
		"version": s.version,
	})
}
