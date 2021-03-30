package core

import "github.com/fox-one/mixin-sdk-go"

// Wallet wallet
type Wallet struct {
	Client *mixin.Client `json:"client"`
	Pin    string        `json:"pin"`
}
