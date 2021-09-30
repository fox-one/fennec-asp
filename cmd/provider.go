package cmd

import (
	"github.com/fox-one/mixin-sdk-go"
)

func provideDapp() *mixin.Client {
	client, err := mixin.NewFromKeystore(&cfg.Dapp)
	if err != nil {
		panic(err)
	}

	return client
}
