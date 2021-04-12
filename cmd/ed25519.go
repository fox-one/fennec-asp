package cmd

import (
	"crypto/ed25519"
	"encoding/base64"

	"github.com/spf13/cobra"
)

var ed25519Cmd = &cobra.Command{
	Use: "ed25519",
	Run: func(cmd *cobra.Command, args []string) {
		pubK, privK, _ := ed25519.GenerateKey(nil)

		cmd.Println("priv: ", base64.StdEncoding.EncodeToString(privK))
		cmd.Println("pub: ", base64.StdEncoding.EncodeToString(pubK))
	},
}

func init() {
	rootCmd.AddCommand(ed25519Cmd)
}
