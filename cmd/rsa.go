package cmd

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/spf13/cobra"
)

var rsaCmd = &cobra.Command{
	Use: "rsa",
	Run: func(cmd *cobra.Command, args []string) {
		// prv := generateRSAKey()
		// pubKeyBytes := x509.MarshalPKCS1PublicKey(&prv.PublicKey)

		// pubPem := pem.EncodeToMemory(
		// 	&pem.Block{
		// 		Type:  "RSA PUBLIC KEY",
		// 		Bytes: pubKeyBytes,
		// 	},
		// )

		// cmd.Println(string(pubPem))

		s := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsjVrFzquCMJOL7oXDWQ7
Jm9iy7ndIF8rkJ/dEtgHQl1T4lSEGEb1I4Q9dPRlp+V1TkbEnJrDDP+IdQf6m+oI
Elh1w7dIqzCAPdqGF7Jjs/Jdq7f9SAn2zmM2AaidfMziOW8rAK7ji31KA2GTUeQb
Eti5ujPAB54uno7CbDwvPMIMhi7YYp7rCh7hlMWDSQjAYTvyIIzTFklXFNzRSlDo
hLxad3GimqqFK1K8/NtkSSAriQO0PvllKz1pA38HxCpmGOh1oLW7oj8IPBbo3YS2
jY+lTRzTWmIQ4JS6jAauzz8NTmGKAQ/mCiZpasmbxi7t/bHx9GJIa2NWJvUJU9AR
WQIDAQAB
-----END PUBLIC KEY-----`

		block, _ := pem.Decode([]byte(s))
		if block == nil {
			panic(errors.New("failed to parse PEM block containing the key"))
		}

		_, e := x509.ParsePKIXPublicKey(block.Bytes)

		cmd.Println(e)
	},
}

func generateRSAKey() *rsa.PrivateKey {
	private, _ := rsa.GenerateKey(rand.Reader, 1024)
	return private
}

func init() {
	rootCmd.AddCommand(rsaCmd)
}
