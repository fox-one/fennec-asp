package cmd

import (
	"fennec/core"

	"github.com/fox-one/mixin-sdk-go"
)

func provideConfig() *core.Config {
	return &cfg
}

func provideDapp() *core.Wallet {
	c, err := mixin.NewFromKeystore(&cfg.Dapp.Keystore)
	if err != nil {
		panic(err)
	}

	return &core.Wallet{
		Client: c,
		Pin:    cfg.Dapp.Pin,
	}
}
