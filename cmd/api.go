package cmd

import (
	"fennec/handler"
	"fmt"
	"net/http"

	"github.com/fox-one/pkg/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "run api server",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		log := logger.FromContext(ctx)

		dapp := provideDapp()

		mux := chi.NewMux()
		mux.Use(middleware.Recoverer)
		mux.Use(middleware.StripSlashes)
		mux.Use(cors.AllowAll().Handler)
		mux.Use(logger.WithRequestID)
		mux.Use(middleware.Logger)
		mux.Use(middleware.NewCompressor(5).Handler)

		{
			mux.Mount("/hc", handler.HealthCheckHandler(rootCmd.Version))
		}

		{
			mux.Mount("/api/v1", handler.RestAPIHandler(dapp))
		}

		port, _ := cmd.Flags().GetInt("port")
		addr := fmt.Sprintf(":%d", port)
		server := http.Server{
			Addr:    addr,
			Handler: mux,
		}

		log.Infoln("api serve at ", addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.WithError(err).Fatal("server aborted")
		}
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	apiCmd.Flags().IntP("port", "p", 80, "server port")
}
