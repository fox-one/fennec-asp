package core

import "github.com/fox-one/mixin-sdk-go"

type Config struct {
	Dapp Dapp `json:"dapp"`
}

type Dapp struct {
	mixin.Keystore
	ClientSecret string `json:"client_secret"`
	Pin          string `json:"pin"`
}
