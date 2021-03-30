package handler

import (
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
		}

		if e := param.Binding(r, &params); e != nil {
			render.BadRequest(w, e)
			return
		}

		if params.WalletName == "" {
			render.BadRequest(w, errors.New("invalid wallet name"))
			return
		}

		key := mixin.GenerateEd25519Key()
		user, keyStore, err := dapp.Client.CreateUser(ctx, key, params.WalletName)

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
