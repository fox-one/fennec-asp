package cmd

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/spf13/cobra"
)

var rsaCmd = &cobra.Command{
	Use: "rsa",
	Run: func(cmd *cobra.Command, args []string) {
		prv := generateRSAKey()
		pubKeyBytes := x509.MarshalPKCS1PublicKey(&prv.PublicKey)

		pubPem := pem.EncodeToMemory(
			&pem.Block{
				Type:  "RSA PUBLIC KEY",
				Bytes: pubKeyBytes,
			},
		)

		cmd.Println(string(pubPem))
	},
}

func generateRSAKey() *rsa.PrivateKey {
	private, _ := rsa.GenerateKey(rand.Reader, 1024)
	return private
}

func init() {
	rootCmd.AddCommand(rsaCmd)
}
