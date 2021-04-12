package cmd

import (
	"encoding/base64"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/spf13/cobra"
)

var ed25519 = &cobra.Command{
	Use: "ed25519",
	Run: func(cmd *cobra.Command, args []string) {
		signer := mixin.GenerateEd25519Key()
		cmd.Println(base64.StdEncoding.EncodeToString([]byte(signer)))
	},
}

func init() {
	rootCmd.AddCommand(ed25519)
}
