package cmd

import (
	"context"
	"fennec/handler/api"
	"fmt"
	"net/http"
	"time"

	"github.com/drone/signal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run Fennec ASP server",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		host, _ := cmd.Flags().GetString("host")
		addr := fmt.Sprintf("%s:%d", host, port)
		ctx := context.Background()

		api := api.New(debugMode,
			rootCmd.Version,
			provideDapp(),
		)

		srv := &http.Server{
			Addr:    addr,
			Handler: api.Handler(),
		}

		done := make(chan struct{}, 1)
		signal.WithContextFunc(ctx, func() {
			logrus.Debug("shutdown server...")

			// create context with timeout
			ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
			defer cancel()

			if err := srv.Shutdown(ctx); err != nil {
				logrus.WithError(err).Error("graceful shutdown server failed")
			}

			close(done)
		})

		logrus.Println("serve at", addr)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			logrus.WithError(err).Fatal("server aborted")
		}

		<-done
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().Int("port", 8081, "server port")
	serverCmd.Flags().String("host", "", "server host")
}
