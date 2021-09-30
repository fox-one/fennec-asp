package config

import "github.com/fox-one/mixin-sdk-go"

type Config struct {
	Dapp mixin.Keystore `json:"dapp"`
}
