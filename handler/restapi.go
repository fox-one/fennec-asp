package handler

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fennec/core"
	"fennec/handler/param"
	"fennec/handler/render"
	"net/http"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/go-chi/chi"
	"github.com/twitchtv/twirp"
)

func RestAPIHandler(dapp *core.Wallet) http.Handler {
	router := chi.NewRouter()
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.Error(w, twirp.NotFoundError("not found"))
	})

	router.Post("/users", createUserHandlerFunc(dapp))

	return router
}

func createUserHandlerFunc(dapp *core.Wallet) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var params struct {
			WalletName string `json:"wallet_name"`
			CipherType string `json:"cipher_type"`
			PrivateKey string `json:"private_key"`
		}

		if e := param.Binding(r, &params); e != nil {
			render.BadRequest(w, e)
			return
		}

		if params.WalletName == "" {
			render.BadRequest(w, errors.New("invalid wallet name"))
			return
		}

		if params.PrivateKey == "" {
			render.BadRequest(w, errors.New("no private key"))
			return
		}

		if params.CipherType == "" {
			params.CipherType = "rsa"
		}

		var signer crypto.Signer

		if params.CipherType == "ed25519" {
			prv, err := parseED25519PrivateKeyFromStr(params.PrivateKey)
			if err != nil {
				render.BadRequest(w, errors.New("parse ed25519 private key error"))
				return
			}
			signer = prv
		} else {
			prv, err := parseRSAPrivateKeyFromPEMStr(params.PrivateKey)
			if err != nil {
				render.BadRequest(w, errors.New("parse rsa private key error"))
				return
			}
			signer = prv
		}

		user, keyStore, err := dapp.Client.CreateUser(ctx, signer, params.WalletName)

		if err != nil {
			render.BadRequest(w, err)
			return
		}

		var response struct {
			User     *mixin.User     `json:"user"`
			Keystore *mixin.Keystore `json:"keystore"`
		}

		response.User = user
		response.Keystore = keyStore

		render.JSON(w, response)
	}
}

func parseRSAPrivateKeyFromPEMStr(pemStr string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func parseED25519PrivateKeyFromStr(str string) (ed25519.PrivateKey, error) {
	content, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}

	return content, nil
}
